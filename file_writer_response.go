package uguis

// FileWriterResponse represents a file writer response.
type FileWriterResponse struct {
	tweet Tweet
	path  string
}

// NewFileWriterResponse creates and returns a file writer response.
func NewFileWriterResponse(tweet Tweet, path string) FileWriterResponse {
	return FileWriterResponse{
		tweet: tweet,
		path:  path,
	}
}
