package jsonhelper

import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "testing"
  "github.com/google/go-cmp/cmp"
)

type ErrorResponse struct {
  Message string   `json:"message"`
  Detail []string  `json:"detail,omitempty"`
}

func WriteJSONResponse(ctx context.Context, w http.ResponseWriter, body any, status int) error {
  w.Header().Set("Context-Type", "application/json; charset=utf-8")

  bodyBytes, err := json.Marshal(body)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    rsp := ErrorResponse{
      Message: http.StatusText(http.StatusInternalServerError),
    }
    _ = json.NewEncoder(w).Encode(rsp)
    return fmt.Errorf("failed to marshal response")
  }

  w.WriteHeader(status)
  if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
    return fmt.Errorf("failed to marshal response")
  }

  return nil
}

func AssertJSON(t *testing.T, want, got []byte) {
  // ヘルパーメソッドとしてマーク、このメソッドの行番号を出さない
  t.Helper()

  var jw, jg any
  if err := json.Unmarshal(want, &jw); err != nil {
    t.Fatalf("cannot unmarshal want %q: %v", want, err)
  }
  if err := json.Unmarshal(got, &jg); err != nil {
    t.Fatalf("cannot unmarshal got %q: %v", got, err)
  }
  if diff := cmp.Diff(jg, jw); diff != "" {
    t.Errorf("got differs: (-got, +want)\n%s", diff)
  }
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
  t.Helper()
  t.Cleanup(func(){ _ = got.Body.Close() })

  gb, err := io.ReadAll(got.Body)
  if err != nil {
    t.Fatal(err)
  }
  if got.StatusCode != status {
    t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
  }

  if len(gb) == 0 && len(body) == 0 {
    return
  }
  AssertJSON(t, body, gb)
}

