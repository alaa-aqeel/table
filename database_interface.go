package table

import (
	"context"
	"database/sql"
)

type IDatabase interface {
	Db() *sql.DB
	QueryRow(ctx context.Context, sql string, args ...any) *sql.Row
	Query(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) error
}
