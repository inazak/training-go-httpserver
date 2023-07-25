package main

import (
  "context"
  "fmt"
  "io"
  "net/http"
  "testing"
  "golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {

  // 後でキャンセルするために生成
  ctx, cancel := context.WithCancel(context.Background())

  // run関数を別goroutineで起動しておく
  eg, ctx := errgroup.WithContext(ctx)
  eg.Go( func() error {
    return run(ctx)
  })

  msg := "message"
  rsp, err := http.Get("http://localhost:18080/" + msg)
  if err != nil {
    t.Errorf("failed to get: %+v", err)
  }
  defer rsp.Body.Close()

  got, err := io.ReadAll(rsp.Body)
  if err != nil {
    t.Fatalf("failed to read body: %v", err)
  }

  // httpサーバの応答を検証する
  expect := fmt.Sprintf("Hello, %s!", msg)
  if string(got) != expect {
    t.Errorf("expect %q, but got %q", expect, got)
  }

  // 別goroutineで起動しているrun関数に終了通知
  cancel()

  // 別goroutineで起動しているrun関数が終了して終わる
  if err := eg.Wait(); err != nil {
    t.Fatal(err)
  }
}

