package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type TaskController struct {
	TaskUsecase domains.TaskUsecase
	Env         *bootstrap.Env
}

// Create AddTask godoc
//
//	@Summary		Add A New Task To DB
//	@Description	Add A New Task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			title			body		string	true	"task's title"
//	@Param			password		body		string	true	"task's description"
//	@Param			isCompleted		body		bool	true	"whether the task is completed"
//
//	@Success		201				{object}	domains.AddTaskResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/api/tasks [post]
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
	c.JSON(http.StatusCreated, domains.AddTaskResponse{
		Message: "Task created successfully",
		Data: domains.AddTaskResponseData{
			ID: task.ID,
		},
	})
}

// GetByOwner GetTasksByOwner godoc
//
//	@Summary		Fetch Tasks
//	@Description	Fetch Tasks By Owner ID. Only valid users may have tasks
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			page			body		string	true	"task's title"
//	@Param			size			body		string	true	"task's description"
//
//	@Success		200				{object}	domains.GetTaskResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		404				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/api/tasks [get]
func (tc *TaskController) GetByOwner(c *gin.Context) {
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

	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "1")
	tasks, err := tc.TaskUsecase.FetchByOwner(c, userId.Hex(), page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if tasks == nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Message: "Task not found",
		})
		return
	}
	c.JSON(http.StatusOK, domains.GetTaskResponse{
		Message: "Successfully fetched tasks",
		Data: domains.GetTaskResponseData{
			Tasks: tasks,
		},
	})
}
