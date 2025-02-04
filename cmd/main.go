package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/interfaces/http/api/routes"
)

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
