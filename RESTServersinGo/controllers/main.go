package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/RESTServersinGo/models"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func GetTasks(c *gin.Context) {

	// var tasks []models.Task;
	tasks := models.GetAllTasks()

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	err := task.CreateTask()
	fmt.Println(task)
	if err != nil {
		c.JSON(201, gin.H{
			"task": task,
		})
		} else {
		c.JSON(400, gin.H{
			"error": err,
		})
	}
}