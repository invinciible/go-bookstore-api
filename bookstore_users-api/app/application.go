package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication To start the app
func StartApplication() {
	mapURLs()
	router.Run(":8080")
}
