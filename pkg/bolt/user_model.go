package bolt

import (
	"errors"
	"go_rest_api/pkg"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func nonEmptyFields_NewUser(nu root.NewUser) error {
	var missing strings.Builder
	missing.WriteString(emptyOr(nu.Username, "username\n"))
	missing.WriteString(emptyOr(nu.Password, "password\n"))
	missing.WriteString(emptyOr(nu.PasswordConfirm, "password_confirm\n"))
	missing.WriteString(emptyOr(nu.Email, "email\n"))
	missing.WriteString(emptyOr(nu.EmailConfirm, "email_confirm\n"))

	if missing.String() == "" {
		return nil
	} else {
		return errors.New("The following fields are required:\n" + missing.String())
	}
}

func verifyFields_NewUser(nu root.NewUser) error {
	var incorrect strings.Builder

	if !emailRegex.MatchString(nu.Email) {
		incorrect.WriteString("Email is not a proper email.\n")
	} else if nu.Email != nu.EmailConfirm {
		incorrect.WriteString("Emails do not match.\n")
	}

	if nu.Password != nu.PasswordConfirm {
		incorrect.WriteString("Passwords do not match.")
	}

	if incorrect.String() == "" {
		return nil
	} else {
		return errors.New("The following fields are malformed:\n" + incorrect.String())
	}
}

func emptyOr(s, or string) string {
	if s == "" {
		return or
	}
	return ""
}
