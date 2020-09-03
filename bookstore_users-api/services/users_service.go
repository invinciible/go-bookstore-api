package services

import (
	"github.com/invinciible/go-bookstore-api/bookstore_users-api/domain/users"
	resterr "github.com/invinciible/go-bookstore-api/bookstore_users-api/utils/errors"
)

// CreateUser service
func CreateUser(user users.User) (*users.User, *resterr.RestError) {
	return &user, nil
}
