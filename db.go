package table

import "context"

func (t *SqlTable) Row(ctx context.Context, query SquirrelBuilder) (IRow, error) {
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	return t.db.QueryRow(ctx, sql, args...), nil
}

func (t *SqlTable) Rows(ctx context.Context, query SquirrelBuilder) (IRows, error) {
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
