package uguis

// Defaults
const (
	defaultSimplePlayerReqCBfSize = 4096
	defaultSimplePlayerResCBfSize = 4096
)

// SimplePlayerOptions represents options for a simple player.
type SimplePlayerOptions struct {
	// ReqCBfSize represents a buffer size of a request channel.
	ReqCBfSize uint32
	// ResCBfSize represents a buffer size of a response channel.
	ResCBfSize uint32
}

// setDefaults sets defaults to the simple player.
func (opts *SimplePlayerOptions) setDefaults() {
	if opts.ReqCBfSize == 0 {
		opts.ReqCBfSize = defaultSimplePlayerReqCBfSize
	}

	if opts.ResCBfSize == 0 {
		opts.ResCBfSize = defaultSimplePlayerResCBfSize
	}
}
