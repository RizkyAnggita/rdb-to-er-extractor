package identification

import (
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func IdentifyStrongEntities(allTable []model.Table) []model.StrongEntity {
	res := []model.StrongEntity{}

	for _, t := range allTable {
		if t.Type == "STRONG" {
			temp := model.StrongEntity{Name: t.Name, Type: t.Type}
			temp.Keys = t.PrimaryKeys
			res = append(res, temp)
		}

	}

	return res
}

func IdentifyWeakEntities(allTable []model.Table, inclDepend []model.InclusionDependency) ([]model.WeakEntity, []model.Relationship) {
	resA := []model.WeakEntity{}
	resB := []model.Relationship{}

	for _, t := range allTable {
		if t.Type == "WEAK" {
			for _, all := range allTable {
				for _, pk := range t.PrimaryKeys {
					if helper.IsExistInPrimaryKeys(pk.ColumnName, all.PrimaryKeys) {
						ownerEntityName := all.Name
						X := pk.ColumnName
						for _, id := range inclDepend {
							if id.RelationAName == t.Name && id.RelationBName == ownerEntityName &&
								id.KeyA == X && id.KeyB == X {
								weakEntity := model.WeakEntity{
									Name:            t.Name,
									Type:            t.Type,
									Keys:            t.DanglingKeys,
									OwnerEntityName: ownerEntityName,
								}
								relationship := model.Relationship{
									Name:        "DEPENDENT",
									Type:        "BINARY",
									Cardinality: "1-N",
									EntityAName: ownerEntityName,
									EntityBName: t.Name,
								}

								resA = append(resA, weakEntity)
								resB = append(resB, relationship)

							}
						}

					}
				}
			}
		}
	}

	return resA, resB
}
