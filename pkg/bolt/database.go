package bolt

import (
	"github.com/boltdb/bolt"
	"time"
)

type Database struct {
	db *bolt.DB
}

type Bucket = bolt.Bucket

const (
	DB_NAME = "go_rest_api.db"
)

func NewDatabase() (*Database, error) {
	db, err := bolt.Open(DB_NAME, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (d *Database) ViewBucket(bkt_name string, fn func(*Bucket) error) error {
	return d.db.View(func(tx *bolt.Tx) error {
		return fn(tx.Bucket([]byte(bkt_name)))
	})
}

func (d *Database) UpdateBucket(bkt_name string, fn func(*Bucket) error) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists([]byte(bkt_name))
		if err != nil {
			return err
		}
		return fn(bkt)
	})
}

func (d *Database) DeleteBucket(bkt_name string) {
	d.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bkt_name))
	})
}

func (d *Database) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
