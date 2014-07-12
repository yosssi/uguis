package uguis

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	defaultDBDir      = ".uguis"
	defaultDBFileName = "uguis.db"
)

// BoltDBOptions represents options for a Bolt database.
type BoltDBOptions struct {
	// Path represents a Bolt database file path.
	Path string
}

// setDefaults sets defaults to the Bolt database options.
func (opts *BoltDBOptions) setDefaults() error {
	if opts.Path == "" {
		if usr, err := user.Current(); err == nil {
			dir := filepath.Join(usr.HomeDir, defaultDBDir)

			// Check if the target directory exists or not.
			if _, err := os.Stat(dir); err != nil {
				if os.IsNotExist(err) {
					// Make the target directory if it does not exist.
					if err := os.Mkdir(dir, os.ModePerm); err != nil {
						return err
					}
				} else {
					return err
				}
			}

			opts.Path = filepath.Join(dir, defaultDBFileName)
		} else {
			opts.Path = defaultDBFileName
		}
	}

	return nil
}
