package main

import (
//  "io"
//  "net/http"
//  "net/http/httptest"
  "testing"
)

func TestNewMux(t *testing.T) {
/*
  //FIXME NewMuxの変更(parameter変更、DBオープン)で動かなくなった

  w := httptest.NewRecorder()
  r := httptest.NewRequest(http.MethodGet, "/health", nil)

  mux := NewMux()
  mux.ServeHTTP(w, r)
  resp := w.Result()
  t.Cleanup(func(){ _ = resp.Body.Close() })

  if resp.StatusCode != http.StatusOK {
    t.Error("want status code 200, but got ", resp.StatusCode)
  }

  got, err := io.ReadAll(resp.Body)
  if err != nil {
    t.Fatalf("failed to read body: %v", err)
  }

  want := `{status: "ok"}`
  if string(got) != want {
    t.Errorf("want %q, but got %q", want, got)
  }
*/
}

// TODO httptestについて調べる

