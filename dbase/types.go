package dbase

type Column struct {
	Name    string
	Type    string
	NotNull bool
	Default string
}

type TableTemplateData struct {
	Schema     string
	TableName  string
	Columns    []Column
	PrimaryKey []string
}

type InsertTemplateData struct {
	Schema    string
	TableName string
	Rows      []string
}
