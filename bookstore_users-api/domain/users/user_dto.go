package users

import (
	"strings"

	resterr "github.com/invinciible/go-bookstore-api/bookstore_users-api/utils/errors"
)

// User Model of user
type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	ID          int64  `json:"Id"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

// Validate validation for user
func (user *User) Validate() *resterr.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return resterr.NewBadRequestError("Invalid Email address")
	}
	return nil
}
