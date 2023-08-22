package repository

import (
  "context"
  "github.com/inazak/training-go-httpserver/model"
)

type Repository interface {
  SetTask(ctx context.Context, task *model.Task) error
  GetTaskList(ctx context.Context) (*model.TaskList, error)
}

