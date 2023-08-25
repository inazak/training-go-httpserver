package service

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository"
)

// これより上位のUI層で詳細なエラー情報を返すことはない
// そのため、この層でのロギングが必要である
type TodoService struct {
	repo   repository.Repository
	logger log.Logger
}

func NewTodoService(ctx context.Context, r repository.Repository, logger log.Logger) *TodoService {
	return &TodoService{
		repo:   r,
		logger: logger,
	}
}

func (st *TodoService) HealthCheck(ctx context.Context) error {
	level.Info(st.logger).Log("msg", "in service.HealthCheck")
	return nil
}

func (st *TodoService) GetTaskList(ctx context.Context) (model.TaskList, error) {
	level.Info(st.logger).Log("msg", "in service.GetTaskList")
	rs, err := st.repo.SelectTaskList(ctx)
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
	err := st.repo.InsertTask(ctx, task)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.AddTask", "err", err)
		return nil, err
	}
	return task, nil
}
