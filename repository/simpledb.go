package repository

import (
	"fmt"
	"context"
	"github.com/inazak/training-go-httpserver/common/clock"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository/database"
)

// repository.Repository の実装
type SimpleDB struct {
	database.Database
}

func NewSimpleDB(db database.Database) *SimpleDB {
	return &SimpleDB{
		Database: db,
	}
}

func (sd *SimpleDB) InsertTask(ctx context.Context, task *model.Task) error {
	task.Created = clock.NowString()
	task.Modified = task.Created

	sql := `INSERT INTO task (userid, title, status, created, modified) VALUES (:userid, :title, :status, :created, :modified);`

	//FIXME この動作はsqlxが前提となり、抽象を破壊している
	result, err := sd.NamedExec(ctx, sql, task)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = model.TaskID(id)
	return nil
}

func (sd *SimpleDB) SelectTaskList(ctx context.Context) (model.TaskList, error) {
	tasklist := model.TaskList{}

	//FIXME この動作はsqlxが前提となり、抽象を破壊している
	sql := `SELECT * FROM task;`
	if err := sd.Select(ctx, &tasklist, sql); err != nil {
		return nil, err
	}
	return tasklist, nil
}

func (sd *SimpleDB) InsertUser(ctx context.Context, user *model.User) error {
	user.Created = clock.NowString()
	user.Modified = user.Created

	sql := `INSERT INTO user (name, password, role, created, modified) VALUES (:name, :password, :role, :created, :modified);`

	//FIXME この動作はsqlxが前提となり、抽象を破壊している
	result, err := sd.NamedExec(ctx, sql, user)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = model.UserID(id)
	return nil
}

func (sd *SimpleDB) SelectUser(ctx context.Context, name string) (*model.User, error) {
	result := []*model.User{}

	sql := `SELECT * FROM user WHERE name = ?;`
	if err := sd.Select(ctx, &result, sql, name); err != nil {
		return nil, err
	}

	if len(result) == 0 {
    return nil, fmt.Errorf("no match")
	}

	if len(result) > 1 {
    return nil, fmt.Errorf("unexpected multi match")
	}
	return result[0], nil
}

