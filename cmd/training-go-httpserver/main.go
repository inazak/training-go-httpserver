package main

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/inazak/training-go-httpserver/common/config"
	"github.com/inazak/training-go-httpserver/common/logging"
	"github.com/inazak/training-go-httpserver/common/jwter"
	"github.com/inazak/training-go-httpserver/httpserver"
	"github.com/inazak/training-go-httpserver/repository"
	"github.com/inazak/training-go-httpserver/repository/database/sqlite"
	"github.com/inazak/training-go-httpserver/service"
	"os"
	_ "embed"
)

// 秘密鍵と公開鍵の生成
// openssl genrsa 4096 > secret.pem
// openssl rsa -pubout < secret.pem  > public.pem

//go:embed cert/secret.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPublicKey []byte

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
	sqlitedb, err := sqlite.NewDatabase(ctx, conf)
	if err != nil {
		return err
	}
	defer sqlitedb.Close()
	db := repository.NewSimpleDB(sqlitedb)

	jtr, err := jwter.NewJWTer(`github.com/inazak/training-go-httpserver`, rawPrivateKey, rawPublicKey)
	if err != nil {
		return err
	}

	svc := service.NewTodoService(ctx, db, jtr, logger)
	mux := httpserver.NewMux(svc)
	hsv := httpserver.NewHttpServer(mux)

	return hsv.Run(ctx, conf.Port, logger)
}
