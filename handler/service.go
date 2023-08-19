package handler

import (
  "context"
  "github.com/inazak/training-go-httpserver/entity"
)

// golang/mock はすでに archived になっている
// インストールは go install github.com/golang/mock/mockgen@v1.6.0
// 下記の通り書いてから go generate ./handler/service.go でモックを生成

//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/mock_$GOFILE

type ListTasksService interface {
  ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
  AddTask(ctx context.Context, title string) (*entity.Task, error)
}

