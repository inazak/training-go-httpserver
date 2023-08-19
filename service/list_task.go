package service

import (
  "context"
  "fmt"

  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/store"
)

// Queryer はDBを隠しているはずなので、DB という名前が美しくない
type ListTask struct {
  DB   store.Queryer
  Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
  ts, err := l.Repo.ListTasks(ctx, l.DB)
  if err != nil {
    return nil, fmt.Errorf("failed to list: %w", err)
  }
  return ts, nil
}

