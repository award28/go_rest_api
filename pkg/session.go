package root

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type Session interface {
	Get(*http.Request, func(*sessions.Session, string) error) error
	Set(*http.Request, http.ResponseWriter, func(*sessions.Session) error) error
}
