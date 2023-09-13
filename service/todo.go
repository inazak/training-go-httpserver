package service

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository"
	"golang.org/x/crypto/bcrypt"
)

// これより上位のUI層で詳細なエラー情報を返すことはない
// そのため、この層でのロギングが必要である
type TodoService struct {
	db     repository.Database
	logger log.Logger
}

func NewTodoService(ctx context.Context, db repository.Database, logger log.Logger) *TodoService {
	return &TodoService{
		db:     db,
		logger: logger,
	}
}

func (st *TodoService) HealthCheck(ctx context.Context) error {
	level.Info(st.logger).Log("msg", "in service.HealthCheck")
	return nil
}

func (st *TodoService) GetTaskList(ctx context.Context) (model.TaskList, error) {
	level.Info(st.logger).Log("msg", "in service.GetTaskList")
	rs, err := st.db.SelectTaskList(ctx)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.SelectTaskList", "err", err)
		return nil, err
	}
	return rs, nil
}

func (st *TodoService) AddTask(ctx context.Context, title string) (*model.Task, error) {

	level.Info(st.logger).Log("msg", "in service.AddTask")
	task := &model.Task{
		Title:  title,
		Status: model.TaskStatusTodo,
	}
	err := st.db.InsertTask(ctx, task)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.AddTask", "err", err)
		return nil, err
	}
	return task, nil
}

func (st *TodoService) AddUser(ctx context.Context, name string, password string, role string) (model.UserID, error) {

	pwhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.AddUser", "err", err)
		return -1, err
	}

	level.Info(st.logger).Log("msg", "in service.AddUser")
	user := &model.User{
		Name:     name,
		Password: string(pwhash),
		Role:     role,
	}
	err = st.db.InsertUser(ctx, user)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.AddUser", "err", err)
		return -1, err
	}
	return user.ID, nil
}

func (st *TodoService) GetUser(ctx context.Context, name string) (*model.User, error) {
	level.Info(st.logger).Log("msg", "in service.GetUser")
	user, err := st.db.SelectUser(ctx, name)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.SelectUser", "err", err)
		return nil, err
	}
	return user, nil
}

