package gitdb

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

func (g *GitDB) Create(bucket string, content interface{}) (string, error) {

	id := uuid.New().String()

	err := g.DB.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		// json encoding
		data, err := json.Marshal(content)
		if err != nil {
			return err
		}

		err = b.Put([]byte(id), data)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return id, nil
}
