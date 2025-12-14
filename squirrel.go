package table

import "github.com/Masterminds/squirrel"

func (t *SqlTable) Statement() squirrel.StatementBuilderType {

	return t.statement
}

func (t *SqlTable) Query() squirrel.SelectBuilder {

	return t.statement.Select().From(t.tableName)
}
