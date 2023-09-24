package service

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/inazak/training-go-httpserver/common/jwter"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// これより上位のUI層で詳細なエラー情報を返すことはない
// そのため、この層でのロギングが必要である
type TodoService struct {
	db     repository.Database
	kvs    repository.KVS
	jwter  *jwter.JWTer
	logger log.Logger
}

func NewTodoService(ctx context.Context, db repository.Database, kvs repository.KVS, jwter *jwter.JWTer, logger log.Logger) *TodoService {
	return &TodoService{
		db:     db,
		kvs:    kvs,
		jwter:  jwter,
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

func (st *TodoService) AddTask(ctx context.Context, id model.UserID, title string) (*model.Task, error) {

	level.Info(st.logger).Log("msg", "in service.AddTask")
	task := &model.Task{
		UserID: id,
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

func (st *TodoService) AddUser(ctx context.Context, name string, password string) (model.UserID, error) {

	pwhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.AddUser", "err", err)
		return -1, err
	}

	level.Info(st.logger).Log("msg", "in service.AddUser")
	user := &model.User{
		Name:     name,
		Password: string(pwhash),
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

func (st *TodoService) Login(ctx context.Context, name string, password string) (string, error) {
	level.Info(st.logger).Log("msg", "in service.Login")
	user, err := st.GetUser(ctx, name)
	if err != nil {
		level.Error(st.logger).Log("msg", "in repository.SelectUser", "err", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		level.Error(st.logger).Log("msg", "in compare password", "err", err)
		return "", err
	}

	//FIXME キー名や失効までの時間を外で定義しなおすこと
	claims := []jwter.Claim{
		{Key: "uid", Value: fmt.Sprintf("%d", user.ID)},
	}
	jwtid, token, err := st.jwter.GenerateToken("accesstoken", time.Minute*10, claims)
	if err != nil {
		level.Error(st.logger).Log("msg", "in generate token", "err", err)
		return "", err
	}

	err = st.kvs.SetUserID(ctx, jwtid, user.ID, 600)
	if err != nil {
		level.Error(st.logger).Log("msg", "in kvs.SetUserID", "err", err)
		return "", err
	}

	return string(token), nil
}

func (st *TodoService) ValidateToken(ctx context.Context, r *http.Request) (model.UserID, error) {
	level.Info(st.logger).Log("msg", "in service.ValidateToken")

	token, err := st.jwter.ParseRequest(r)
	if err != nil {
		level.Error(st.logger).Log("msg", "in jwter.ParseRequest", "err", err)
		return -1, err
	}

	id, err := st.kvs.GetUserID(ctx, token.JwtID())
	if err != nil {
		level.Error(st.logger).Log("msg", "in kvs.GetUserID", "err", err)
		return -1, err
	}

	return id, nil
}
