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
//	@Param			page			query		string	true	"page number, acting as offset"
//	@Param			size			query		string	true	"page size, acting as limit"
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
		Data:    tasks,
	})
}

// GetById GetTaskById godoc
//
//	@Summary		Fetch Task
//	@Description	Fetch Tasks By ID. Only valid task may get returned
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			id				path		string	true	"Task ID"
//
//	@Success		200				{object}	domains.GetTaskByIdResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		403				{object}	domains.ErrorResponse
//	@Failure		404				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/api/tasks/{id} [get]
func (tc *TaskController) GetById(c *gin.Context) {
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
	id := c.Param("id")
	task, err := tc.TaskUsecase.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Message: "Task not found",
		})
		return
	}
	if userId != task.Owner {
		c.JSON(http.StatusForbidden, domains.ErrorResponse{
			Message: "You can't see this task",
		})
		return
	}
	c.JSON(http.StatusOK, domains.GetTaskByIdResponse{
		Message: "Successfully fetched task",
		Data:    *task,
	})
}

// Update UpdateTaskById godoc
//
//	@Summary		Edit Task
//	@Description	Edit Tasks By ID. Only valid task may be edited
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			id				path		string	true	"Task ID"
//	@Param			title			body		string	true	"Title of the task"
//	@Param			description		body		string	true	"Description of the task"
//	@Param			isCompleted		body		string	true	"whether the task is completed"
//
//	@Success		200				{object}	domains.GetTaskByIdResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		403				{object}	domains.ErrorResponse
//	@Failure		404				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/api/tasks/{id} [put]
func (tc *TaskController) Update(c *gin.Context) {
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
	task := &domains.Task{}
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}
	id := c.Param("id")
	fetchedTask, err := tc.TaskUsecase.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Message: "Task not found",
		})
		return
	}
	if userId != fetchedTask.Owner {
		c.JSON(http.StatusForbidden, domains.ErrorResponse{
			Message: "You can't see this task",
		})
		return
	}
	task.UpdatedAt = time.Now()
	task.CreatedAt = fetchedTask.CreatedAt
	task.Owner = userId
	task.ID = fetchedTask.ID
	err = tc.TaskUsecase.EditById(c, id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domains.UpdateTaskResponse{
		Message: "Successfully updated task",
		Data:    *task,
	})
}
