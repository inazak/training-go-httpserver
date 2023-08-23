package httpserver

import (
  "context"
  "fmt"
  "io"
  "net/http"
  "testing"
	"github.com/go-kit/log"
  "golang.org/x/sync/errgroup"
)

var logger = log.NewNopLogger()

func TestRun(t *testing.T) {

  // 後でキャンセルするために生成
  ctx, cancel := context.WithCancel(context.Background())

  mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
  })

  hs := NewHttpServer(mux)

  // ポート番号に0を選択すると利用可能なポートを動的に選択する
  err := hs.Listen(0)
  if err != nil {
    t.Fatalf("failed to listen: %v", err)
  }

  // run関数を別goroutineで起動しておく
  eg, ctx := errgroup.WithContext(ctx)
  eg.Go( func() error {
    return hs.Run(ctx, logger)
  })

  msg := "message"
  url := fmt.Sprintf("http://%s/%s", hs.lsnr.Addr().String(), msg)

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

