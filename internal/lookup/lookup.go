package lookup

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vishisth/vpc-flowlog-parser/pkg/logger"
)

// LookupMap is a nested map where the first key is dstport and the second key is protocol
type LookupMap map[string]map[string]string

// LoadLookupTable loads the lookup CSV file and returns a LookupMap
func LoadLookupTable(lookupFilePath string) (LookupMap, error) {
	file, err := os.Open(lookupFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open lookup file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// Read header
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header from lookup file: %v", err)
	}

	// Identify column indices
	dstportIdx, protocolIdx, tagIdx := -1, -1, -1
	for i, header := range headers {
		switch strings.ToLower(strings.TrimSpace(header)) {
		case "dstport":
			dstportIdx = i
		case "protocol":
			protocolIdx = i
		case "tag":
			tagIdx = i
		}
	}

	if dstportIdx == -1 || protocolIdx == -1 || tagIdx == -1 {
		return nil, fmt.Errorf("lookup file must contain 'dstport', 'protocol', and 'tag' columns")
	}

	lookup := make(LookupMap)

	// Read each record
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading lookup file: %v", err)
		}

		dstport := strings.TrimSpace(record[dstportIdx])
		protocol := strings.ToLower(strings.TrimSpace(record[protocolIdx]))
		tag := strings.TrimSpace(record[tagIdx])

		if dstport == "" || protocol == "" || tag == "" {
			continue // Skip incomplete records
		}

		if _, exists := lookup[dstport]; !exists {
			lookup[dstport] = make(map[string]string)
		}
		lookup[dstport][protocol] = tag
	}

	logger.Infof("Loaded %d port/protocol mappings from lookup table.", len(lookup))

	return lookup, nil
}
