package repository

import (
	"testing"
	"context"
	"time"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository/kvs/ttlmap"
)

func TestSimpleKVS(t *testing.T) {
	m := ttlmap.NewTTLMap()
	k := NewSimpleKVS(m)

	ctx := context.Background()
	if err := k.SetUserID(ctx, "aaa", model.UserID(1), 60) ; err != nil {
		t.Errorf("unexpected error")
	}
	if err := k.SetUserID(ctx, "bbb", model.UserID(2), 1) ; err != nil {
		t.Errorf("unexpected error")
	}

	time.Sleep(time.Second * 2)

	if _, err := k.GetUserID(ctx, "aaa") ; err != nil {
		t.Errorf("key is expired")
	}
	if _, err := k.GetUserID(ctx, "bbb") ; err == nil {
		t.Errorf("key is alived")
	}
}
