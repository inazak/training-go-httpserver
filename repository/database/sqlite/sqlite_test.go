package sqlite

import (
  "testing"
  "context"
  "os"
  "github.com/inazak/training-go-httpserver/common/config"
)

var testDBPath = "./todo_test.db"

func TestDatabaseOpen(t *testing.T) {

	t.Setenv("TODO_DBPATH", testDBPath)

  cfg, err := config.NewConfig()
  if err != nil {
    t.Fatalf("creating new config failed: %v", err)
  }

  db, err := NewDatabase(context.Background(), cfg)
  if err != nil {
    t.Fatalf("creating new database failed: %v", err)
  }

	t.Cleanup( func() {
    db.Close()
    os.Remove(testDBPath)
  })
}


