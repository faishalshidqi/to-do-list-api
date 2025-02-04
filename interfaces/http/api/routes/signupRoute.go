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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, domains.UserCollectionName)
	signupController := controllers.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepository, timeout),
		Env:           env,
	}
	group.POST("/api/auth/register", signupController.Signup)
}
