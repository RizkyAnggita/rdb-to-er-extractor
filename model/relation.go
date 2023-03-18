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

type DanglingKey struct {
	ColumnName string
}

type Table struct {
	Name         string
	Columns      []Column
	PrimaryKeys  []PrimaryKey
	ForeignKeys  []ForeignKey
	DanglingKeys []DanglingKey
	Type         string
}

type Column struct {
	Name     string
	Type     string
	Nullable bool
}

type InclusionDependency struct {
	RelationAName string
	RelationBName string
	KeyA          string
	KeyB          string
}
