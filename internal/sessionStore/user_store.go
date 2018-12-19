package sessionStore

import (
	"encoding/gob"
	"errors"
	"go_rest_api/internal"
	"net/http"
)

var (
	noSessionUserErr = errors.New("No current user in session.")
)

const (
	USER = "user"
)

type UserStore struct {
	store *Store
}

func NewUserStore(store *Store) *UserStore {
	gob.Register(&root.User{})
	return &UserStore{
		store: store,
	}
}

func (us *UserStore) GetSessionUser(r *http.Request) (user *root.User, err error) {
	err = us.store.Get(r, func(s *Session) error {
		val, ok := s.Values[USER]
		if !ok {
			return noSessionUserErr
		}

		u := &root.User{}
		u, ok = val.(*root.User)
		if !ok {
			return noSessionUserErr
		}
		user = u
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserStore) SetSessionUser(r *http.Request, w http.ResponseWriter, u *root.User) error {
	return us.store.Set(r, w, func(s *Session) error {
		s.Values[USER] = u
		return nil
	})
}

func (us *UserStore) DeleteSessionUser(r *http.Request, w http.ResponseWriter) error {
	return us.store.Set(r, w, func(s *Session) error {
		_, ok := s.Values[USER]
		if !ok {
			return noSessionUserErr
		}
		delete(s.Values, USER)
		return nil
	})
}
