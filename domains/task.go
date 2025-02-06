package domains

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	TaskCollectionName = "tasks"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Owner       primitive.ObjectID `json:"owner" bson:"owner"`
	IsCompleted bool               `json:"isCompleted" bson:"isCompleted" default:"false"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type AddTaskResponseData struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

type AddTaskResponse struct {
	Message string              `json:"message"`
	Data    AddTaskResponseData `json:"data"`
}

type TaskRepository interface {
	Add(c context.Context, task *Task) error
	FetchByOwner(c context.Context, owner string) ([]Task, error)
	FetchById(c context.Context, id string) (*Task, error)
	EditById(c context.Context, id string, task *Task) error
	DeleteById(c context.Context, id string) error
	MarkAsCompleted(c context.Context, id string) error
	FetchCompleted(c context.Context, owner string) ([]Task, error)
}

type TaskUsecase interface {
	Add(c context.Context, task *Task) error
	FetchByOwner(c context.Context, owner string) ([]Task, error)
	FetchById(c context.Context, id string) (*Task, error)
	EditById(c context.Context, id string, task *Task) error
	DeleteById(c context.Context, id string) error
	MarkAsCompleted(c context.Context, id string) error
	FetchCompleted(c context.Context, owner string) ([]Task, error)
}
