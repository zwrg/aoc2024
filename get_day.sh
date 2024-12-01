#!/bin/bash

BINARY="bin/get_day"
BUILD_SCRIPT="scripts/build.sh"

if [ ! -f "$BINARY" ]; then
    echo "Binary not found. Running build script..."
    if [ ! -f "$BUILD_SCRIPT" ]; then
        echo "Build script not found. Exiting..."
        exit 1
    fi
    bash "$BUILD_SCRIPT"
fi

if [ -f "$BINARY" ]; then
    echo "Running $BINARY with parameter: $1"
    "$BINARY" "$1"
else
    echo "Build failed or binary not found. Exiting..."
    exit 1
fi
