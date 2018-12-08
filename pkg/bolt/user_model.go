package bolt

import (
	"go_rest_api/pkg"
)

type userModel struct {
	Username string
	Password string
}

func newUserModel(u *root.User) *userModel {
	return &userModel{
		Username: u.Username,
		Password: u.Password}
}

func (u *userModel) toRootUser() *root.User {
	return &root.User{
		Username: u.Username,
		Password: u.Password}
}
