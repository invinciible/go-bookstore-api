package app

import (
	"github.com/invinciible/go-bookstore-api/bookstore_users-api/controllers/ping"
	"github.com/invinciible/go-bookstore-api/bookstore_users-api/controllers/users"
)

func mapURLs() {

	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUsers)
	router.GET("/users/:userId", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
}
