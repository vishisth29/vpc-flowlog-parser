set -e

# Define variables
APP_NAME="flowlogparser"
OUTPUT_DIR="bin"

# Create output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# Build the application
go build -o $OUTPUT_DIR/$APP_NAME ./cmd/flowlogparser

echo "Build completed. Executable is located at $OUTPUT_DIR/$APP_NAME"
