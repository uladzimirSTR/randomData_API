package dbase

import (
	"context"
	"log"
)

func CreateTable(
	ctx context.Context,
	table string,
	schema string,
	columns []Column,
	primaryKey []string,
) {
	query, err := RenderTemplateFromFile("./templates/create_table.sql.tmpl", TableTemplateData{
		Schema:     schema,
		TableName:  table,
		Columns:    columns,
		PrimaryKey: primaryKey,
	})

	if err != nil {
		log.Fatalf("render template: %v", err)
	}

	if err := ExecuteSQL(ctx, query); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

}
