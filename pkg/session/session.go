package session

import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"go_rest_api/pkg"
	"os"
)

type Session struct {
	name  string
	store *sessions.CookieStore
}

func NewSession(name string) *Session {
	key := []byte(os.Getenv("SESSION_KEY"))
	if len(key) == 0 {
		key = securecookie.GenerateRandomKey(32)
		os.Setenv("SESSION_KEY", key)
	}

	return &Session{
		name:  name,
		store: sessions.NewCookieStore(key),
	}
}

func (as *Session) Get(r *http.Request, fn func(*sessions.Session) error) error {
	session, err := as.store(r, as.name)
	if err != nil {
		return err
	}

	return fn(session)
}

func (as *Session) Set(r *http.Request, w http.ResponseWriter, fn func(*sessions.Session) error) error {
	session, err := as.store(r, as.name)
	if err != nil {
		return err
	}

	err := fn(session)
	if err != nil {
		return err
	}

	session.Save(r, w)

	return nil
}
