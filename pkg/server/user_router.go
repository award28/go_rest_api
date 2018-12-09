package server

import (
	"fmt"
	"go_rest_api/pkg"
	"net/http"
)

type UserRouter struct {
	userService root.UserService
}

func NewUserRouter(u root.UserService) *UserRouter {
	userRouter := &UserRouter{u}

	//http.Handle("/", handlers.Handler{env, handlers.GetIndex})
	return userRouter
}

func (ur *UserRouter) createUserHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "%s", "Load index page")
	return nil
}
