#!/bin/bash

# Test script for Network Scanner
set -e

echo "Running tests..."

# Run all tests with verbose output and coverage
go test -v -cover ./...

echo "Tests completed successfully"
