package table

import (
	"github.com/Masterminds/squirrel"
)

type SqlTable struct {
	tableName string
	pk        string
	cols      string
	statement squirrel.StatementBuilderType
	db        IDatabase
}

func Table(db IDatabase, tableName, pk string) ITable {

	return &SqlTable{
		db:        db,
		tableName: tableName,
		pk:        pk,
		cols:      "*",
		statement: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (t *SqlTable) SetPk(value string) ITable {
	t.pk = value

	return t
}

func (t *SqlTable) SetTableName(value string) ITable {
	t.tableName = value

	return t
}

func (t *SqlTable) Column(cols string) ITable {
	t.cols = cols

	return t
}
