package service

import (
  "context"
  "github.com/inazak/training-go-httpserver/model"
  "github.com/inazak/training-go-httpserver/repository"
)

type SimpleTodoService struct {
  repo repository.Repository
}

func NewSimpleTodoService(ctx context.Context, r repository.Repository) *SimpleTodoService {
  return &SimpleTodoService{
    repo: r,
  }
}

func(st *SimpleTodoService) HealthCheck(ctx context.Context) error {
  return nil
}

func(st *SimpleTodoService) GetTaskList(ctx context.Context) (model.TaskList, error) {
  rs, err := st.repo.SelectTaskList(ctx)
  if err != nil {
    return nil, err
  }
  return rs, nil
}

func(st *SimpleTodoService) AddTask(ctx context.Context, title string) (*model.Task, error) {

  task := &model.Task{
    Title: title,
    Status: model.TaskStatusTodo,
  }
  err := st.repo.InsertTask(ctx, task)
  if err != nil {
    return nil, err
  }
  return task, nil
}


