package uguis

// Defaults
const (
	defaultSimpleTwitterClientReqCBfSize = 4096
	defaultSimpleTwitterClientResCBfSize = 4096
)

// SimpleTwitterClientOptions represents options for a simple twitter client.
type SimpleTwitterClientOptions struct {
	// ReqCBfSize represents a buffer size of a request channel.
	ReqCBfSize uint32
	// ResCBfSize represents a buffer size of a response channel.
	ResCBfSize uint32
}

// setDefaults sets defaults to the simple twitter options.
func (opts *SimpleTwitterClientOptions) setDefaults() {
	if opts.ReqCBfSize == 0 {
		opts.ReqCBfSize = defaultSimpleTwitterClientReqCBfSize
	}

	if opts.ResCBfSize == 0 {
		opts.ResCBfSize = defaultSimpleTwitterClientResCBfSize
	}
}
