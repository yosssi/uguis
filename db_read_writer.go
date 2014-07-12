package uguis

// DBReadWriter is an interface for reading/writing data from/to the databse.
type DBReadWriter interface {
	// Get gets a value from a database.
	Get(bucketName, key []byte) ([]byte, error)
	// Put puts the key/value to a database.
	Put(bucketName, key, value []byte) error
	// Close closes the database read writer's process.
	Close() error
}
