package model

import "fmt"

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

type StrongEntity struct {
	Name string
	Type string
	Keys []PrimaryKey
}

type WeakEntity struct {
	Name            string
	Type            string
	Keys            []DanglingKey
	OwnerEntityName string
}

type Relationship struct {
	Name        string
	Type        string
	Cardinality string
	EntityAName string
	EntityBName string
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

func (id InclusionDependency) IsEqualTo(other InclusionDependency) bool {
	return id.KeyA == other.KeyA && id.KeyB == other.KeyB &&
		id.RelationAName == other.RelationAName && id.RelationBName == other.RelationBName
}

func (id InclusionDependency) Print() {
	fmt.Printf("ID: %s.%s << %s.%s\n", id.RelationAName, id.KeyA, id.RelationBName, id.KeyB)
}
