package main

import (
  "context"
  "fmt"
  "log"
  "net"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "golang.org/x/sync/errgroup"
)

type Server struct {
  srv *http.Server
  l   net.Listener
}

// mux は multiplexer の略
func NewServer(l net.Listener, mux http.Handler) *Server {
  return &Server {
    srv: &http.Server{ Handler: mux},
    l: l,
  }
}


func (s *Server) Run(ctx context.Context) error {

  // シグナルを受け取るcontext、CTRL-Cを受け取るとDoneが呼ばれる
  ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
  defer stop()

  eg, ctx := errgroup.WithContext(ctx)

  // 別goroutineでhttpserverを起動する
  eg.Go( func() error {
    if err := s.srv.Serve(s.l); err != nil {
      // ErrServerClosedは http.Server.Shutdown()が正常終了なので、異常ではない
      if err != http.ErrServerClosed {
        log.Printf("failed to close: %v", err)
        return err
      }
    }
    return nil
  })

  url := fmt.Sprintf("http://%s", s.l.Addr().String())
  log.Printf("start with: %s", url)

  // run関数の呼び出し元がcontextを使って終了を指示した場合
  <-ctx.Done()
  if err := s.srv.Shutdown(context.Background()); err != nil {
    log.Printf("failed to shutdown: %v", err)
  }

  return eg.Wait() // 戻り値は eg.Go()で起動していた無名関数の戻り値
  // errgroup は全てのgoroutineが終了するまで待つ、errorがあった場合は
  // goroutineの中で最初のerrorを返した値を戻り値とする
}

