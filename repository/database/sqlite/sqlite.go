package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/inazak/training-go-httpserver/common/config"
	"github.com/inazak/training-go-httpserver/repository/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//database.Database の実装
type DB struct {
	SqlDB  *sql.DB
	SqlxDB *sqlx.DB
}

func NewDatabase(ctx context.Context, cfg *config.Config) (*DB, error) {

	db := &DB{}

	sqldb, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=1", cfg.DBPath))
	if err != nil {
		return db, err
	}
	db.SqlDB = sqldb

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.SqlDB.PingContext(ctx); err != nil {
		return db, err
	}

	//FIXME どうするか
	err = RunMigrateUp(db.SqlDB)
	if err != nil {
		return db, err
	}

	sqlxdb := sqlx.NewDb(db.SqlDB, "sqlite3")
	db.SqlxDB = sqlxdb

	return db, nil
}

func (d *DB) Close() {
	// sqlx.DB の Close だけではダメなのだろうか
	if d.SqlDB != nil {
		d.SqlDB.Close()
	}
	if d.SqlxDB != nil {
		d.SqlxDB.Close()
	}
}

func (db *DB) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.SqlxDB.SelectContext(ctx, dest, query, args...)
}

func (db *DB) NamedExec(ctx context.Context, query string, arg interface{}) (database.Result, error) {
	result, err := db.SqlxDB.NamedExecContext(ctx, query, arg)
	if err != nil {
		return nil, err
	}
	return result, nil
}
