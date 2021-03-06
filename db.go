package uguis

// DB is an interface for storing data.
type DB interface {
	// Get gets a value from a database.
	Get(bucketName, key []byte) ([]byte, error)
	// Put puts the key/value to a database.
	Put(bucketName, key, value []byte) error
	// Close closes the database.
	Close() error
}
