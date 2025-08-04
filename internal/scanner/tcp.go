package scanner

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

// Default number of concurrent scanners
var numWorkers = 20

// SetConcurrentScanners sets the number of concurrent port scanners
func SetConcurrentScanners(workers int) {
	if workers > 0 {
		numWorkers = workers
	}
}

// ScanTCPPort attempts to connect to a specific TCP port on a host
// Returns true if the port is open, false otherwise
func ScanTCPPort(host string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// ScanPortRange scans a range of ports on a host and returns a slice of open ports
// Uses goroutines for concurrent scanning with a worker pool pattern
func ScanPortRange(host string, startPort, endPort int) []int {
	var openPorts []int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// Create a buffered chanel for jobs to control concurrency
	jobs := make(chan int, 100)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range jobs {
				if ScanTCPPort(host, port, 2*time.Second) {
					mutex.Lock()
					openPorts = append(openPorts, port)
					mutex.Unlock()
				}
			}
		}()
	}

	// Send jobs to workers
	for port := startPort; port <= endPort; port++ {
		jobs <- port
	}
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()

	// Sorts the results for consistent output
	sort.Ints(openPorts)
	return openPorts
}

// IsHostAlive checks if a host is reachable by attempting to connect
// to common ports (80, 443, 22) with a short timeout
func IsHostAlive(host string) bool {
	commonPorts := []int{80, 443, 22}
	timeout := 1 * time.Second

	for _, port := range commonPorts {
		if ScanTCPPort(host, port, timeout) {
			return true
		}
	}

	// Try ICMP ping as a fallback
	return pingHost(host)
}

func pingHost(host string) bool {
	addr, err := net.ResolveIPAddr("ip4:icmp", host)
	if err != nil {
		return false
	}

	conn, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
