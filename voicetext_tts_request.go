package uguis

import "github.com/yosssi/go-voicetext"

// VoicetextTTSRequest represents a voicetext request.
type VoicetextTTSRequest struct {
	text string
	opts *voicetext.TTSOptions
}

// NewVoicetextTTSRequest creates and returns a a voicetext request.
func NewVoicetextTTSRequest(text string, opts *voicetext.TTSOptions) VoicetextTTSRequest {
	return VoicetextTTSRequest{
		text: text,
		opts: opts,
	}
}
