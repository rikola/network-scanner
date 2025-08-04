# Network Scanner

A comprehensive network scanner tool built in Go that can discover hosts, scan ports, and identify services.

## Features

- TCP connect scanning for port discovery
- Host discovery through TCP and ICMP
- Concurrent scanning for performance
- Configurable timeout and worker count

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/network-scanner.git
cd network-scanner

# Build the application
go build -o scanner ./cmd/scanner
```

## Usage

```bash
# Basic port scanning
./scanner --host example.com --ports 80,443

# Scan a range of ports
./scanner --host 192.168.1.1 --ports 1-1000

# Customize scan parameters
./scanner --host example.com --ports 1-100 --timeout 5s --concurrent 50 --verbose
```

### Command-line Options

- `--host`: Target host to scan (required)
- `--ports`, `-p`: Ports to scan (e.g., '80,443' or '1-1000') (required)
- `--timeout`, `-t`: Timeout for each port scan (default: 2s)
- `--concurrent`, `-c`: Number of concurrent scans (default: 20)
- `--verbose`, `-v`: Enable verbose output

## Examples

### Scan Common Web Ports

```bash
./scanner --host example.com --ports 80,443,8080,8443
```

### Comprehensive Scan

```bash
./scanner --host 192.168.1.1 --ports 1-10000 --timeout 1s --concurrent 100 --verbose
```

## License

MIT