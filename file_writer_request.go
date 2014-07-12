package uguis

// FileWriterRequest represents a file writer request.
type FileWriterRequest struct {
	tweet Tweet
	file  file
}

// NewFileWriterRequest creates and returns a file writer request.
func NewFileWriterRequest(tweet Tweet, file file) FileWriterRequest {
	return FileWriterRequest{
		tweet: tweet,
		file:  file,
	}
}
