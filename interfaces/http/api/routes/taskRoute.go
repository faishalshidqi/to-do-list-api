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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	taskRepository := repository.NewTaskRepository(db, domains.TaskCollectionName)
	taskController := controllers.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(taskRepository, timeout),
		Env:         env,
	}
	group.POST("/api/tasks", taskController.Create)
	group.GET("/api/tasks", taskController.GetByOwner)
	group.GET("/api/tasks/:id", taskController.GetById)
	group.PUT("/api/tasks/:id", taskController.Update)
	group.DELETE("/api/tasks/:id", taskController.Delete)
	group.PUT("/api/tasks/:id/mark", taskController.MarkAsCompleted)
	group.GET("/api/tasks/completed", taskController.FetchCompletedTasks)
}
