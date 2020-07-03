package gitdb

import (
	"errors"

	"github.com/boltdb/bolt"
)

func (g *GitDB) Update(bucket string, id string, content Content) error {

	err := g.DB.Update(func(tx *bolt.Tx) error {

		data, err := content.Serialize()
		if err != nil {
			return err
		}

		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(id))
		if value == nil {
			return errors.New("Key not found!")
		}

		err = b.Put([]byte(id), data)
		return err

	})

	return err
}
