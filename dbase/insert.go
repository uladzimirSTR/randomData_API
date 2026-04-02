package dbase

import (
	"fmt"
	"strings"

	randomdata "github.com/uladzimirSTR/randomData_API/randomData"
)

func BuildUsersInsertQuery(
	templatePath string,
	schema string,
	tableName string,
	users []randomdata.User,
) (string, []any, error) {
	if len(users) == 0 {
		return "", nil, fmt.Errorf("users slice is empty")
	}

	args := make([]any, 0, len(users)*5)
	rows := make([]string, 0, len(users))

	placeholder := 1

	for _, user := range users {
		row := fmt.Sprintf(
			"$%d, $%d, $%d, $%d, $%d",
			placeholder,
			placeholder+1,
			placeholder+2,
			placeholder+3,
			placeholder+4,
		)

		rows = append(rows, row)

		args = append(args,
			user.ID,
			user.Email,
			user.Name,
			user.CreatedAt,
			user.UpdatedAt,
		)

		placeholder += 5
	}

	data := InsertTemplateData{
		Schema:    schema,
		TableName: tableName,
		Rows:      rows,
	}

	query, err := RenderTemplateFromFile(templatePath, data)
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func NormalizeRows(rows []string) []string {
	out := make([]string, 0, len(rows))
	for _, row := range rows {
		out = append(out, strings.TrimSpace(row))
	}
	return out
}
