package uguis

// Player is an interface for playing a sound file.
type Player interface {
	// Play plays a sound file.
	Play(path string)
	// Close closes the player.
	Close() error
}
