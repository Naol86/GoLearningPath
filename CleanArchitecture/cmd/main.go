package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/CleanArchitecture/api/route"
	"github.com/naol86/learn-go/CleanArchitecture/config"
)


func main() {
	app := config.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)
	gin.Run(env.ServerAddress)
}