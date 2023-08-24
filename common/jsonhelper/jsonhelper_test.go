package jsonhelper

import (
  "context"
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestWriteJSONResponse(t *testing.T) {

  w := httptest.NewRecorder()

  ctx := context.Background()
  rsp := struct{ Test string `json:"test"` }{ Test: "Hello" }
  WriteJSONResponse(ctx, w, rsp, http.StatusOK)

  resp := w.Result()
  AssertResponse(t, resp, http.StatusOK, []byte(`{ "test": "Hello"}`))
}

