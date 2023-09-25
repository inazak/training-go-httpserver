package repository

import (
	"context"
	"github.com/inazak/training-go-httpserver/common/config"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository/database/sqlite"
	"os"
	"sort"
	"testing"
)

var testDBPath = "./todo_test.db"

func TestUseSqliteDB(t *testing.T) {

	t.Setenv("TODO_DBPATH", testDBPath)

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("creating new config failed: %v", err)
	}

	db, err := sqlite.NewDatabase(context.Background(), cfg)
	if err != nil {
		t.Fatalf("creating new database failed: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
		os.Remove(testDBPath)
	})

	ps := model.TaskList{
		{UserID: 1, Title: "Hydrogen", Status: model.TaskStatusTodo},
		{UserID: 1, Title: "Helium", Status: model.TaskStatusDone},
		{UserID: 1, Title: "Lithium", Status: model.TaskStatusTodo},
	}

	ctx := context.Background()
	sd := NewSimpleDB(db)

	sd.InsertUser(ctx, &model.User{Name: "firstuser"})

	user, err := sd.SelectUser(ctx, "firstuser")
	if err != nil {
		t.Fatalf("failed SelectUser: %v", err)
	}
	if user.ID != 1 {
		t.Fatalf("unmatch UserID: expect=%d, got=%d", 1, user.ID)
	}

	for i, p := range ps {
		err = sd.InsertTask(ctx, p)
		if err != nil {
			t.Errorf("failed no.%d AddTask: %v", i, err)
		}
	}

	rs, err := sd.SelectTaskList(ctx, model.UserID(1))
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
