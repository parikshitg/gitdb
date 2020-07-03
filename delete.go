package gitdb

import (
	"github.com/boltdb/bolt"
)

func (g *GitDB) Delete(bucket string, id string) error {

	err := g.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(id))
	})
	if err != nil {
		return err
	}

	return err
}
