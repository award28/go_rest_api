package bolt

import (
	"github.com/boltdb/bolt"
	"time"
)

type Session struct {
	session *bolt.DB
}

type Bucket = bolt.Bucket

const (
	DB_NAME = "go_rest_api.db"
)

func NewSession() (*Session, error) {
	db, err := bolt.Open(DB_NAME, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return &Session{db}, nil
}

func (s *Session) ViewBucket(bkt_name string, fn func(*Bucket) error) error {
	return s.session.View(func(tx *bolt.Tx) error {
		return fn(tx.Bucket([]byte(bkt_name)))
	})

}
func (s *Session) UpdateBucket(bkt_name string, fn func(*Bucket) error) error {
	return s.session.Update(func(tx *bolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists([]byte(bkt_name))
		if err != nil {
			return err
		}
		return fn(bkt)
	})
}

func (s *Session) DeleteBucket(bkt_name string) {
	s.session.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bkt_name))
	})
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
