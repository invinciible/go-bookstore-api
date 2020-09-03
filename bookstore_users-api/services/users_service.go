package services

import (
	"github.com/invinciible/go-bookstore-api/bookstore_users-api/domain/users"
	resterr "github.com/invinciible/go-bookstore-api/bookstore_users-api/utils/errors"
)

// CreateUser service
func CreateUser(user users.User) (*users.User, *resterr.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser service
func GetUser(userID int64) (*users.User, *resterr.RestError) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
