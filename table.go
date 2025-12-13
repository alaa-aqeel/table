package table

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type SquirrelBuilder interface {
	ToSql() (string, []any, error)
}

type SqlTable struct {
	tableName string
	pk        string
	cols      string
	statement squirrel.StatementBuilderType
	db        IDatabase
}

func Table(db IDatabase, tableName, pk string) *SqlTable {

	return &SqlTable{
		db:        db,
		tableName: tableName,
		pk:        pk,
		cols:      "*",
		statement: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (t *SqlTable) SetPk(value string) *SqlTable {
	t.pk = value

	return t
}

func (t *SqlTable) SetTableName(value string) *SqlTable {
	t.tableName = value

	return t
}

func (t *SqlTable) Query() squirrel.SelectBuilder {

	return t.statement.Select(t.cols).From(t.tableName)
}

func (t *SqlTable) Column(cols string) *SqlTable {
	t.cols = cols

	return t
}

func (t *SqlTable) Row(ctx context.Context, query SquirrelBuilder) (*sql.Row, error) {
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	return t.db.QueryRow(ctx, sql, args...), nil
}

func (t *SqlTable) Rows(ctx context.Context, query SquirrelBuilder) (*sql.Rows, error) {
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	return t.db.Query(ctx, sql, args...)
}

func (t *SqlTable) Exec(ctx context.Context, query SquirrelBuilder) error {
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	return t.db.Exec(ctx, sql, args...)
}
