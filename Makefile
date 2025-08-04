# Network Scanner Build System

# Build variables
BINARY_NAME=scanner
BUILD_DIR=./build
CMD_DIR=./cmd/scanner
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# Go commands
GOLINT=golangci-lint
GOFMT=gofmt
GOTEST=go test
GOBUILD=go build
GOCLEAN=go clean
GORUN=go run
GOVET=go vet
GOMOD=go mod

.PHONY: all build clean run test lint fmt vet mod-tidy cross-build release help

# Default target
all: lint test build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
run:
	$(GORUN) $(CMD_DIR) --help

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run linter
lint:
	@echo "Running linter..."
	@if command -v $(GOLINT) > /dev/null; then \
		$(GOLINT) run; \
	else \
		echo "golangci-lint not installed, skipping lint"; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -w -s .

# Run vet
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Update dependencies
mod-tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

# Cross-compile for multiple platforms
cross-build:
	@echo "Cross-compiling for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	# Linux (amd64)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(CMD_DIR)
	# Windows (amd64)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(CMD_DIR)
	# macOS (amd64)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(CMD_DIR)
	# macOS (arm64)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(CMD_DIR)
	@echo "Cross-compilation complete"

# Create release using goreleaser
release:
	@echo "Creating release with goreleaser..."
	@if command -v goreleaser > /dev/null; then \
		goreleaser release --clean; \
	else \
		echo "goreleaser not installed, run: go install github.com/goreleaser/goreleaser@latest"; \
	fi

# Show help
help:
	@echo "Available targets:"
	@echo "  all          - Run lint, test, and build"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  mod-tidy     - Tidy Go modules"
	@echo "  cross-build  - Cross-compile for multiple platforms"
	@echo "  release      - Create release with goreleaser"
	@echo "  help         - Show this help"
