package store

import (
  "context"
  "github.com/inazak/training-go-httpserver/entity"
)

// RDBによる永続化を行ったTaskの扱い

func (r *Repository) ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error) {
  tasks := entity.Tasks{}

  // sqlxにより構造体の名前に結果がマッピングされる
  sql := `SELECT * FROM task;`
  if err := db.SelectContext(ctx, &tasks, sql); err != nil {
    return nil, err
  }
  return tasks, nil
}

func (r *Repository) AddTask(ctx context.Context, db Execer, task *entity.Task) error {
  task.Created  = r.Clocker.Now()
  task.Modified = task.Created

  sql := `INSERT INTO task (title, status, created, modified) ` +
         `VALUES (:title, :status, :created, :modified);`

  result, err := db.NamedExecContext(ctx, sql, task)
  if err != nil {
    return err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return err
  }

  task.ID = entity.TaskID(id)
  return nil
}

