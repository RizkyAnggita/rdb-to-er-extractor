package extract

import (
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func ClassifyStrongRelation(table *model.Table, allTable []model.Table) {
	// to check if a table is strong relation or not,
	// we need to check if its primary key appear as key of other relation, than it's not strong

	// If the number of primary key of a relation is one, definitely is a strong entity relation
	if len(table.PrimaryKeys) == 1 {
		table.Type = "STRONG"
		return
	}

	// otherwise, we need to check if every key of primary keys appear in other relation
	for _, t := range allTable {
		if t.Name != table.Name {
			if isAppear := isPKExistAsKeyInOtherTable(table.PrimaryKeys, t.PrimaryKeys); isAppear {
				return
			}
		}
	}

	table.Type = "STRONG"
	return
}

func ClassifyWeakRelation(table *model.Table, allTable []model.Table) {
	properSubset := helper.GenerateProperSubsetPK(table.PrimaryKeys)

	for _, subSet := range properSubset {
		// if a proper subset of its primary key, called K1, contains the keys of entity relations (strong
		// and/or weak).
		isContainOtherKey := isPKExistAsKeyInOtherEntityRelation(subSet, allTable)
		if isContainOtherKey {
			//The remaining attributes of the primary key, called K2, do not contain a key of any other relation
			remainingPK := setDifference(subSet, table.PrimaryKeys)

			isRemainingPKExistOther := isPKExistAsKeyInOtherEntityRelation(remainingPK, allTable)
			if !isRemainingPKExistOther {
				for _, key := range remainingPK {
					table.DanglingKeys = append(table.DanglingKeys, model.DanglingKey{ColumnName: key.ColumnName})
				}
				table.Type = "WEAK"
				return
			}
		}
	}
	return
}

func ClassifyRegularRelationshipRelation(table *model.Table, allTable []model.Table) {
	if len(table.PrimaryKeys) < 2 {
		return
	}

	count := 0

	for _, t := range allTable {
		if t.Name != table.Name && (t.Type == "STRONG" || t.Type == "WEAK") {
			if isAppear := isPKExistAsKeyInOtherTable(table.PrimaryKeys, t.PrimaryKeys); isAppear {
				count += 1
			}
		}
	}

	if count == len(table.PrimaryKeys) {
		table.Type = "REGULAR"
		return
	}

	return
}

func setDifference(subset []model.PrimaryKey, set []model.PrimaryKey) []model.PrimaryKey {
	diff := []model.PrimaryKey{}

	for _, s := range set {
		if !helper.IsExistInPrimaryKeys(s.ColumnName, subset) {
			diff = append(diff, s)
		}
	}

	return diff
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

func isPKExistAsKeyInOtherEntityRelation(primaryKey []model.PrimaryKey, allTable []model.Table) bool {
	for _, pk := range primaryKey {
		currPk := pk.ColumnName

		for _, at := range allTable {
			if at.Type == "STRONG" || at.Type == "WEAK" {
				if helper.IsExistInPrimaryKeys(currPk, at.PrimaryKeys) {
					return true
				}
			}
		}
	}

	return false
}
