package bolt

import (
	"encoding/json"
	"errors"
	"go_rest_api/internal"
	"net/http"
)

type UserService struct {
	db       *Database
	bkt_name string
	hash     root.Hash
}

var (
	doesNotExistErr     = errors.New("User does not exist.")
	usernameConflictErr = errors.New("Username is taken.")
	notHashableErr      = errors.New("Could not hash password.")
	credentialsErr      = errors.New("Incorrect Credentials.")
)

func NewUserService(db *Database, bkt_name string, hash root.Hash) *UserService {
	return &UserService{db, bkt_name, hash}
}

// TODO
func (us *UserService) Me() (*root.User, error) {
	return nil, nil
}

func (us *UserService) Login(c *root.Credentials) (ru *root.User, err error) {
	// Verify credentials
	if err := nonEmptyFields_Credentials(*c); err != nil {
		return nil, root.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	// Compare credentials to stored credentials
	if user, err := us.GetByUsername(c.Username); err != nil {
		if err == doesNotExistErr {
			return nil, root.StatusError{
				Code: http.StatusNotFound,
				Err:  err,
			}
		} else {
			return nil, err
		}
	} else if err = us.hash.Compare(user.Password, c.Password); err != nil {
		return nil, root.StatusError{
			Code: http.StatusUnauthorized,
			Err:  credentialsErr,
		}
	} else {
		user.Password = ""
		return user, nil
	}
}

func (us *UserService) Signup(nu *root.NewUser) (ru *root.User, err error) {
	// Verify new user details
	if err := nonEmptyFields_NewUser(*nu); err != nil {
		return nil, root.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	if err := verifyFields_NewUser(*nu); err != nil {
		return nil, root.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	// Check if user with username exists
	if _, err := us.GetByUsername(nu.Username); err != doesNotExistErr {
		return nil, root.StatusError{
			Code: http.StatusConflict,
			Err:  usernameConflictErr,
		}
	}

	// Hash Password
	nu.Password, err = us.hash.Generate(nu.Password)
	if err != nil {
		return nil, notHashableErr
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
	err = us.db.ViewBucket(us.bkt_name, func(bkt *Bucket) error {
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
	err = us.db.UpdateBucket(us.bkt_name, func(bkt *Bucket) error {
		if buf, err := json.Marshal(&u); err != nil {
			return err
		} else if err := bkt.Put([]byte(u.Username), buf); err != nil {
			return err
		}
		return nil
	})
	return
}
