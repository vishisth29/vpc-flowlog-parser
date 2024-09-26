package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/vishisth/vpc-flowlog-parser/internal/lookup"
	"github.com/vishisth/vpc-flowlog-parser/internal/parser"
	"github.com/vishisth/vpc-flowlog-parser/pkg/logger"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Define command-line flags
	flowLogPath := flag.String("flowlog", "examples/flow_log.txt", "Path to the VPC Flow Log file")
	lookupPath := flag.String("lookup", "examples/lookup.csv", "Path to the lookup CSV file")
	outputPath := flag.String("output", "output.txt", "Path to the output file")

	flag.Parse()

	// Resolve absolute paths
	flowLogAbsPath, err := filepath.Abs(*flowLogPath)
	if err != nil {
		log.Fatalf("Error resolving flow log file path: %v", err)
	}
	lookupAbsPath, err := filepath.Abs(*lookupPath)
	if err != nil {
		log.Fatalf("Error resolving lookup file path: %v", err)
	}
	outputAbsPath, err := filepath.Abs(*outputPath)
	if err != nil {
		log.Fatalf("Error resolving output file path: %v", err)
	}

	logger.Infof("Flow Log File: %s", flowLogAbsPath)
	logger.Infof("Lookup File: %s", lookupAbsPath)
	logger.Infof("Output File: %s", outputAbsPath)

	// Load lookup table
	lookupMap, err := lookup.LoadLookupTable(lookupAbsPath)
	if err != nil {
		log.Fatalf("Error loading lookup table: %v", err)
	}

	// Process flow logs
	tagCounts, portProtocolCounts, err := parser.ProcessFlowLogs(flowLogAbsPath, lookupMap)
	if err != nil {
		log.Fatalf("Error processing flow logs: %v", err)
	}

	// Write output
	err = parser.WriteOutput(outputAbsPath, tagCounts, portProtocolCounts)
	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}

	fmt.Printf("Processing complete. Output written to '%s'.\n", outputAbsPath)
}
