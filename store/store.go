package store

import (
  "errors"
  "github.com/inazak/training-go-httpserver/entity"
)

var (
  Tasks = &TaskStore{ Tasks: map[entity.TaskID]*entity.Task{} }

  ErrNotFound = errors.New("not found")
)

type TaskStore struct {
  LastID entity.TaskID
  Tasks map[entity.TaskID]*entity.Task
}

func(ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
  ts.LastID += 1
  t.ID = ts.LastID
  ts.Tasks[t.ID] = t
  return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
  tasks := make([]*entity.Task, len(ts.Tasks))
  i := 0
  for _, t := range ts.Tasks {
    tasks[i] = t
    i += 1
  }
  return tasks
}

