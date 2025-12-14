package table

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (t *SqlTable) Get(ctx context.Context, wheres any) (IRows, error) {
	row, err := t.Rows(ctx,
		t.Query().Column(t.cols).Where(wheres),
	)

	return row, err
}

func (t *SqlTable) Find(ctx context.Context, key string, value any) (IRow, error) {
	row, err := t.Row(ctx,
		t.Query().Column(t.cols).Where(squirrel.Eq{key: value}).Limit(1),
	)

	return row, err
}

func (t *SqlTable) FindPk(ctx context.Context, pks any) (IRow, error) {
	return t.Find(ctx, t.pk, pks)
}

func (t *SqlTable) Paginate(ctx context.Context, limit, offset int, wheres any) (IRows, error) {
	rows, err := t.Rows(ctx,
		t.
			Query().
			Column(t.cols).
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
