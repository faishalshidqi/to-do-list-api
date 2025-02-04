package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/interfaces/http/api/routes"
)

//	@title			To-Do List API
//	@version		1.0
//	@description	This is a To-Do List API

// @host		localhost:8080
// @BasePath	/
// @securityDefinitions.apikey
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.MongoDB)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	router := gin.Default()
	routes.Setup(env, timeout, db, router)
	router.Run(env.ServerAddress)
}
