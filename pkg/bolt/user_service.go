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

var doesNotExistErr = errors.New("User does not exist.")

func NewUserService(s *Session, bkt_name string, hash root.Hash) *UserService {
	return &UserService{s, bkt_name, hash}
}

func (us *UserService) Login(c *root.Credentials) (ru *root.User, err error) {
	// Verify credentials

	// Compare credentials to stored credentials

	// Login User

	return nil, nil
}

func (us *UserService) Signup(nu *root.NewUser) (ru *root.User, err error) {
	// Verify new user details
	if err := nonEmptyFields_NewUser(*nu); err != nil {
		return nil, err
	}
	if err := verifyFields_NewUser(*nu); err != nil {
		return nil, err
	}

	// Check if user with username exists
	if _, err := us.GetByUsername(nu.Username); err != doesNotExistErr {
		return nil, errors.New("Username is taken.")
	}

	// Hash Password
	nu.Password, err = us.hash.Generate(nu.Password)
	if err != nil {
		return nil, errors.New("Could not hash password.")
	}

	// Create user
	u := &root.User{
		Username: nu.Username,
		Password: nu.Password,
		Email:    nu.Email,
	}
	if err = us.Create(u); err != nil {
		return nil, err
	}

	// Login User
	return u, nil
}

func (us *UserService) GetByUsername(username string) (ru *root.User, err error) {
	err = us.session.ViewBucket(us.bkt_name, func(bkt *Bucket) error {
		if bkt == nil {
			return doesNotExistErr
		} else if buf := bkt.Get([]byte(username)); buf == nil {
			return doesNotExistErr
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

func (us *UserService) Create(u *root.User) (err error) {
	err = us.session.UpdateBucket(us.bkt_name, func(bkt *Bucket) error {
		if buf, err := json.Marshal(&u); err != nil {
			return err
		} else if err := bkt.Put([]byte(u.Username), buf); err != nil {
			return err
		}
		return nil
	})
	return
}
