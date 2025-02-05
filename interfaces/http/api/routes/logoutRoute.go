package routes

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/domains"
	"todo-list-api/infrastructures/mongo"
	"todo-list-api/infrastructures/repository"
	"todo-list-api/interfaces/http/api/controllers"
)

func NewLogoutRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	refreshAuthnRepository := repository.NewRefreshAuthnRepository(db, domains.RefreshTokenCollectionName)
	logoutController := controllers.LogoutController{
		RefreshAuthnRepo: refreshAuthnRepository,
		Env:              env,
	}
	group.DELETE("/api/auth", logoutController.Logout)
}
