#!/bin/bash

OUTPUT_DIR="bin"
BINARY_NAME="get_day"

if [ ! -d "$OUTPUT_DIR" ]; then
    mkdir "$OUTPUT_DIR"
fi

echo "Building..."
go build -o "$OUTPUT_DIR/$BINARY_NAME" get_day.go
echo "Build complete. Binary available at $OUTPUT_DIR/$BINARY_NAME"
