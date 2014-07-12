package uguis

// TwitterClient is an interface for calling the Twitter APIs.
type TwitterClient interface {
	// Call calls a Twitter API.
	Call(req twitterRequest)
	// Close closes the twitter client.
	Close() error
}
