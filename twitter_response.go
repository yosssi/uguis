package uguis

// TwitterResponse in an interface for returning tweets.
type TwitterResponse interface {
	Tweets() []Tweet
}
