package main

import (
  "context"
  "os"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
  "github.com/inazak/training-go-httpserver/common/config"
  "github.com/inazak/training-go-httpserver/common/logging"
  "github.com/inazak/training-go-httpserver/httpserver"
  "github.com/inazak/training-go-httpserver/service"
  "github.com/inazak/training-go-httpserver/repository"
  "github.com/inazak/training-go-httpserver/repository/database/sqlite"
)

func main() {
  logger := logging.NewLogger()

  if err := runHttpServer(logger); err != nil {
    level.Error(logger).Log("msg", "failed to runHttpServer", "err", err)
    os.Exit(1)
  }
}


func runHttpServer(logger log.Logger) error {

  ctx := context.Background()

  conf, err := config.NewConfig()
  if err != nil {
    return err
  }

  level.Info(logger).Log("msg", "open sqlite db")
  db, err := sqlite.NewDatabase(ctx, conf)
  if err != nil {
    return err
  }
  defer db.Close()

  rep := repository.NewSimpleDB(db)
  svc := service.NewTodoService(ctx, rep, logger)
  mux := httpserver.NewMux(svc)
  hsv := httpserver.NewHttpServer(mux)

  return hsv.Run(ctx, conf.Port, logger)
}

