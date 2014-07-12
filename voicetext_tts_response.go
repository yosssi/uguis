package uguis

import "github.com/yosssi/go-voicetext"

// VoicetextTTSResponse represents a voicetext response.
type VoicetextTTSResponse struct {
	tweet  Tweet
	result *voicetext.Result
}

// NewVoicetextTTSResponse creates and returns a a voicetext response.
func NewVoicetextTTSResponse(tweet Tweet, result *voicetext.Result) VoicetextTTSResponse {
	return VoicetextTTSResponse{
		tweet:  tweet,
		result: result,
	}
}
