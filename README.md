# Network Scanner

A comprehensive network scanner tool built in Go that can discover hosts, scan ports, and identify services.

[![Build](https://github.com/rikola/network-scanner/actions/workflows/build.yml/badge.svg)](https://github.com/rikola/network-scanner/actions/workflows/build.yml)

## Features

- TCP connect scanning for port discovery
- Host discovery through TCP and ICMP
- Concurrent scanning for performance
- Configurable timeout and worker count
- Cross-platform support (Linux, macOS, Windows)

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/network-scanner.git
cd network-scanner

# Using Make
make build

# Or using the build script
chmod +x scripts/build.sh
./scripts/build.sh

# Install to /usr/local/bin (requires sudo)
sudo ./scripts/install.sh
```

### Using Docker

```bash
# Build and run with Docker
docker build -t network-scanner .
docker run --cap-add=NET_ADMIN --cap-add=NET_RAW network-scanner --host example.com --ports 80,443

# Or using Docker Compose
docker-compose up
```

## Usage

```bash
# Basic port scanning
scanner --host example.com --ports 80,443

# Scan a range of ports
scanner --host 192.168.1.1 --ports 1-1000

# Customize scan parameters
scanner --host example.com --ports 1-100 --timeout 5s --concurrent 50 --verbose
```

### Command-line Options

- `--host`: Target host to scan (required)
- `--ports`, `-p`: Ports to scan (e.g., '80,443' or '1-1000') (required)
- `--timeout`, `-t`: Timeout for each port scan (default: 2 s)
- `--concurrent`, `-c`: Number of concurrent scans (default: 20)
- `--verbose`, `-v`: Enable verbose output

## Examples

### Scan Common Web Ports

```bash
scanner --host example.com --ports 80,443,8080,8443
```

### Comprehensive Scan

```bash
scanner --host 192.168.1.1 --ports 1-10000 --timeout 1s --concurrent 100 --verbose
```

## Development

### Build System

This project uses Make as its primary build tool. Here are the available commands:

```bash
# Build the application
make build

# Run tests
make test

# Lint the code
make lint

# Format the code
make fmt

# Clean build artifacts
make clean

# Cross-compile for multiple platforms
make cross-build

# Create a release with GoReleaser
make release
```

### Releases

To create a new release:

1. Tag the release: `git tag -a v1.0.0 -m "Release v1.0.0"`
2. Push the tag: `git push origin v1.0.0`
3. The GitHub Actions workflow will automatically build and publish the release

Alternatively, you can use GoReleaser locally:

```bash
goreleaser release --clean
```

## Docker

To build and run with Docker:

```bash
# Build the Docker image
docker build -t network-scanner .

# Run the scanner with Docker
docker run --cap-add=NET_ADMIN --cap-add=NET_RAW network-scanner --host example.com --ports 80,443
```

## License

MIT
