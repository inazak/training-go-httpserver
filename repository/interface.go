package repository

import (
	"context"
	"github.com/inazak/training-go-httpserver/model"
)

type Repository interface {
	InsertTask(ctx context.Context, task *model.Task) error
	SelectTaskList(ctx context.Context) (model.TaskList, error)
}
