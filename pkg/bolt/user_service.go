package bolt

import (
	"encoding/json"
	"errors"
	"go_rest_api/pkg"
)

type UserService struct {
	session  *Session
	bkt_name string
	hash     root.Hash
}

func NewUserService(s *Session, bkt_name string, hash root.Hash) *UserService {
	return &UserService{s, bkt_name, hash}
}

func (us *UserService) Create(u *root.User) (err error) {
	err = us.session.UpdateBucket(us.bkt_name, func(bkt *Bucket) error {
		u.Password, err = us.hash.Generate(u.Password)
		if err != nil {
			return errors.New("Could not hash password.")
		}

		if buf := bkt.Get([]byte(u.Username)); buf != nil {
			return errors.New("User already exists.")
		}
		if buf, err := json.Marshal(&u); err != nil {
			return err
		} else if err := bkt.Put([]byte(u.Username), buf); err != nil {
			return err
		}
		return nil
	})
	return
}

func (us *UserService) GetByUsername(username string) (ru *root.User, err error) {
	err = us.session.ViewBucket(us.bkt_name, func(bkt *Bucket) error {
		if bkt == nil {
			return errors.New("User does not exist.")
		}

		if buf := bkt.Get([]byte(username)); buf == nil {
			return errors.New("User does not exist.")
		} else if err := json.Unmarshal(buf, &ru); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
