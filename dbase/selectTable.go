package dbase

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	rnd "github.com/uladzimirSTR/randomData_API/randomData"
)

func GetUsers(
	ctx context.Context,
	pool *pgxpool.Pool,
	schema string,
	tableName string,
	args ...any,
) ([]string, error) {
	data := SelectTemplateData{
		Schema:    schema,
		TableName: tableName,
		DateType:  args[0],
		Start:     args[1],
		End:       args[2],
	}

	fmt.Printf("%+v\n", data)

	query, err := RenderTemplateFromFile("./templates/select_users.sql.tmpl", data)
	if err != nil {
		log.Fatalf("render template: %v", err)
	}

	fmt.Printf("query: %s\n", query)

	rows, err := pool.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		log.Fatalf("query execution failed: %v", err)
	}

	var users []string

	for rows.Next() {
		var user rnd.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Fatalf("row scan failed: %v", err)
		}
		users = append(users, fmt.Sprintf("%+v", user))
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("rows iteration error: %v", err)
	}

	return users, nil
}
