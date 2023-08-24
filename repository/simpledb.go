package repository

import (
  "context"
  "github.com/inazak/training-go-httpserver/model"
  "github.com/inazak/training-go-httpserver/common/clock"
  "github.com/inazak/training-go-httpserver/repository/database"
)

// repository.Repository の実装
type SimpleDB struct {
  database.Database
}

func NewSimpleDB(db database.Database) *SimpleDB {
  return &SimpleDB {
    Database: db,
  }
}

func (sd *SimpleDB) InsertTask(ctx context.Context, task *model.Task) error {
  task.Created  = clock.NowString()
  task.Modified = task.Created

  sql := `INSERT INTO task (title, status, created, modified) VALUES (:title, :status, :created, :modified);`

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

  // sqlxにより構造体の名前に結果がマッピングされる
  sql := `SELECT * FROM task;`
  if err := sd.Select(ctx, &tasklist, sql); err != nil {
    return nil, err
  }
  return tasklist, nil
}


