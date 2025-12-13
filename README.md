# Table 
This package provides a simple way to interact with a database using the sql package.
It includes a type `SqlTable` which allows you to perform CRUD operations on a table.
The `SqlTable` type has methods to insert, update, delete and fetch data from the table.
It also includes a `Dto` type which is used to define the structure of the data that is fetched from the table.
The `UserTable` type is an example of how to use the `SqlTable` type to interact with a specific table in the database.


# Install
```sh
$ go get github.com/alaa-aqeel/table@v0.0.1
```


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



# Table Methods 

## Delete
```go
user := user.User(db)

---- delete 
err := user.Delete(ctx, map[string]any{
    "username": "2001",
})
err := user.DeletePk(ctx, "id")
```


## Update 
```go
--- update 
err := user.Update(ctx, map[string]any{ // wheres username == 2001
    "username": "2001",
}, map[string]any{ //  data
    "name": "alaa aqeel",
})

err := user.UpdatePk(ctx, "id", map[string]any{ //  data
    "name": "alaa aqeel",
})
```

##  Insert 
```go
---- Insert 
err := user.Insert(ctx, map[string]any{ //  data
	"id": "uuid"
    "name": "alaa aqeel",
	"username": "alaa_aqeel"
})
pk, err := user.InsertPk(ctx, map[string]any{ //  data
	"id": "uuid"
    "name": "alaa aqeel",
	"username": "alaa_aqeel"
})

err = user.InsertMany(ctx, []string{"name", "username"}, []map[string]any{
	{
		"name":     "aqeel",
		"username": "alaa",
	},
	{
		"name":     "alaa",
		"username": "aqeel",
	},
})
```

## Fetch 
```go 
row, err = user.Find(ctx, "pk value") // find by pk
row, err = user.One(ctx, "key", "value") 
row, err = user.All(ctx, "limit int", "offset int", map[string]any{
	// wheres
})
```

## Custom 
Use [Masterminds/squirrel](github.com/Masterminds/squirrel) to build query
```go
func (u *UserTable) GetAll(ctx context.Context) ([]UserDto, error) {

	q := u.SqlTable.
		Query().
		Columns("name", "username").
		Where(squirrel.Eq{"active": true})

	rows, err := u.SqlTable.Rows(ctx, q)

	if err != nil {
		return nil, err
	}

	var users []UserDto
	for rows.Next() {
		var user UserDto
		err = rows.Scan(&user.Name, &user.Username)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
```