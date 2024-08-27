package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect with mysql database
var DB *gorm.DB


func Connect() {
	// connect with mysql database
	dsn := "naol:naol.@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection fail")
	}
	DB = db
	fmt.Println("database is connected")
}

func GetDB() *gorm.DB {
	return DB
}