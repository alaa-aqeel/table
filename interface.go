package table

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type SquirrelBuilder interface {
	ToSql() (string, []any, error)
}

type IDatabase interface {
	Db() *sql.DB
	QueryRow(ctx context.Context, sql string, args ...any) *sql.Row
	Query(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) error
}

type IRow interface {
	Scan(dest ...any) error
	Err() error
}

type IRows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Columns() ([]string, error)
	Close() error
}

type ISquirrel interface {
	Query() squirrel.SelectBuilder
	Statement() squirrel.StatementBuilderType
}

type IDb interface {
	Row(ctx context.Context, query SquirrelBuilder) (IRow, error)
	Rows(ctx context.Context, query SquirrelBuilder) (IRows, error)
	Exec(ctx context.Context, query SquirrelBuilder) error
}

type ITableFetch interface {
	FindPk(ctx context.Context, pks any) (IRow, error)
	Find(ctx context.Context, key string, value any) (IRow, error)
	Paginate(ctx context.Context, limit, offset int, wheres any) (IRows, error)
	Count(ctx context.Context, wheres any) (int64, error)
	Get(ctx context.Context, wheres any) (IRows, error)
}

type ITableInsert interface {
	Insert(ctx context.Context, data map[string]any) error
	InsertPk(ctx context.Context, data map[string]any) (any, error)
	InsertMany(ctx context.Context, cols []string, values []map[string]any) error
}

type ITableUpdate interface {
	Update(ctx context.Context, wheres any, data map[string]any) error
	UpdatePk(ctx context.Context, pks any, data map[string]any) error
}

type ITableDelete interface {
	Delete(ctx context.Context, wheres any) error
	DeletePk(ctx context.Context, pks any) error
}

type ITable interface {
	IDb
	ISquirrel
	ITableFetch
	ITableInsert
	ITableUpdate
	ITableDelete

	SetPk(value string) ITable
	SetTableName(value string) ITable
	Column(cols string) ITable
}
