package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/invinciible/go-bookstore-api/bookstore_users-api/domain/users"
	"github.com/invinciible/go-bookstore-api/bookstore_users-api/services"
	resterr "github.com/invinciible/go-bookstore-api/bookstore_users-api/utils/errors"
)

// CreateUsers creates new user.
func CreateUsers(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		re := resterr.NewBadRequestError("Invalid JSON")
		c.JSON(re.Status, re)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser get the user.
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userErr != nil {
		err := resterr.NewBadRequestError("User ID should be an integer")
		c.JSON(err.Status, err)
	}
	result, getErr := services.GetUser(userID)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// SearchUser finds user.
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
