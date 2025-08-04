# Network Scanner Tool - Complete Project Guide

## Project Overview

Build a comprehensive network scanner in Go that can discover hosts, scan ports, identify services, and perform basic reconnaissance. This project will teach you fundamental network security concepts while leveraging your Go expertise.

## Learning Objectives

- Understand TCP/UDP protocols at a low level
- Learn network reconnaissance techniques
- Practice concurrent programming for performance
- Implement service detection and banner grabbing
- Build modular, extensible security tools

## Phase 1: Basic Port Scanner (Week 1)

### Core Functionality
- TCP connect scans
- Command-line interface
- Basic host discovery
- Simple port range scanning

### Implementation Steps

#### 1. Project Structure
```
network-scanner/
├── cmd/
│   └── scanner/
│       └── main.go
├── internal/
│   ├── scanner/
│   │   ├── tcp.go
│   │   ├── host.go
│   │   └── types.go
│   └── utils/
│       └── network.go
├── pkg/
│   └── ports/
│       └── common.go
├── go.mod
└── README.md
```

#### 2. Basic TCP Scanner Implementation
```go
// Key functions to implement:
func ScanTCPPort(host string, port int, timeout time.Duration) bool
func ScanPortRange(host string, startPort, endPort int) []int
func IsHostAlive(host string) bool
```

#### 3. Command Line Interface
```bash
# Target usage examples:
./scanner -host 192.168.1.1 -ports 1-1000
./scanner -host 192.168.1.1 -ports 22,80,443,8080
./scanner -range 192.168.1.0/24 -ports 22,80,443
```

### Technical Challenges
1. **Socket Programming**: Direct TCP connection handling
2. **Timeout Management**: Preventing hung connections
3. **Error Handling**: Distinguishing between filtered/closed ports
4. **IP Address Parsing**: CIDR notation and range handling

### Success Metrics
- Scan 1000 ports in under 30 seconds
- Accurately detect open/closed/filtered states
- Handle network errors gracefully
- Clean, readable output format

## Phase 2: Performance Optimization (Week 2)

### Advanced Features
- Concurrent scanning with goroutines
- Rate limiting to avoid overwhelming targets
- SYN scan implementation (raw sockets)
- UDP scanning capabilities

### Implementation Focus

#### 1. Concurrency Patterns
```go
// Worker pool pattern for controlled concurrency
type ScanJob struct {
    Host string
    Port int
}

type ScanResult struct {
    Host   string
    Port   int
    Open   bool
    Error  error
}

func WorkerPool(jobs <-chan ScanJob, results chan<- ScanResult, workers int)
```

#### 2. Rate Limiting
```go
// Implement token bucket or sliding window
type RateLimiter struct {
    tokens chan struct{}
    ticker *time.Ticker
}
```

#### 3. Raw Socket Implementation (Linux/macOS)
```go
// SYN scan for stealth and performance
func SYNScan(host string, port int) (bool, error) {
    // Requires root privileges
    // Use golang.org/x/net/ipv4 and raw sockets
}
```

### Performance Targets
- 10,000+ ports per minute with proper concurrency
- Configurable thread count (default: 100-500)
- Memory usage under 50MB for large scans
- Graceful handling of network congestion

## Phase 3: Service Detection & Banner Grabbing (Week 3)

### Enhanced Capabilities
- Service fingerprinting
- Banner grabbing for common services
- Version detection
- Service-specific probes

### Implementation Details

#### 1. Service Detection Engine
```go
type Service struct {
    Name        string
    Port        int
    Protocol    string
    Probe       string
    Match       []string
    VersionInfo string
}

func DetectService(host string, port int) (*Service, error)
```

#### 2. Common Service Probes
- **HTTP/HTTPS**: GET / HTTP/1.1 requests
- **SSH**: Protocol version exchange
- **FTP**: Banner reading
- **SMTP**: HELO command
- **Telnet**: Initial banner
- **DNS**: Query response analysis

