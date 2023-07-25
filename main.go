package main

import (
  "context"
  "fmt"
  "log"
  "net"
  "net/http"
  "os"
  "golang.org/x/sync/errgroup"
  "github.com/inazak/training-go-httpserver/config"
)

func main() {
  if err := run(context.Background()); err != nil {
    log.Printf("failed to terminate server: %v", err)
    os.Exit(1)
  }
}


func run(ctx context.Context) error {

  // 環境変数を取得する
  cfg, err := config.New()
  if err != nil {
    return err
  }

  // listenするがもっと後ろでもよくないか
  l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
  if err != nil {
    return fmt.Errorf("failed to listen port %d: %v", cfg.Port, err)
  }

  url := fmt.Sprintf("http://%s", l.Addr().String())
  log.Printf("start with: %s", url)

  // *http.Server には他も設定項目がある
  s := &http.Server {
    // net.Listenerを使うので Addr を設定しない
    Handler: http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
      fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    }),
  }

  // ここでrun関数の呼び出し元のcontextを引きついでいるが
  // それがDone()した場合は影響があるのか、分からない
  eg, ctx := errgroup.WithContext(ctx)

  // 別goroutineでhttpserverを起動する
  eg.Go( func() error {
    if err := s.Serve(l); err != nil {
      // ErrServerClosedは http.Server.Shutdown()が正常終了なので、異常ではない
      if err != http.ErrServerClosed {
        log.Printf("failed to close: %v", err)
        return err
      }
    }
    return nil
  })

  // run関数の呼び出し元がcontextを使って終了を指示した場合
  <-ctx.Done()
  if err := s.Shutdown(context.Background()); err != nil {
    log.Printf("failed to shutdown: %v", err)
  }

  return eg.Wait() // 戻り値は eg.Go()で起動していた無名関数の戻り値
  // errgroup は全てのgoroutineが終了するまで待つ、errorがあった場合は
  // goroutineの中で最初のerrorを返した値を戻り値とする
}

