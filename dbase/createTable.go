package dbase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTable(
	pool *pgxpool.Pool,
	table string,
	schema string,
	columns []Column,
	primaryKey []string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query, err := RenderTemplateFromFile(
		"./templates/create_table.sql.tmpl",
		TableTemplateData{
			Schema:     schema,
			TableName:  table,
			Columns:    columns,
			PrimaryKey: primaryKey,
		},
	)

	tag, err := pool.Exec(ctx, query)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Printf("table %s.%s created successfully, affected rows: %d\n", schema, table, tag.RowsAffected())
}
