package scanner

type ServiceMatch struct {
	Regex   string
	Service string
	Version string
}

type Service struct {
	Name        string
	Port        int
	Protocol    string
	Probe       string
	Match       []string
	VersionInfo string
}

// ServiceSignatures Create a signature database for service identification
var ServiceSignatures = map[string][]ServiceMatch{
	"http": {
		{Regex: `Server: Apache/(\d+\.\d+\.\d+)`, Service: "Apache", Version: "$1"},
		{Regex: `Server: nginx/(\d+\.\d+\.\d+)`, Service: "Nginx", Version: "$1"},
	},
	"ssh": {
		{Regex: `SSH-2.0-OpenSSH_(\d+\.\d+)`, Service: "OpenSSH", Version: "$1"},
	},
}

// DetectService matches port detection with a known service profile
func DetectService(host string, port int) (*Service, error) {
	// TODO: Needs implementation
	return &Service{}, nil
}
