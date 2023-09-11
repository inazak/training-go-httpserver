package httpserver

import (
	"github.com/golang/mock/gomock"
	"github.com/inazak/training-go-httpserver/common/mock"
	"testing"
)

func TestMux(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockService(ctrl)
	mux := NewMux(m)

	if mux == nil {
		t.Errorf("failed to create")
	}
}
