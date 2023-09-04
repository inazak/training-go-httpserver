package service

import (
	"context"
	"github.com/inazak/training-go-httpserver/model"
)

// golang/mock はすでに archived になっている
// インストールは go install github.com/golang/mock/mockgen@v1.6.0
// 下記の通り書いてから go generate ./service.go でモックを生成

//go:generate mockgen -source=$GOFILE -package=mock -destination=../common/mock/service_$GOFILE
type Service interface {
	HealthCheck(ctx context.Context) error
	GetTaskList(ctx context.Context) (model.TaskList, error)
	AddTask(ctx context.Context, title string) (*model.Task, error)
	AddUser(ctx context.Context, name string, password string, role string) (model.UserID, error)
}
