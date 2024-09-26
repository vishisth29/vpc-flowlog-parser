# VPC Flow Log Parser

A Go-based tool to parse AWS VPC Flow Logs, map each log entry to predefined tags based on destination port and protocol, and generate summary reports.

## Features

- **Efficient Parsing:** Handles large VPC Flow Log files (up to 10 MB) efficiently.
- **Flexible Tagging:** Maps log entries to tags based on a customizable lookup table.
- **Comprehensive Reporting:** Generates counts of tags and port/protocol combinations.
- **Extensible:** Easily extendable to support additional protocols or functionalities.
- **Logging:** Provides informative logs for easier debugging and monitoring.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)
- Git (for cloning the repository)

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/vishisth29/vpc-flowlog-parser.git
   cd vpc-flowlog-parser
