package routes

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/infrastructures/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	router := gin.Group("")
	NewSignupRouter(env, timeout, db, router)
	NewLoginRouter(env, timeout, db, router)
	NewRefreshAuthnRouter(env, timeout, db, router)
	//NewLogoutRouter(env, timeout, db, router)

	NewTaskRouter(env, timeout, db, router)
}
