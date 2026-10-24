package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64  `bun:"id,pk,autoincrement"`
	Name  string `bun:"name,notnull"`
	Email string `bun:"email,unique"`
}

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://admin:admin@localhost:5432/mydatabase?sslmode=disable")))
	db := bun.NewDB(sqldb, pgdialect.New())

	defer db.Close()

	ctx := context.Background()

	// Define model
	type User struct {
		ID   int64  `bun:",pk,autoincrement"`
		Name string `bun:",notnull"`
	}

	// Create table
	db.NewCreateTable().Model((*User)(nil)).Exec(ctx)

	// Insert user
	user := &User{Name: "John Doe"}
	db.NewInsert().Model(user).Exec(ctx)

	// Query user
	err := db.NewSelect().Model(user).Where("id = ?", user.ID).Scan(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User: %+v\n", user)

}
