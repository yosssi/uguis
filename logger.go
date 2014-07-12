package uguis

// Logger is an interface for prining a log.
type Logger interface {
	// Print prints a log.
	Print(lg Log)
	// Close closes the logger.
	Close() error
}
