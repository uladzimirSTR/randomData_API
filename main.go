package main

import (
	"context"
	"time"

	db "github.com/uladzimirSTR/randomData_API/dbase"
	rnd "github.com/uladzimirSTR/randomData_API/randomData"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	db.CreateTable(ctx, "users", "random_data", []db.Column{
		{Name: "id", Type: "BIGSERIAL", NotNull: true},
		{Name: "email", Type: "TEXT", NotNull: true},
		{Name: "name", Type: "TEXT"},
		{Name: "created_at", Type: "TIMESTAMP", NotNull: true, Default: "NOW()"},
		{Name: "updated_at", Type: "TIMESTAMP", NotNull: true, Default: "NOW()"},
	}, []string{"id"})

	users := rnd.GenerateRandomUsers()
	// fmt.Printf("%+v\n", users)

	db.InsertUsers(ctx, "random_data", "users", users)
}
