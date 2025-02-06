package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type TaskController struct {
	TaskUsecase domains.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *TaskController) Create(c *gin.Context) {
	token, err := tokenize.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	userId, err := tokenize.ValidateJWT(token, tc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var task domains.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Owner = userId

	err = tc.TaskUsecase.Add(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domains.TaskResponse{
		Message: "Task created successfully",
		Data: domains.TaskResponseData{
			ID: task.ID,
		},
	})
}
