package table

import (
	"context"
	"fmt"
	"sort"
)

func (t *SqlTable) Insert(ctx context.Context, data map[string]any) error {

	return t.Exec(ctx, t.statement.Insert(t.tableName).SetMap(data))
}

func (t *SqlTable) InsertPk(ctx context.Context, data map[string]any) (any, error) {
	var pk any
	row, err := t.Row(ctx, t.statement.Insert(t.tableName).SetMap(data).Suffix("RETURNING \""+t.pk+"\""))
	if err != nil {
		return nil, err
	}
	err = row.Scan(&pk)

	return pk, err
}

func (t *SqlTable) InsertMany(ctx context.Context, cols []string, values []map[string]any) error {

	q := t.statement.Insert(t.tableName)
	sort.Strings(cols)
	q = q.Columns(cols...)

	for _, value := range values {
		vals := make([]any, 0, len(cols))
		for _, col := range cols {
			vals = append(vals, value[col])
		}
		fmt.Println(cols)
		fmt.Println(vals)
		q = q.Values(vals...)
	}

	return t.Exec(ctx, q)
}
