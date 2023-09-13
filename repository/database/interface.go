package database

import (
	"context"
)

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Database interface {
	Close()
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(ctx context.Context, query string, arg interface{}) (Result, error)
}
