package dbimpl

import (
	"backend/pkg/infra/storage/db"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type sqlDB struct {
	db *sqlx.DB
}

func NewSQL(db *sqlx.DB) db.DB {
	return &sqlDB{db: db}
}

func (sq *sqlDB) Close() error {
	err := sq.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (sq *sqlDB) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return sq.db.GetContext(ctx, dest, sq.db.Rebind(query), args...)
}

func (sq *sqlDB) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return sq.db.SelectContext(ctx, dest, sq.db.Rebind(query), args...)
}

func (sq *sqlDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return sq.db.ExecContext(ctx, sq.db.Rebind(query), args...)
}

func (sq *sqlDB) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return sq.db.NamedExecContext(ctx, sq.db.Rebind(query), arg)
}
