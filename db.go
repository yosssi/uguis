package uguis

// DB is an interface for storing data.
type DB interface {
	// Close closes the database.
	Close() error
}
