package dbase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	randomdata "github.com/uladzimirSTR/randomData_API/randomData"
)

func InsertUsers(
	pool *pgxpool.Pool,
	schema string,
	tableName string,
	users []randomdata.User,
) {
	if len(users) == 0 {
		log.Fatalf("users slice is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows := make([]string, 0, len(users))

	for _, user := range users {
		row := fmt.Sprintf(
			"%d, '%s', '%s', '%s', '%s'",
			user.ID,
			user.Email,
			user.Name,
			user.CreatedAt,
			user.UpdatedAt,
		)

		rows = append(rows, row)
	}

	data := InsertTemplateData{
		Schema:    schema,
		TableName: tableName,
		Rows:      rows,
	}

	query, err := RenderTemplateFromFile("./templates/insert_users.sql.tmpl", data)
	if err != nil {
		log.Fatalf("render template: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("ping database: %v", err)
	}

	tag, err := pool.Exec(ctx, query)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Printf("users inserted successfully, affected rows: %d\n", tag.RowsAffected())
}
