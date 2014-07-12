package uguis

// Defaults
const (
	defaultSimpleVoicetextClientReqCBfSize = 4096
	defaultSimpleVoicetextClientResCBfSize = 4096
)

// SimpleVoicetextClientOptions represents options for a simple voicetext client.
type SimpleVoicetextClientOptions struct {
	// ReqCBfSize represents a buffer size of a request channel.
	ReqCBfSize uint32
	// ResCBfSize represents a buffer size of a response channel.
	ResCBfSize uint32
}

// setDefaults sets defaults to the simple voicetext options.
func (opts *SimpleVoicetextClientOptions) setDefaults() {
	if opts.ReqCBfSize == 0 {
		opts.ReqCBfSize = defaultSimpleVoicetextClientReqCBfSize
	}

	if opts.ResCBfSize == 0 {
		opts.ResCBfSize = defaultSimpleVoicetextClientResCBfSize
	}
}
