package server

import (
	"encoding/json"
	"errors"
	"go_rest_api/pkg"
	"net/http"
)

type UserRouter struct {
	userService root.UserService
}

func NewUserRouter(u root.UserService, handle func(string, http.Handler)) {
	userRouter := &UserRouter{u}

	handle("/", Handler{userRouter.createUserHandler})
}

func (ur *UserRouter) createUserHandler(w http.ResponseWriter, r *http.Request) error {
	user, err := decodeUser(r)
	if err != nil {
		return StatusError{
			Code: 400,
			Err:  err,
		}
	}

	err = ur.userService.Create(&user)
	if err != nil {
		return StatusError{
			Code: 400,
			Err:  err,
		}
	}
	return nil
}

func decodeUser(r *http.Request) (root.User, error) {
	var u root.User
	if r.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}
