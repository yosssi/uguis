package uguis

// Defaults
const (
	defaultSimpleFileWriterReqCBfSize = 4096
	defaultSimpleFileWriterResCBfSize = 4096
)

// SimpleFileWriterOptions represents options for a simple file writer.
type SimpleFileWriterOptions struct {
	// ReqCBfSize represents a buffer size of a request channel.
	ReqCBfSize uint32
	// ResCBfSize represents a buffer size of a response channel.
	ResCBfSize uint32
}

// setDefaults sets defaults to the simple file writer.
func (opts *SimpleFileWriterOptions) setDefaults() {
	if opts.ReqCBfSize == 0 {
		opts.ReqCBfSize = defaultSimpleFileWriterReqCBfSize
	}

	if opts.ResCBfSize == 0 {
		opts.ResCBfSize = defaultSimpleFileWriterResCBfSize
	}
}
