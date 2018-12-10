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

func NewUserRouter(u root.UserService, userHandle func(string, http.Handler)) {
	userRouter := &UserRouter{u}

	userHandle("/login", ErrorHandler{userRouter.loginHandler})
	userHandle("/signup", ErrorHandler{userRouter.signupHandler})
}

func (ur *UserRouter) signupHandler(w http.ResponseWriter, r *http.Request) error {
	new_user, err := decodeNewUser(r)
	if err != nil {
		return root.StatusError{
			Code: 400,
			Err:  err,
		}
	}

	_, err = ur.userService.Signup(&new_user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRouter) loginHandler(w http.ResponseWriter, r *http.Request) error {
	credentials, err := decodeCredentials(r)

	_, err = ur.userService.Login(&credentials)
	if err != nil {
		return root.StatusError{
			Code: 404,
			Err:  err,
		}
	}

	return nil
}

func decodeCredentials(r *http.Request) (root.Credentials, error) {
	var c root.Credentials
	if r.Body == nil {
		return c, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	return c, err
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

func decodeNewUser(r *http.Request) (root.NewUser, error) {
	var nu root.NewUser
	if r.Body == nil {
		return nu, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nu)
	return nu, err
}
