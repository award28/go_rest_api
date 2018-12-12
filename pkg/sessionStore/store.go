package sessionStore

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

type Store struct {
	name  string
	store *sessions.CookieStore
}

type Session = sessions.Session

func NewStore(name string) *Store {
	key := []byte(os.Getenv("SESSION_KEY"))
	return &Store{
		name:  name,
		store: sessions.NewCookieStore(key),
	}
}

func (s *Store) Get(r *http.Request, fn func(*Session) error) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}

	return fn(session)
}

func (s *Store) Set(r *http.Request, w http.ResponseWriter, fn func(*Session) error) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}

	err = fn(session)
	if err != nil {
		return err
	}

	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
