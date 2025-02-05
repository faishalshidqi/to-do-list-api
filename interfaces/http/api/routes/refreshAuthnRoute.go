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

func NewRefreshAuthnRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, domains.UserCollectionName)
	refreshAuthnRepository := repository.NewRefreshAuthnRepository(db, domains.RefreshTokenCollectionName)
	refreshAuthnController := controllers.RefreshAuthnController{
		RefreshAuthnRepo:    refreshAuthnRepository,
		RefreshAuthnUsecase: usecase.NewRefreshAuthnUsecase(userRepository, timeout),
		Env:                 env,
	}
	group.PUT("/api/auth", refreshAuthnController.RefreshToken)
}
