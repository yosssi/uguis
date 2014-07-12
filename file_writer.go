package uguis

// FileWriter is an interface for writing a file.
type FileWriter interface {
	// Write writes a file.
	Write(f file)
	// Close closes the file writer.
	Close() error
}
