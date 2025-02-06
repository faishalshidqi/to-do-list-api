package usecase

import (
	"context"
	"time"
	"todo-list-api/domains"
)

type taskUsecase struct {
	taskRepository domains.TaskRepository
	contextTimeout time.Duration
}

func (tu *taskUsecase) FetchById(c context.Context, id string) (*domains.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchById(ctx, id)
}

func (tu *taskUsecase) EditById(c context.Context, id string, task *domains.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.EditById(ctx, id, task)
}

func (tu *taskUsecase) DeleteById(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.DeleteById(ctx, id)
}

func (tu *taskUsecase) MarkAsCompleted(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.MarkAsCompleted(ctx, id)
}

func (tu *taskUsecase) FetchCompleted(c context.Context, owner string) ([]domains.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchCompleted(ctx, owner)
}

func (tu *taskUsecase) Add(c context.Context, task *domains.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Add(ctx, task)
}

func (tu *taskUsecase) FetchByOwner(c context.Context, owner string, page, size string) ([]domains.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByOwner(ctx, owner, page, size)
}

func NewTaskUsecase(taskRepository domains.TaskRepository, timeout time.Duration) domains.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}
