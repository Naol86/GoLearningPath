package views

import (
	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/RESTServersinGo/controllers"
)

func Routes(route *gin.Engine) {
	api := route.Group("/api")

	api.GET("/hello", controllers.Hello)

	// get all tasks
	api.GET("/tasks", controllers.GetTasks)
	// create a task
	api.POST("/tasks", controllers.CreateTask)

}