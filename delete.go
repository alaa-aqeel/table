package table

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (t *SqlTable) Delete(ctx context.Context, wheres any) error {
	return t.Exec(ctx,
		t.statement.
			Delete(t.tableName).
			Where(wheres),
	)
}

func (t *SqlTable) DeletePk(ctx context.Context, pks any) error {

	return t.Delete(ctx, squirrel.Eq{t.pk: pks})
}
