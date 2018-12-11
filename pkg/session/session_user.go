package session

import (
	"encoding/gob"
	"go_rest_api/pkg"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	gob.Register(&root.User{})
	return &UserMiddleware{}
}

func (um *UserMiddleware) AuthenticateUser() {}
