package uguis

import (
	"os"
	"runtime"
)

const fallbackHostname = "Uguis"

// ApplicationOptions represents options for an application.
type ApplicationOptions struct {
	// Hostname represents the host name.
	Hostname string
	// CPUs represents the maximum number of CPUs.
	CPUs int
}

// setDefaults sets defaults to the application options.
func (opts *ApplicationOptions) setDefaults() {
	if opts.Hostname == "" {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = fallbackHostname
		}
		opts.Hostname = hostname
	}

	if opts.CPUs == 0 || opts.CPUs > runtime.NumCPU() {
		opts.CPUs = runtime.NumCPU()
	}
}
