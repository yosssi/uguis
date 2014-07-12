package uguis

// simpleDBReadWriter represents a simple database read writer.
type simpleDBReadWriter struct {
	db DB
}

// Get gets a value from a database.
func (dbRW *simpleDBReadWriter) Get(bucketName, key []byte) ([]byte, error) {
	return dbRW.db.Get(bucketName, key)
}

// Put puts the key/value to a database.
func (dbRW *simpleDBReadWriter) Put(bucketName, key, value []byte) error {
	return dbRW.db.Put(bucketName, key, value)
}

// Close does nothing.
func (dbRW *simpleDBReadWriter) Close() error {
	return nil
}

// NewSimpleDBReadWriter creates and returns a simple database read writer.
func NewSimpleDBReadWriter(db DB, opts *SimpleDBReadWriterOptions) DBReadWriter {
	// Initialize options.
	if opts == nil {
		opts = &SimpleDBReadWriterOptions{}
	}
	opts.setDefaults()

	return &simpleDBReadWriter{
		db: db,
	}
}
