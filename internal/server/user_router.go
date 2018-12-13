package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_rest_api/internal"
	"net/http"
)

type UserRouter struct {
	userService root.UserService
	userStore   root.UserStore
}

func NewUserRouter(userService root.UserService, userStore root.UserStore, userHandle func(string, http.Handler)) {
	userRouter := &UserRouter{
		userService: userService,
		userStore:   userStore,
	}

	userHandle("/", ErrorHandler{userRouter.meHandler})
	userHandle("/login", ErrorHandler{userRouter.loginHandler})
	userHandle("/signup", ErrorHandler{userRouter.signupHandler})
	userHandle("/logout", ErrorHandler{userRouter.logoutHandler})
}

func (ur *UserRouter) meHandler(w http.ResponseWriter, r *http.Request) error {
	user, err := ur.userStore.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil
	}

	fmt.Fprintf(w, "username: %s, password: %s", user.Username, user.Password)
	return nil
}

func (ur *UserRouter) loginHandler(w http.ResponseWriter, r *http.Request) error {
	credentials, err := decodeCredentials(r)

	user, err := ur.userService.Login(&credentials)
	if err != nil {
		return root.StatusError{
			Code: 404,
			Err:  err,
		}
	}

	err = ur.userStore.SetSessionUser(r, w, user)
	if err != nil {
		return nil
	}

	fmt.Fprintf(w, "Welcome, %s!", user.Username)
	return nil
}

func (ur *UserRouter) signupHandler(w http.ResponseWriter, r *http.Request) error {
	new_user, err := decodeNewUser(r)
	if err != nil {
		return root.StatusError{
			Code: 400,
			Err:  err,
		}
	}

	user, err := ur.userService.Signup(&new_user)
	if err != nil {
		return err
	}

	err = ur.userStore.SetSessionUser(r, w, user)
	if err != nil {
		return nil
	}

	fmt.Fprintf(w, "Welcome, %s!", user.Username)
	return nil
}

func (ur *UserRouter) logoutHandler(w http.ResponseWriter, r *http.Request) error {
	err := ur.userStore.DeleteSessionUser(r, w)
	if err != nil {
		return nil
	}

	fmt.Fprintf(w, "You are now logged out.")
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
