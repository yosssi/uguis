package uguis

// Player is an interface for playing a sound file.
type Player interface {
	// Play plays a sound file.
	Play(path PlayerRequest)
	// Close closes the player.
	Close() error
	// ResC returns a response channel.
	ResC() <-chan PlayerResponse
}
