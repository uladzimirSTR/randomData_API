package dbase

import (
	"context"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertValues(
	pool *pgxpool.Pool,
	schema string,
	tableName string,
	columns []string,
	primaryKey []string,
	rows []string,
) {
	if len(rows) == 0 {
		log.Fatalf("rows slice is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateColumns := make([]string, 0)

	for _, col := range columns {
		if !slices.Contains(append(primaryKey, "updated_at"), col) {
			updateColumns = append(updateColumns, col)
		}
	}

	query, err := RenderTemplateFromFile(
		"./templates/insert_values.sql.tmpl",
		InsertTemplateData{
			Schema:        schema,
			TableName:     tableName,
			Rows:          rows,
			Columns:       columns,
			UpdateColumns: updateColumns,
			PrimaryKey:    primaryKey,
		},
	)

	if err != nil {
		log.Fatalf("render template: %v", err)
	}

	tag, err := pool.Exec(ctx, query)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Printf("users inserted successfully, affected rows: %d\n", tag.RowsAffected())
}
