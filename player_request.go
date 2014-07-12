package uguis

// PlayerRequest represents a player request.
type PlayerRequest struct {
	tweet Tweet
	path  string
}

// NewPlayerRequest creates and returns a player request.
func NewPlayerRequest(tweet Tweet, path string) PlayerRequest {
	return PlayerRequest{
		tweet: tweet,
		path:  path,
	}
}
