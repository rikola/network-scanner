package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"network-scanner/internal/scanner"
)

// version is set during build using ldflags
var version = "dev"

// Target usage examples:
// ./scanner -host 192.168.1.1 -ports 1-1000
// ./scanner -host 192.168.1.1 -ports 22,80,443,8080
// ./scanner -range 192.168.1.0/24 -ports 22,80,443

var (
	// Command line flags
	host         string
	portRangeStr string
	timeoutStr   string
	concurrent   int
	verbose      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scanner",
	Short: "A fast TCP port scanner",
	Long: `Network Scanner is a comprehensive network scanning tool built in Go.

It can discover hosts, scan ports, and identify services to help with network reconnaissance.

Examples:
  scanner --host 192.168.1.1 --ports 1-1000
  scanner --host example.com --ports 22,80,443,8080
  scanner --host 192.168.1.1 --ports 1-100 --timeout 5s --concurrent 50`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate required flags
		if host == "" {
			fmt.Println("Error: host is required")
			cmd.Help()
			os.Exit(1)
		}

		if portRangeStr == "" {
			fmt.Println("Error: ports is required")
			cmd.Help()
			os.Exit(1)
		}

		// Parse ports
		ports, err := parsePorts(portRangeStr)
		if err != nil {
			fmt.Printf("Error parsing ports: %v\n", err)
			os.Exit(1)
		}

		// Parse timeout
		timeout, err := time.ParseDuration(timeoutStr)
		if err != nil {
			fmt.Printf("Error parsing timeout: %v\n", err)
			os.Exit(1)
		}

		// Check if host is alive
		if verbose {
			fmt.Printf("Checking if host %s is alive...\n", host)
		}

		if scanner.IsHostAlive(host) {
			fmt.Printf("Host %s is up\n", host)

			// Execute port scan based on the provided range
			var openPorts []int

			// If there's a hyphen, it's a range scan
			if strings.Contains(portRangeStr, "-") && len(ports) >= 2 {
				startPort := ports[0]
				endPort := ports[len(ports)-1]

				if verbose {
					fmt.Printf("Scanning port range %d-%d on %s with %d concurrent scanners\n",
						startPort, endPort, host, concurrent)
				}

				// Set the concurrent scanner count
				scanner.SetConcurrentScanners(concurrent)

				// Scan port range
				openPorts = scanner.ScanPortRange(host, startPort, endPort)
			} else {
				// Scan individual ports
				if verbose {
					fmt.Printf("Scanning %d individual ports on %s\n", len(ports), host)
				}

				for _, port := range ports {
					if verbose {
						fmt.Printf("Scanning port %d...\n", port)
					}

					if scanner.ScanTCPPort(host, port, timeout) {
						openPorts = append(openPorts, port)
					}
				}
			}

			// Print results
			if len(openPorts) > 0 {
				fmt.Println("\nOpen ports:")
				for _, port := range openPorts {
					fmt.Printf("  %d\n", port)
				}
			} else {
				fmt.Println("\nNo open ports found")
			}
		} else {
			fmt.Printf("Host %s appears to be down\n", host)
		}
	},
}

// parsePorts handles both individual ports (80,443) and port ranges (1-1000)
func parsePorts(portsFlag string) ([]int, error) {
	var ports []int

	// Handle ranges like "1-1000"
	if strings.Contains(portsFlag, "-") {
		parts := strings.Split(portsFlag, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid port range format")
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start port: %v", err)
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid end port: %v", err)
		}

		if start > end {
			return nil, fmt.Errorf("start port cannot be greater than end port")
		}

		if start < 1 || end > 65535 {
			return nil, fmt.Errorf("port numbers must be between 1 and 65535")
		}

		// Just return the start and end for range scanning
		return []int{start, end}, nil
	}

	// Handle comma-separated individual ports
	portsList := strings.Split(portsFlag, ",")
	for _, portStr := range portsList {
		portStr = strings.TrimSpace(portStr)
		if portStr == "" {
			continue
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid port number '%s': %v", portStr, err)
		}

		if port < 1 || port > 65535 {
			return nil, fmt.Errorf("port number %d out of range (1-65535)", port)
		}

		ports = append(ports, port)
	}

	return ports, nil
}

func init() {
	// Define command-line flags
	rootCmd.Flags().StringVarP(&host, "host", "", "", "Target host to scan (required)")
	rootCmd.Flags().StringVarP(&portRangeStr, "ports", "p", "", "Ports to scan (e.g., '80,443' or '1-1000') (required)")
	rootCmd.Flags().StringVarP(&timeoutStr, "timeout", "t", "2s", "Timeout for each port scan (e.g., '500ms', '2s')")
	rootCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 20, "Number of concurrent scans")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Mark required flags
	rootCmd.MarkFlagRequired("host")
	rootCmd.MarkFlagRequired("ports")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mainold() {
	host := "example.com"

	// Check if host is alive
	if scanner.IsHostAlive(host) {
		fmt.Printf("Host %s is up\n", host)

		if scanner.ScanTCPPort(host, 80, 2*time.Second) {
			fmt.Println("Port 80 is open")
		}

		// Scan port range
		openPorts := scanner.ScanPortRange(host, 1, 1024)
		fmt.Printf("Open ports: %v\n", openPorts)
	} else {
		fmt.Printf("Host %s appears to be down\n", host)
	}
}
