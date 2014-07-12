package uguis

import "github.com/boltdb/bolt"

// boltDB represents a Bolt database.
type boltDB struct {
	*bolt.DB
}

// Get gets a value from a database.
func (db *boltDB) Get(bucketName, key []byte) ([]byte, error) {
	var value []byte

	err := db.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)

		if bucket == nil {
			return nil
		}

		value = bucket.Get(key)

		return nil
	})

	return value, err
}

// Put puts the key/value to a database.
func (db *boltDB) Put(bucketName, key, value []byte) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)

		if bucket == nil {
			var err error

			bucket, err = tx.CreateBucket(bucketName)
			if err != nil {
				return nil
			}
		}

		return bucket.Put(key, value)
	})

	return err
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
