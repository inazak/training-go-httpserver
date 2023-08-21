package main

import (
  "context"
  "fmt"
  "net/http"
  "os"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
  "github.com/inazak/training-go-httpserver/common/config"
  "github.com/inazak/training-go-httpserver/common/logging"
  "github.com/inazak/training-go-httpserver/server"
)

func run(ctx context.Context, logger log.Logger) error {

  conf, err := config.NewConfig()
  if err != nil {
    return err
  }

  mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
  })

  hs := server.NewHttpServer(mux)
  err = hs.Listen(conf.Port)
  if err != nil {
    level.Error(logger).Log("msg", "failed to listen port", "port", conf.Port, "err", err)
  }

  return hs.Run(ctx, logger)
}

func main() {
  logger := logging.NewLogger()

  if err := run(context.Background(), logger); err != nil {
    level.Error(logger).Log("msg", "failed server.run", "err", err)
    os.Exit(1)
  }
}

