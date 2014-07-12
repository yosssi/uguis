package uguis

import "github.com/boltdb/bolt"

// boltDB represents a Bolt database.
type boltDB struct {
	*bolt.DB
}

// Close closes the Bolt database.
func (db *boltDB) Close() error {
	return db.DB.Close()
}

// NewBoltDB creates and returns a Bolt database.
func NewBoltDB(opts *BoltDBOptions) (DB, error) {
	// Initialize options.
	if opts == nil {
		opts = &BoltDBOptions{}
	}
	if err := opts.setDefaults(); err != nil {
		return nil, err
	}

	// Open a Bolt database.
	bDB, err := bolt.Open(opts.Path, 0666, nil)
	if err != nil {
		return nil, err
	}

	// Create a Bolt database.
	db := &boltDB{
		DB: bDB,
	}

	return db, nil
}
