package uguis

// VoicetextClient is an interface for calling Voicetext Web APIs.
type VoicetextClient interface {
	// TTS calls the Voicetext TTS API.
	TTS(req VoicetextTTSRequest)
	// Close closes the twitter client.
	Close() error
}
