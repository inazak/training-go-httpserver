package store

import (
  "context"
  "os"
  "testing"
  "github.com/inazak/training-go-httpserver/config"
)

func TestDBOpen(t *testing.T) {

  dbpath := "./todo_test.db"

	t.Setenv("TODO_DBPATH", dbpath)

  cfg, err := config.New()
  if err != nil {
    t.Fatalf("creating new config failed: %v", err)
  }

  _, closer, err := New(context.Background(), cfg)
  if err != nil {
    t.Fatalf("db open err: %v", err)
  }

  closer()
  os.Remove(dbpath)
}

