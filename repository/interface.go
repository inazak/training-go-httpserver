package repository

import (
	"context"
	"github.com/inazak/training-go-httpserver/model"
)

type Database interface {
	InsertTask(ctx context.Context, task *model.Task) error
	SelectTaskList(ctx context.Context, id model.UserID) (model.TaskList, error)
	InsertUser(ctx context.Context, user *model.User) error
	SelectUser(ctx context.Context, name string) (*model.User, error)
}

type KVS interface {
	SetUserID(ctx context.Context, key string, id model.UserID, ttl int64) error
	GetUserID(ctx context.Context, key string) (model.UserID, error)
}
