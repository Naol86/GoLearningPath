package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/RESTServersinGo/databases"
	"github.com/naol86/learn-go/RESTServersinGo/views"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	// create gin server with default middleware
	databases.Connect();
	router := gin.Default()

	views.Routes(router)

	router.Run(":8080")
}
