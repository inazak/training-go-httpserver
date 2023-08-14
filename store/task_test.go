package store

import (
  "context"
  "sort"
  "testing"
  "github.com/inazak/training-go-httpserver/clock"
  "github.com/inazak/training-go-httpserver/entity"
)

func TestAddTaskAndListTask(t *testing.T) {

  db, err := createDBForTest(t)
  if err != nil {
    t.Fatalf("creating new config failed: %v", err)
  }

	t.Cleanup( func() {
    db.Close()
    removeDBForTest()
  })

  ps := entity.Tasks{
    { Title: "Hydrogen", Status: entity.TaskStatusTodo },
    { Title: "Helium",   Status: entity.TaskStatusDone },
    { Title: "Lithium",  Status: entity.TaskStatusTodo },
  }

  ctx := context.Background()
	c := clock.FixedClocker{}
  r := &Repository{ Clocker: c }
  for i, p := range ps {
    err = r.AddTask(ctx, db.SqlxDB, p)
    if err != nil {
      t.Errorf("failed no.%d AddTask: %v", i, err)
    }
  }

  rs, err := r.ListTasks(ctx, db.SqlxDB)
  if err != nil {
    t.Errorf("failed ListTask: %v", err)
  }

  sort.Slice(rs, func(i, j int) bool {
    return rs[i].ID < rs[j].ID
  })

  for i, r := range rs {
    if r.Title != ps[i].Title {
      t.Errorf("unmatch ListTask line %d, expect=%s, got=%s", i, ps[i].Title, r.Title)
    }
  }

}


