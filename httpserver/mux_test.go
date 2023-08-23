package httpserver

import (
  "context"
  "testing"
  "github.com/inazak/training-go-httpserver/common/mock"
  "github.com/golang/mock/gomock"
)

func TestMux(t *testing.T) {
  ctx := context.Background()

  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  m := mock.NewMockService(ctrl)
  mux := NewMux(ctx, m)

  if mux == nil {
    t.Errorf("failed to create")
  }
}

