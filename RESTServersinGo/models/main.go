package models

import (
	"time"

	"github.com/naol86/learn-go/RESTServersinGo/databases"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}


var DB *gorm.DB;

func init() {
	databases.Connect()
	DB = databases.GetDB()
	DB.AutoMigrate(&Task{})
	// DB.AutoMigrate(&Task{})
}

func GetAllTasks() []Task {
	var tasks []Task
	DB.Find(&tasks)
	return tasks
}

func (t *Task) CreateTask() *Task {
	DB.Create(&t)
	return t
}