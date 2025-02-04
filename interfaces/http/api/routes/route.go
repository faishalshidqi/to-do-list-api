package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo-list-api/commons/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
}
