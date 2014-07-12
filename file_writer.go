package uguis

// FileWriter is an interface for writing a file.
type FileWriter interface {
	// Write writes a file.
	Write(req FileWriterRequest)
	// Close closes the file writer.
	Close() error
	// ResC returns a response channel.
	ResC() <-chan FileWriterResponse
}
