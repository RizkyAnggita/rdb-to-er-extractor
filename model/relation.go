package model

type Database struct {
	Tables []Table
}

type ForeignKey struct {
	ColumnName           string
	ReferencedTableName  string
	ReferencedColumnName string
}

type PrimaryKey struct {
	ColumnName string
}

type Table struct {
	Name        string
	Columns     []Column
	PrimaryKeys []PrimaryKey
	ForeignKeys []ForeignKey
	Type        string
}

type Column struct {
	Name     string
	Type     string
	Nullable bool
}
