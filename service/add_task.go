package service

import (
  "context"
  "fmt"

  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/store"
)

// store.Execer はDBを隠しているはずなので、DBという名前が美しくない
type AddTask struct {
  DB   store.Execer
  Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {

  task := &entity.Task{
    Title:  title,
    Status: entity.TaskStatusTodo,
  }

  err := a.Repo.AddTask(ctx, a.DB, task)
  if err != nil {
    return nil, fmt.Errorf("failed to register: %w", err)
  }

  return task, nil
}

