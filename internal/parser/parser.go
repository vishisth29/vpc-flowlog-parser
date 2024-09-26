package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/vishisth/vpc-flowlog-parser/internal/lookup"
	"github.com/vishisth/vpc-flowlog-parser/internal/protocol"
	"github.com/vishisth/vpc-flowlog-parser/pkg/logger"
)

// ProcessFlowLogs processes the flow log file and returns tag counts and port/protocol counts
func ProcessFlowLogs(flowLogFilePath string, lookupMap lookup.LookupMap) (map[string]int, map[string]int, error) {
	file, err := os.Open(flowLogFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open flow log file: %v", err)
	}
	defer file.Close()

	tagCounts := make(map[string]int)
	portProtocolCounts := make(map[string]int)

	reader := csv.NewReader(file)
	reader.Comma = ' ' // Space-separated fields
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	lineNum := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		lineNum++
		if err != nil {
			logger.Warnf("Failed to read line %d: %v", lineNum, err)
			continue
		}

		// Handle empty lines
		if len(record) == 0 {
			continue
		}

		// Convert slice to a single string and split by whitespace to handle multiple spaces
		lineStr := strings.Join(record, " ")
		fields := strings.Fields(lineStr)

		// AWS VPC Flow Logs version 2 has 14 fields
		if len(fields) < 14 {
			logger.Warnf("Line %d is malformed and will be skipped.", lineNum)
			continue
		}

		// Extract dstport and protocol number based on AWS VPC Flow Logs version 2 format
		// Field indices (0-based):
		// 6: dstport
		// 7: protocol
		dstport := strings.TrimSpace(fields[6])
		protocolNum := strings.TrimSpace(fields[7])

		protocolName := protocol.GetProtocolName(protocolNum)

		if protocolName == "unknown" {
			logger.Warnf("Unknown protocol number '%s' on line %d.", protocolNum, lineNum)
		}

		// Update port/protocol counts
		portProtoKey := fmt.Sprintf("%s,%s", dstport, protocolName)
		portProtocolCounts[portProtoKey]++

		// Determine tag
		tag := "Untagged"
		if protocols, exists := lookupMap[dstport]; exists {
			if mappedTag, exists := protocols[protocolName]; exists {
				tag = mappedTag
			}
		}

		tagCounts[tag]++
	}

	return tagCounts, portProtocolCounts, nil
}

// WriteOutput writes the tag counts and port/protocol counts to the output file
func WriteOutput(outputFilePath string, tagCounts, portProtocolCounts map[string]int) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ','

	// Write Tag Counts
	_, err = file.WriteString("Tag Counts:\n\nTag,Count\n")
	if err != nil {
		return fmt.Errorf("failed to write to output file: %v", err)
	}

	// Sort tags alphabetically
	tags := make([]string, 0, len(tagCounts))
	for tag := range tagCounts {
		tags = append(tags, tag)
	}
	sort.Strings(tags)

	for _, tag := range tags {
		count := tagCounts[tag]
		line := []string{tag, strconv.Itoa(count)}
		if err := writer.Write(line); err != nil {
			return fmt.Errorf("failed to write tag counts: %v", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error flushing tag counts: %v", err)
	}

	// Add a newline between sections
	_, err = file.WriteString("\nPort/Protocol Combination Counts:\n\nPort,Protocol,Count\n")
	if err != nil {
		return fmt.Errorf("failed to write to output file: %v", err)
	}

	// Collect and sort port/protocol keys
	portProtoKeys := make([]string, 0, len(portProtocolCounts))
	for key := range portProtocolCounts {
		portProtoKeys = append(portProtoKeys, key)
	}

	sort.Slice(portProtoKeys, func(i, j int) bool {
		partsI := strings.Split(portProtoKeys[i], ",")
		partsJ := strings.Split(portProtoKeys[j], ",")

		// Convert ports to integers for proper numerical sorting
		portA, errA := strconv.Atoi(partsI[0])
		portB, errB := strconv.Atoi(partsJ[0])

		if errA != nil || errB != nil {
			// Fallback to string comparison if conversion fails
			return portProtoKeys[i] < portProtoKeys[j]
		}

		if portA != portB {
			return portA < portB
		}
		return partsI[1] < partsJ[1]
	})

	for _, key := range portProtoKeys {
		count := portProtocolCounts[key]
		parts := strings.Split(key, ",")
		if len(parts) != 2 {
			continue // skip malformed keys
		}
		port := parts[0]
		protocol := parts[1]
		line := []string{port, protocol, strconv.Itoa(count)}
		if err := writer.Write(line); err != nil {
			return fmt.Errorf("failed to write port/protocol counts: %v", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error flushing port/protocol counts: %v", err)
	}

	return nil
}
