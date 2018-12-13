package root

import (
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailConfirm    string `json:"email_confirm"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type UserService interface {
	Me() (*User, error)
	Login(c *Credentials) (*User, error)
	Signup(nu *NewUser) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(u *User) error
}

type UserStore interface {
	GetSessionUser(*http.Request) (*User, error)
	SetSessionUser(*http.Request, http.ResponseWriter, *User) error
	DeleteSessionUser(*http.Request, http.ResponseWriter) error
}
