package table

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (t *SqlTable) One(ctx context.Context, key string, value any) (IRow, error) {
	row, err := t.Row(ctx,
		t.Query().Where(squirrel.Eq{key: value}),
	)

	return row, err
}

func (t *SqlTable) Find(ctx context.Context, pks any) (IRow, error) {
	return t.One(ctx, t.pk, pks)
}

func (t *SqlTable) All(ctx context.Context, limit, offset int, wheres map[string]any) (IRows, error) {
	rows, err := t.Rows(ctx,
		t.
			Query().
			Where(wheres).
			Limit(uint64(limit)).
			Offset(uint64(offset)),
	)

	return rows, err
}

func (t *SqlTable) Filter(ctx context.Context, limit, offset int, wheres any) (IRows, error) {
	rows, err := t.Rows(ctx,
		t.
			Query().
			Where(wheres).
			Limit(uint64(limit)).
			Offset(uint64(offset)),
	)

	return rows, err
}

func (t *SqlTable) Count(ctx context.Context, wheres any) (int64, error) {
	row, err := t.Row(ctx,
		t.
			statement.
			Select("COUNT(*)").
			From("users").
			Where(wheres),
	)

	if err != nil {
		return 0, err
	}

	var total int64
	err = row.Scan(&total)

	return total, err
}
