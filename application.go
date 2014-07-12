package uguis

// Application represents an application.
type Application struct {
	// Hostname represents the hostname.
	Hostname string
	// CPUs represents the maximum number of CPUs.
	CPUs int
}

// NewApplication creates and returns an application.
func NewApplication(opts *ApplicationOptions) *Application {
	// Initialize options.
	if opts == nil {
		opts = &ApplicationOptions{}
	}
	opts.setDefaults()

	return &Application{
		Hostname: opts.Hostname,
		CPUs:     opts.CPUs,
	}
}
