package gitdb

import (
	"github.com/boltdb/bolt"
)

type GitDB struct {
	DB *bolt.DB
}

type Content interface {
	GetHeader() *Header
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	Copy() interface{}
}

type Header struct {
	ID string `json:"id"`
}
