package db

import (
	"context"
	"database/sql"
)

type DB interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Close() error
}
