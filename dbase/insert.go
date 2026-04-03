package dbase

import (
	"context"
	"fmt"
	"log"

	randomdata "github.com/uladzimirSTR/randomData_API/randomData"
)

func InsertUsers(
	ctx context.Context,
	schema string,
	tableName string,
	users []randomdata.User,
) {
	if len(users) == 0 {
		log.Fatalf("users slice is empty")
	}

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

	if err := ExecuteSQL(ctx, query); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
