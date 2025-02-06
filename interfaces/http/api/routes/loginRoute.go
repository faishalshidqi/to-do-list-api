package routes

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo-list-api/applications/usecase"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/domains"
	"todo-list-api/infrastructures/mongo"
	"todo-list-api/infrastructures/repository"
	"todo-list-api/interfaces/http/api/controllers"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, domains.UserCollectionName)
	loginController := controllers.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(userRepository, timeout),
		Env:          env,
	}
	group.POST("/api/auth", loginController.Login)
}
