package uguis

import "github.com/yosssi/go-voicetext"

// VoicetextTTSRequest represents a voicetext request.
type VoicetextTTSRequest struct {
	tweet Tweet
	opts  *voicetext.TTSOptions
}

// NewVoicetextTTSRequest creates and returns a voicetext request.
func NewVoicetextTTSRequest(tweet Tweet, opts *voicetext.TTSOptions) VoicetextTTSRequest {
	return VoicetextTTSRequest{
		tweet: tweet,
		opts:  opts,
	}
}
