package domains

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID          primitive.ObjectID `json:"-" bson:"_id"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Owner       primitive.ObjectID `json:"owner" bson:"owner" binding:"required"`
	IsCompleted bool               `json:"isCompleted" bson:"isCompleted"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type TaskRepository interface {
	Add(c context.Context, task *Task) error
	FetchByOwner(c context.Context, owner string) ([]Task, error)
}

type TaskUsecase interface {
	Add(c context.Context, task *Task) error
	FetchByOwner(c context.Context, owner string) ([]Task, error)
}
