# VPC Flow Log Parser

A Go-based tool to parse AWS VPC Flow Logs, map each log entry to predefined tags based on destination port and protocol, and generate summary reports.

## Features

- **Efficient Parsing:** Handles large VPC Flow Log files (up to 10 MB) efficiently.
- **Flexible Tagging:** Maps log entries to tags based on a customizable lookup table.
- **Comprehensive Reporting:** Generates counts of tags and port/protocol combinations.
- **Extensible:** Easily extendable to support additional protocols or functionalities.
- **Logging:** Provides informative logs for easier debugging and monitoring.

## Prerequisites

- Go (version 1.16 or later) [https://golang.org/dl/]
- Git (for cloning the repository)

## Installation

1. **Clone the Repository:**

   ```bash
   git clone [https://github.com/vishisth29/vpc-flowlog-parser.git](https://github.com/vishisth29/vpc-flowlog-parser.git)
   cd vpc-flowlog-parser

2. **Initialize Go Modules:**

   ```bash
   go mod tidy

3. **Build the Application:**

   ```bash
   chmod +x scripts/build.sh
   ./scripts/build.sh

## Usage
The vpc-flowlog-parser executable can be run with command-line flags to specify input and output files.

**Command-Line Flags**
-flowlog : Path to the VPC Flow Log file. (Default: examples/flow_log.txt)
-lookup : Path to the lookup CSV file. (Default: examples/lookup.csv)
-output : Path to the output file. (Default: output.txt)


1. **Running the Example**

   ```bash
   ./bin/flowlogparser -flowlog=examples/flow_log.txt -lookup=examples/lookup.csv -output=output.txt


2. **Running Your Own Test Files**

   ```bash
   ./bin/flowlogparser -flowlog=path/to/<TEST_FLOW_LOG>.txt -lookup=path/to/<TEST_LOOKUP>.csv -output=path/to/<TEST_OUTPUT>.txt

# Further Improvements for a Production System

While the current implementation is robust for handling up to 10 MB of flow log data and 10,000 lookup mappings, the following enhancements can be made to scale and optimize the tool for production environments:

## 1. Concurrency and Parallel Processing

- **Goroutines and Channels**:  
  Utilize Go's concurrency features to process multiple lines of the flow log simultaneously, improving performance on multi-core systems.

- **Worker Pools**:  
  Implement worker pools to manage the number of concurrent processing threads, preventing resource exhaustion.

## 2. Enhanced Protocol Support

- **Dynamic Protocol Mapping**:  
  Instead of a static map, fetch protocol mappings from an external source or allow users to define them in a configuration file.

- **Support for IPv6**:  
  Extend parsing capabilities to handle IPv6 addresses and related flow log fields.

## 3. Improved Error Handling and Reporting

- **Detailed Logs**:  
  Implement more granular logging levels (e.g., DEBUG, INFO, WARN, ERROR) and log rotation to manage log file sizes.

- **Error Metrics**:  
  Collect metrics on parsing errors and warnings for monitoring purposes.

## 4. Configuration Management

- **Configuration Files**:  
  Allow users to define configurations (e.g., file paths, protocol mappings) in YAML or TOML files instead of command-line flags.

- **Environment Variables**:  
  Support environment variables for configuration to integrate with containerized deployments.

## 5. Scalability and Performance Optimization

- **Streaming Processing**:  
  Implement streaming parsers to handle very large flow log files without loading the entire file into memory.

- **Caching Mechanisms**:  
  Use caching for frequently accessed lookup mappings to reduce lookup times.

"
