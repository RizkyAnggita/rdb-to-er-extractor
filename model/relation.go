package model

import "fmt"

type Database struct {
	Tables []Table
}

type ForeignKey struct {
	ColumnName           string
	Type                 string
	ReferencedTableName  string
	ReferencedColumnName string
}

type PrimaryKey struct {
	ColumnName string
	Type       string
}

type DanglingKey struct {
	ColumnName string
	Type       string
}

type Table struct {
	Name         string
	Columns      []Column
	PrimaryKeys  []PrimaryKey
	ForeignKeys  []ForeignKey
	DanglingKeys []DanglingKey
	Type         string
}

type Entities struct {
	StrongEntities      []StrongEntity
	WeakEntities        []WeakEntity
	AssociativeEntities []AssociativeEntity
}

type Relationships struct {
	BinaryRelationships    []Relationship
	InclusionRelationships []Relationship
	DependentRelationships []Relationship
}

type StrongEntity struct {
	Name    string
	Type    string
	Keys    []PrimaryKey
	Columns []Column
}

type WeakEntity struct {
	Name            string
	Type            string
	Keys            []DanglingKey
	OwnerEntityName string
	Columns         []Column
}

type AssociativeEntity struct {
	Name        string
	Type        string
	Keys        []PrimaryKey
	EntityAName string
	EntityBName string
	Columns     []Column
}

type Relationship struct {
	Name          string
	Type          string
	Cardinality   string
	EntityAName   string
	EntityBName   string
	Columns       []Column
	Identificator string
}

type Column struct {
	Name          string
	Type          string
	IsNullable    string
	IsMultivalues bool
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
