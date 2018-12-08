package root

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	Create(u *User) error
	GetByUsername(username string) (*User, error)
}
