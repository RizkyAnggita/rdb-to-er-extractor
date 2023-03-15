package extract

import (
	"rdb-to-er-extractor/model"
)

func IsStrongRelation(table model.Table, allTable []model.Table) bool {
	// to check if a table is strong relation or not,
	// we need to check if its primary key appear as key of other relation, than it's not strong

	// If the number of primary key of a relation is one, definitely is a strong entity relation
	if len(table.PrimaryKeys) == 1 {
		return true
	}

	// otherwise, we need to check if every key of primary keys appear in other relation
	for _, t := range allTable {
		if t.Name != table.Name {
			if isAppear := isPKExistAsKeyInOtherTable(table.PrimaryKeys, t.PrimaryKeys); isAppear {
				return false
			}
		}
	}

	return true
}

func IsRegularRelationshipRelation(table model.Table, allTable []model.Table) bool {
	if len(table.PrimaryKeys) < 2 {
		return false
	}

	count := 0

	for _, t := range allTable {
		if t.Name != table.Name && (t.Type == "STRONG" || t.Type == "WEAK") {
			if isAppear := isPKExistAsKeyInOtherTable(table.PrimaryKeys, t.PrimaryKeys); isAppear {
				count += 1
			}
		}
	}

	return count == len(table.PrimaryKeys)
}

func isPKExistAsKeyInOtherTable(primaryKey []model.PrimaryKey, otherPrimaryKey []model.PrimaryKey) bool {
	for _, pk := range primaryKey {
		currPk := pk.ColumnName

		for _, opk := range otherPrimaryKey {

			if currPk == opk.ColumnName {
				return true
			}
		}
	}

	return false
}
