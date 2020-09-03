package users

import (
	"fmt"

	resterr "github.com/invinciible/go-bookstore-api/bookstore_users-api/utils/errors"
)

var (
	userDb = make(map[int64]*User)
)

// Get user from DB using userId as primary key.
func (user *User) Get() *resterr.RestError {
	result := userDb[user.ID]
	if result == nil {
		return resterr.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save user to DB.
func (user *User) Save() *resterr.RestError {
	current := userDb[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return resterr.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return resterr.NewBadRequestError(fmt.Sprintf("user %d already exist", user.ID))
	}
	userDb[user.ID] = user
	return nil
}
