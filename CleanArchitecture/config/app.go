package config

import (
	"github.com/naol86/learn-go/CleanArchitecture/pkg/mongo"
)


type Application struct {
	Env *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection(){
	CloseMongoDBConnection(app.Mongo);
}