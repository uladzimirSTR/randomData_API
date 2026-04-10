package dbase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	obj "github.com/uladzimirSTR/randomData_API/objects"
)

func GetUsers(
	pool *pgxpool.Pool,
	schema string,
	tableName string,
	params map[string]string,
) ([]string, error) {

	data := SelectTemplateData{
		Schema:    schema,
		TableName: tableName,
		DateType:  params["dateCol"],
		Start:     params["start"],
		End:       params["end"],
		Limit:     10,
		Offset:    0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query, err := RenderTemplateFromFile("./templates/select_users.sql.tmpl", data)
	if err != nil {
		log.Fatalf("render template: %v", err)
	}

	rows, err := pool.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		log.Fatalf("query execution failed: %v", err)
	}

	var users []string

	for rows.Next() {
		var user obj.User
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
