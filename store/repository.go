package store

import (
  "context"
  "database/sql"
  "fmt"
  "time"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jmoiron/sqlx"
  "github.com/inazak/training-go-httpserver/config"
  "github.com/inazak/training-go-httpserver/clock"
)

type DB struct {
  SqlDB  *sql.DB
  SqlxDB *sqlx.DB
}

func (d *DB) Close() {
  if d.SqlDB  != nil { d.SqlDB.Close() }
  if d.SqlxDB != nil { d.SqlxDB.Close() }
}


// 二つ目の戻り値は *sql.DB.Close を実行する関数を返す
func New(ctx context.Context, cfg *config.Config) (*DB, error) {

  db := &DB{}

  sqldb, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=1", cfg.DBPath))
  if err != nil {
    return nil, err
  }
  db.SqlDB = sqldb

  ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
  defer cancel()
  if err := db.SqlDB.PingContext(ctx); err != nil {
    return db, err
  }

  err = RunMigrateUp(db.SqlDB)
  if err != nil {
    return db, err
  }

  sqlxdb := sqlx.NewDb(db.SqlDB, "sqlite3")
  db.SqlxDB = sqlxdb

  return db, nil
}


// 後にある _ Beginner = (*sqlx.DB)(nil) が成り立つのは
// *sqlx.DB が下記だから
//
// type DB struct {
// 	*sql.DB
//   ...
// }
type Beginner interface {
  BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Preparer interface {
  PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

type Execer interface {
  ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
  NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Queryer interface {
  Preparer
  QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
  QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
  GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
  SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

var (
  _ Beginner = (*sqlx.DB)(nil)
  _ Preparer = (*sqlx.DB)(nil)
  _ Execer   = (*sqlx.DB)(nil)
  _ Execer   = (*sqlx.Tx)(nil)
  _ Queryer  = (*sqlx.DB)(nil)
)

type Repository struct {
  Clocker clock.Clocker
}

