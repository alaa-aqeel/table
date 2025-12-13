package table

import (
	"context"
)

func (t *SqlTable) Insert(ctx context.Context, data map[string]any) (any, error) {
	var pk any
	row, err := t.Row(ctx, t.statement.Insert(t.tableName).SetMap(data))
	if err != nil {
		return nil, err
	}
	err = row.Scan(&pk)

	return pk, err
}

func (t *SqlTable) InsertMany(ctx context.Context, cols []string, values [][]any) error {

	q := t.statement.Insert(t.tableName)
	q.Columns(cols...)
	for _, value := range values {
		q.Values(value...)
	}

	return t.Exec(ctx, q)
}
