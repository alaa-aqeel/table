# Table 
This package provides a simple way to interact with a database using the sql package.
It includes a type `SqlTable` which allows you to perform CRUD operations on a table.
The `SqlTable` type has methods to insert, update, delete and fetch data from the table.
It also includes a `Dto` type which is used to define the structure of the data that is fetched from the table.
The `UserTable` type is an example of how to use the `SqlTable` type to interact with a specific table in the database.

## Dto 
```go
type UserDto struct {
	Name     string
	Username string
}
```

## User Table 
```go


type UserTable struct {
	*table.SqlTable
}

func User(db database.IDatabase) *UserTable {

	return &UserTable{
		SqlTable: table.Table(db, "users", "id"),
	}
}

func (u *UserTable) GetByUsername(username string) (UserDto, error) {
	row, err := u.SqlTable.One(context.Background(), "username", username)
	if err != nil {
		return UserDto{}, err
	}

	var dto UserDto
	row.Scan(&dto.Name, &dto.Username)

	return dto, nil
}
```

## Database Interface 
```go
type IDatabase interface {
	Db() *sql.DB
	QueryRow(ctx context.Context, sql string, args ...any) *sql.Row
	Query(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) error
}
```

## Main 
```go
import (
    "github.com/alaa-aqeel/example/database"
	"github.com/alaa-aqeel/example/models/user"
	_ "github.com/jackc/pgx/v5/stdlib"
)

db := database.NewDatabase()
err = db.Connect(ctx, "pgx", os.Getenv("DATABASE_URL"))
if err != nil {
    log.Fatal("[database]: " + err.Error())
}
defer db.Close()
ctx := context.Background()

user := user.User(db)
err = user.Delete(ctx, map[string]any{
    "username": "2001",
})
if err != nil {
    log.Fatal(err)
}
data, err := user.GetByUsername("2001")
```

