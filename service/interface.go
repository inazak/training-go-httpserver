package service

import (
	"context"
	"github.com/inazak/training-go-httpserver/model"
	"net/http"
)

// インストールは go install github.com/golang/mock/mockgen@v1.6.0
// 下記の通り書いてから go generate ./service.go でモックを生成

//go:generate mockgen -source=$GOFILE -package=mock -destination=../common/mock/service_$GOFILE
type Service interface {
	HealthCheck(ctx context.Context) error
	GetTaskList(ctx context.Context, id model.UserID) (model.TaskList, error)
	AddTask(ctx context.Context, id model.UserID, title string) (*model.Task, error)
	AddUser(ctx context.Context, name string, password string) (model.UserID, error)
	GetUser(ctx context.Context, name string) (*model.User, error)
	Login(ctx context.Context, name string, password string) (string, error)
	ValidateToken(ctx context.Context, r *http.Request) (model.UserID, error)
}