#### 3. Banner Database
```go
// Create signature database for service identification
var ServiceSignatures = map[string][]ServiceMatch{
    "http": {
        {Regex: `Server: Apache/(\d+\.\d+\.\d+)`, Service: "Apache", Version: "$1"},
        {Regex: `Server: nginx/(\d+\.\d+\.\d+)`, Service: "Nginx", Version: "$1"},
    },
    "ssh": {
        {Regex: `SSH-2.0-OpenSSH_(\d+\.\d+)`, Service: "OpenSSH", Version: "$1"},
    },
}
```

### Advanced Features
- Custom probe definitions
- Aggressive vs passive scanning modes
- SSL/TLS certificate analysis
- HTTP title and header extraction

## Phase 4: Host Discovery & Network Mapping (Week 4)

### Network Discovery Features
- ICMP ping sweeps
- ARP scanning (local network)
- TCP/UDP ping alternatives
- Subnet enumeration
- MAC address resolution

### Implementation Components

#### 1. ICMP Implementation
```go
// Requires raw sockets or external ping
func ICMPPing(host string, timeout time.Duration) (bool, time.Duration, error)
```

#### 2. ARP Scanning
```go
// For local network discovery
func ARPScan(network string) ([]Host, error) {
    // Use raw sockets or system ARP table
}
```

#### 3. Alternative Host Discovery
```go
// TCP SYN to common ports
func TCPPing(host string, ports []int) bool

// UDP probe to common services
func UDPPing(host string, ports []int) bool
```

### Network Mapping Features
- Traceroute implementation
- TTL analysis for OS detection
- Network topology mapping
- Subnet relationship detection

## Advanced Extensions

### 1. OS Fingerprinting
- TCP/IP stack fingerprinting
- TTL analysis
- Window size detection
- TCP options analysis

### 2. Vulnerability Detection
- Known vulnerable service versions
- Default credential testing
- Common misconfigurations
- CVE database integration

### 3. Evasion Techniques
- Randomized scan timing
- Source port randomization
- Decoy scanning
- Fragment scanning

### 4. Output Formats
- JSON export for automation
- XML (Nmap compatibility)
- CSV for spreadsheet analysis
- HTML reports with graphs

## Code Quality & Security

### Best Practices
- Input validation and sanitization
- Proper error handling and logging
- Configuration file support
- Signal handling for graceful shutdown
- Resource cleanup (file handles, goroutines)

### Security Considerations
- Rate limiting to avoid DoS
- Privilege escalation handling
- Network interface selection
- Firewall interaction awareness

## Testing Strategy

### Unit Tests
- Port scanning accuracy
- Host discovery reliability
- Service detection precision
- Performance benchmarks

### Integration Tests
- Full subnet scans
- Service enumeration accuracy
- Output format validation
- Error condition handling

### Real-World Testing
1. **Home Network**: Start with your own infrastructure
2. **VulnHub VMs**: Test against known vulnerable systems
3. **Docker Containers**: Controlled test environments
4. **Cloud Instances**: Your own AWS/GCP instances

## Deployment & Distribution

### Build Pipeline
```bash
# Cross-compilation for multiple platforms
GOOS=linux GOARCH=amd64 go build -o scanner-linux-x64
GOOS=windows GOARCH=amd64 go build -o scanner-windows-x64.exe
GOOS=darwin GOARCH=amd64 go build -o scanner-macos-x64
```

### Installation Methods
- Pre-compiled binaries
- Docker container
- Go module installation
- Package manager distribution (brew, apt)

## Success Criteria

By the end of this project, your scanner should:
- ✅ Discover live hosts on a network segment
- ✅ Accurately identify open ports and services
- ✅ Perform banner grabbing and version detection
- ✅ Complete subnet scans in reasonable time
- ✅ Generate useful reports for further analysis
- ✅ Handle various network conditions gracefully

## Next Steps Integration

This scanner becomes the foundation for:
- **Vulnerability Assessment**: Feed results into vuln scanners
- **Network Monitoring**: Baseline for change detection
- **Red Team Operations**: Reconnaissance phase tooling
- **Blue Team Defense**: Asset discovery and monitoring

## Recommended Timeline

- **Week 1**: Basic TCP scanner with CLI
- **Week 2**: Concurrency and performance optimization
- **Week 3**: Service detection and banner grabbing
- **Week 4**: Host discovery and advanced features

Each week should include both implementation and testing against real networks (your own lab environment).