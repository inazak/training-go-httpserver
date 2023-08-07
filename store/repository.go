package store

import (
  "context"
  "database/sql"
  "fmt"
  "time"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jmoiron/sqlx"
  "github.com/inazak/training-go-httpserver/config"
)


// 二つ目の戻り値は *sql.DB.Close を実行する関数を返す
func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {

  db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=1", cfg.DBPath))
  if err != nil {
    return nil, nil, err
  }

  ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
  defer cancel()
  if err := db.PingContext(ctx); err != nil {
    return nil, func() { _ = db.Close() }, err
  }

  err = RunMigrateUp(db)
  if err != nil {
    return nil, func() { _ = db.Close() }, err
  }

  xdb := sqlx.NewDb(db, "sqlite3")
  return xdb, func() { _ = db.Close() }, nil
}


