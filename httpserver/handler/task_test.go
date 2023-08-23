package handler

import (
  "testing"
  "bytes"
  "net/http"
  "net/http/httptest"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/common/jsonhelper"
  "github.com/inazak/training-go-httpserver/common/mock"
  "github.com/inazak/training-go-httpserver/model"
  "github.com/golang/mock/gomock"
)

func TestAddTask(t *testing.T) {
  t.Parallel()

  ps := map[string]struct {
    prepareMock func(m *mock.MockService)
    request     string
    response    string
    status      int
  }{
    "ok": {
      prepareMock: func(m *mock.MockService){
        m.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(&model.Task{ID: 1}, nil)
      },
      request: `{ "title": "ok" }`,
      response: `{ "id": 1 }`,
      status: http.StatusOK,
    },
    "badRequest": {
      prepareMock: func(m *mock.MockService){
        // AddTaskが呼ばれないテストなので、EXPECT()しない
      },
      request: `{ "xxx": "ng" }`,
      response: `{ "message": "Key: 'Title' Error:Field validation for 'Title' failed on the 'required' tag" }`,
      status: http.StatusBadRequest,
    },
  }

  for n, p := range ps {
    p := p
    t.Run(n, func(t *testing.T) {
      t.Parallel()

      w := httptest.NewRecorder()
      r := httptest.NewRequest(
        http.MethodPost,
        "/task",
        bytes.NewReader([]byte(p.request)),
      )

      ctrl := gomock.NewController(t)
      defer ctrl.Finish()

      m := mock.NewMockService(ctrl)
      p.prepareMock(m)

      sut := TaskHandler{
        serv: m,
        vali: validator.New(),
      }
      sut.ServeHTTP(w, r)

      resp := w.Result()
      jsonhelper.AssertResponse(t, resp, p.status, []byte(p.response))
    })
  }
}

