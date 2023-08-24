package api

import (
  "testing"
  "bytes"
  "net/http"
  "net/http/httptest"
  "github.com/inazak/training-go-httpserver/common/jsonhelper"
  "github.com/inazak/training-go-httpserver/common/mock"
  "github.com/golang/mock/gomock"
)

func TestServeHealthCheck(t *testing.T) {

  request  := ""
  response := `{ "status": "ok" }`
  status   := http.StatusOK

  w := httptest.NewRecorder()
  r := httptest.NewRequest(
    http.MethodPost,
    "/health",
    bytes.NewReader([]byte(request)),
  )

  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  m := mock.NewMockService(ctrl)
  m.EXPECT().HealthCheck(gomock.Any()).Return(nil)

  h := NewHandler(m)
  h.ServeHealthCheck(w, r)

  resp := w.Result()
  jsonhelper.AssertResponse(t, resp, status, []byte(response))
}

