package routes

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/infrastructures/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshAuthnRouter(env, timeout, db, publicRouter)
	//NewLogoutRouter(env, timeout, db, publicRouter)
}
