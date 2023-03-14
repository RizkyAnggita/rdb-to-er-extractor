package model

type ForeignKey struct {
	ColumnName           string
	ReferencedTableName  string
	ReferencedColumnName string
}
