package dbase

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	obj "github.com/uladzimirSTR/randomData_API/objects"
)

func GetUsers(
	pool *pgxpool.Pool,
	schema string,
	tableName string,
	params map[string]any,
) ([]obj.User, error) {

	var limit int = params["limit"].(int)
	var offset int = params["offset"].(int)

	data := SelectTemplateData{
		Schema:    schema,
		TableName: tableName,
		DateType:  params["dateCol"],
		Start:     params["start"],
		End:       params["end"],
		Limit:     limit,
		Offset:    offset,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query, err := RenderTemplateFromFile("./templates/select_users.sql.tmpl", data)
	if err != nil {
		return nil, fmt.Errorf("render template: %w", err)
	}

	rows, err := pool.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}

	defer rows.Close()

	var users []obj.User

	for rows.Next() {
		var user obj.User

		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		// users = append(users, fmt.Sprintf("%+v", user))
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}
