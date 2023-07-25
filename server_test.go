package main

import (
  "context"
  "fmt"
  "io"
  "log"
  "net"
  "net/http"
  "testing"
  "golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {

  // ポート番号に0を選択すると利用可能なポートを動的に選択する
  l, err := net.Listen("tcp", "localhost:0")
  if err != nil {
    log.Fatalf("failed to listen port: %v", err)
  }

  // 後でキャンセルするために生成
  ctx, cancel := context.WithCancel(context.Background())

  mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
  })

  // run関数を別goroutineで起動しておく
  eg, ctx := errgroup.WithContext(ctx)
  eg.Go( func() error {
    s := NewServer(l, mux)
    return s.Run(ctx)
  })

  msg := "message"
  url := fmt.Sprintf("http://%s/%s", l.Addr().String(), msg)
  log.Printf("try request to %q", url)

  rsp, err := http.Get(url)
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

