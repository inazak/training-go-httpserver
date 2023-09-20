package repository

import (
	"context"
	"fmt"
	"github.com/inazak/training-go-httpserver/model"
	"github.com/inazak/training-go-httpserver/repository/kvs"
)

// repository.Memory の実装
type SimpleKVS struct {
	kvs.KVS
}

func NewSimpleKVS(k kvs.KVS) *SimpleKVS {
	return &SimpleKVS{
		KVS: k,
	}
}

func (k *SimpleKVS) SetUserID(ctx context.Context, key string, id model.UserID, ttl int64) error {
	k.Set(key, id, ttl)
	return nil
}

func (k *SimpleKVS) GetUserID(ctx context.Context, key string) (model.UserID, error) {
	v, ok := k.Get(key)
	if !ok {
		return -1, fmt.Errorf("expired or deleted key") //FIXME -1
	}
	id := v.(model.UserID)
	return id, nil
}
