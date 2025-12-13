package table

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

func (t *SqlTable) One(ctx context.Context, key string, value any) (*sql.Row, error) {
	row, err := t.Row(ctx,
		t.Query().Where(squirrel.Eq{key: value}),
	)

	return row, err
}

func (t *SqlTable) Find(ctx context.Context, pks any) (*sql.Row, error) {
	return t.One(ctx, t.pk, pks)
}

func (t *SqlTable) All(ctx context.Context, limit, offset int, wheres map[string]any) (*sql.Rows, error) {
	rows, err := t.Rows(ctx,
		t.
			Query().
			Where(wheres).
			Limit(uint64(limit)).
			Offset(uint64(offset)),
	)

	return rows, err
}
