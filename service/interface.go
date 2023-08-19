package service

import (
  "context"

  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/store"
)

type TaskAdder interface {
  AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
  ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}


