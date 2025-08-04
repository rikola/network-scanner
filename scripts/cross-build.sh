#!/bin/bash

# Cross-platform build script for Network Scanner
set -e

# Get the current directory
CURRENT_DIR=$(pwd)
BUILD_DIR="$CURRENT_DIR/build"
CMD_DIR="$CURRENT_DIR/cmd/scanner"
BINARY_NAME="scanner"

# Get version from git or use default
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS="-ldflags=-X main.version=$VERSION"

# Create build directory if it doesn't exist
mkdir -p "$BUILD_DIR"

echo "Cross-compiling $BINARY_NAME version $VERSION..."

# Build for Linux (amd64)
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build "$LDFLAGS" -o "$BUILD_DIR/$BINARY_NAME-linux-amd64" "$CMD_DIR"

# Build for Windows (amd64)
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build "$LDFLAGS" -o "$BUILD_DIR/$BINARY_NAME-windows-amd64.exe" "$CMD_DIR"

# Build for macOS (amd64)
echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build "$LDFLAGS" -o "$BUILD_DIR/$BINARY_NAME-darwin-amd64" "$CMD_DIR"

# Build for macOS (arm64)
echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build "$LDFLAGS" -o "$BUILD_DIR/$BINARY_NAME-darwin-arm64" "$CMD_DIR"

echo "Cross-compilation complete. Binaries available in $BUILD_DIR"
