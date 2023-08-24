package httpserver

import (
  "testing"
  "github.com/inazak/training-go-httpserver/common/mock"
  "github.com/golang/mock/gomock"
)

func TestMux(t *testing.T) {

  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  m := mock.NewMockTodoService(ctrl)
  mux := NewMux(m)

  if mux == nil {
    t.Errorf("failed to create")
  }
}

