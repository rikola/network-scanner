#!/bin/bash

# Build script for Network Scanner
set -e

# Get the current directory
CURRENT_DIR=$(pwd)
BUILD_DIR="$CURRENT_DIR/build"
CMD_DIR="$CURRENT_DIR/cmd/scanner"
BINARY_NAME="scanner"

# Get version from git or use default
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS="-ldflags \"-X main.version=$VERSION\""

# Create build directory if it doesn't exist
mkdir -p "$BUILD_DIR"

echo "Building $BINARY_NAME version $VERSION..."

# Build for the current platform
go build $LDFLAGS -o "$BUILD_DIR/$BINARY_NAME" "$CMD_DIR"

echo "Build complete: $BUILD_DIR/$BINARY_NAME"
