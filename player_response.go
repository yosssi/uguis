package uguis

// PlayerResponse represents a player response.
type PlayerResponse struct {
	tweet Tweet
	path  string
}

// NewPlayerResponse creates and returns a player response.
func NewPlayerResponse(tweet Tweet, path string) PlayerResponse {
	return PlayerResponse{
		tweet: tweet,
		path:  path,
	}
}
