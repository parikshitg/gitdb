package gitdb

import (
	"errors"

	"github.com/boltdb/bolt"
)

func (g *GitDB) Read(bucket string, id string, content Content) error {

	var value []byte

	err := g.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))
		value = b.Get([]byte(id))

		return nil
	})
	if err != nil {
		return err
	}

	if value == nil {
		return errors.New("Not Found")
	}

	err = content.Deserialize(value)
	return err
}
