package table

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (t *SqlTable) Update(ctx context.Context, wheres any, data map[string]any) error {
	return t.Exec(ctx,
		t.statement.
			Update(t.tableName).
			Where(wheres).
			SetMap(data),
	)
}

func (t *SqlTable) UpdatePk(ctx context.Context, pks any, data map[string]any) error {
	return t.Update(ctx, squirrel.Eq{t.pk: pks}, data)
}
