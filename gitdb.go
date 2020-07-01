package gitdb

import (
	"github.com/boltdb/bolt"
)

type GitDB struct {
	DB *bolt.DB
}
