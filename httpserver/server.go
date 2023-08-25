package httpserver

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer struct {
	*http.Server
}

func NewHttpServer(mux http.Handler) *HttpServer {
	return &HttpServer{
		Server: &http.Server{Handler: mux},
	}
}

func (hs *HttpServer) Run(ctx context.Context, port int, logger log.Logger) error {

	// シグナルを受け取るcontext、CTRL-Cを受け取るとDoneが呼ばれる
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// ポート番号に0を選択すると利用可能なポートを動的に選択する
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)

	// 別goroutineでhttpserverを起動する
	eg.Go(func() error {
		if err := hs.Serve(l); err != nil {
			// ErrServerClosedは http.Server.Shutdown()が正常終了なので、異常ではない
			if err != http.ErrServerClosed {
				return err
			}
		}
		return nil
	})

	level.Info(logger).Log("msg", "start server", "addr", l.Addr().String())

	// run関数の呼び出し元がcontextを使って終了を指示した場合
	<-ctx.Done()
	if err := hs.Shutdown(context.Background()); err != nil {
		level.Error(logger).Log("msg", "failed to shutdown", "err", err)
	}

	return eg.Wait() // 戻り値は eg.Go()で起動していた無名関数の戻り値
	// errgroup は全てのgoroutineが終了するまで待つ、errorがあった場合は
	// goroutineの中で最初のerrorを返した値を戻り値とする
}
