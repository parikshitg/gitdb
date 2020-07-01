package gitdb

import (
	"github.com/boltdb/bolt"
)

func Open(path string) (*GitDB, error) {
	var err error

	var g = &GitDB{}

	g.DB, err = bolt.Open(path, 0644, nil)
	return g, error
}
