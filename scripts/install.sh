#!/bin/bash

# Installation script for Network Scanner
set -e

# Get the current directory
CURRENT_DIR=$(pwd)
BUILD_DIR="$CURRENT_DIR/build"
CMD_DIR="$CURRENT_DIR/cmd/scanner"
BINARY_NAME="scanner"
INSTALL_DIR="/usr/local/bin"

# Check if running as root for system-wide install
if [ "$EUID" -ne 0 ] && [ "$INSTALL_DIR" = "/usr/local/bin" ]; then
    echo "Error: Please run as root for system-wide installation"
    echo "Use: sudo ./scripts/install.sh"
    echo "Or set a user-writable install location: INSTALL_DIR=~/bin ./scripts/install.sh"
    exit 1
fi

# Create build directory if it doesn't exist
mkdir -p "$BUILD_DIR"

# Get version from git or use default
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS="-ldflags \"-X main.version=$VERSION\""

echo "Building $BINARY_NAME version $VERSION..."

# Build for the current platform
go build $LDFLAGS -o "$BUILD_DIR/$BINARY_NAME" "$CMD_DIR"

echo "Build complete: $BUILD_DIR/$BINARY_NAME"

# Create installation directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Copy binary to installation directory
cp "$BUILD_DIR/$BINARY_NAME" "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "Installation complete: $INSTALL_DIR/$BINARY_NAME"
echo "Run '$BINARY_NAME --help' to get started"
