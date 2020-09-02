package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUsers creates new user.
func CreateUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// GetUser get the user.
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// SearchUser finds user.
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
