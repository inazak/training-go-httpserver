package store

import (
  "context"
  "os"
  "testing"
  "github.com/inazak/training-go-httpserver/config"
)

var testdbpath = "./todo_test.db"

func createDBForTest(t *testing.T) (*DB, error) {

	t.Setenv("TODO_DBPATH", testdbpath)

  cfg, err := config.New()
  if err != nil {
    return nil, err
  }

  db, err := New(context.Background(), cfg)
  if err != nil {
    return nil, err
  }

  return db, nil
}

func removeDBForTest() {
  os.Remove(testdbpath)
}

func TestDBOpen(t *testing.T) {

  db, err := createDBForTest(t)
  if err != nil {
    t.Fatalf("creating new config failed: %v", err)
  }

	t.Cleanup( func() {
    db.Close()
    removeDBForTest()
  })
}

