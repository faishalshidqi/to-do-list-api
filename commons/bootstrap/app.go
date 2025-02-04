package bootstrap

import "todo-list-api/infrastructures/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return app
}

func (app *Application) CloseDBConnection() {
	CloseMongoConnnection(app.Mongo)
}
