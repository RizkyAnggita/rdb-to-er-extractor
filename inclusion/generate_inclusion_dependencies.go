package inclusion

import (
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func HeuristicSupertypeRelationship(arrTable []model.Table) (res []model.InclusionDependency) {
	// Heuristic 5.2.1 point 1
	for i := 0; i < len(arrTable); i++ {
		pks := arrTable[i].PrimaryKeys
		if arrTable[i].Type == "STRONG" {
			for j := i + 1; j < len(arrTable); j++ {
				if arrTable[j].Type == "STRONG" {
					for _, pk := range pks {
						if isExistInPK := helper.IsExistInPrimaryKeys(pk.ColumnName, arrTable[j].PrimaryKeys); isExistInPK {
							inclusionDependency := model.InclusionDependency{
								RelationAName: arrTable[i].Name,
								RelationBName: arrTable[j].Name,
								KeyA:          pk.ColumnName,
								KeyB:          pk.ColumnName,
							}
							res = append(res, inclusionDependency)
						}
					}
				}
			}
		}
	}

	return res
}

func HeuristicRelationshipByForeignKey(arrTable []model.Table) (res []model.InclusionDependency) {
	// Heuristic 5.2.1 point 2
	lenTable := len(arrTable)
	for i := 0; i < lenTable; i++ {
		relationType := arrTable[i].Type
		if relationType == "STRONG" || relationType == "WEAK" {
			pks := arrTable[i].PrimaryKeys
			for j := 0; j < lenTable; j++ {
				for _, pk := range pks {
					if helper.IsExistInForeignKeys(pk.ColumnName, arrTable[j].ForeignKeys) && arrTable[i].Name != arrTable[j].Name {
						if !helper.IsExistInPrimaryKeys(pk.ColumnName, arrTable[j].PrimaryKeys) {
							res = append(res, model.InclusionDependency{
								RelationAName: arrTable[j].Name,
								RelationBName: arrTable[i].Name,
								KeyA:          pk.ColumnName,
								KeyB:          pk.ColumnName,
							})
						}
					} else {
						for _, fk := range arrTable[j].ForeignKeys {
							if pk.ColumnName == fk.ReferencedColumnName && pk.ColumnName != fk.ColumnName {
								res = append(res, model.InclusionDependency{
									RelationAName: arrTable[j].Name,
									RelationBName: arrTable[i].Name,
									KeyA:          fk.ColumnName,
									KeyB:          pk.ColumnName,
								})
							}
						}
					}
				}
			}
		}
	}

	return
}

func HeuristicRelationShipOwnerAndParticipatingEntity(arrTable []model.Table) (res []model.InclusionDependency) {
	// Heuristic 5.2.1 point 1
	lenTable := len(arrTable)

	for i := 0; i < lenTable; i++ {
		if arrTable[i].Type != "STRONG" {
			pks := arrTable[i].PrimaryKeys
			for j := 0; j < lenTable; j++ {
				jType := arrTable[j].Type
				if arrTable[i].Name != arrTable[j].Name && (jType == "STRONG" || jType == "WEAK") {
					for _, pk := range pks {
						if helper.IsExistInPrimaryKeys(pk.ColumnName, arrTable[j].PrimaryKeys) {
							res = append(res, model.InclusionDependency{
								RelationAName: arrTable[i].Name,
								RelationBName: arrTable[j].Name,
								KeyA:          pk.ColumnName,
								KeyB:          pk.ColumnName,
							})
						}
					}
				}
			}
		}
	}

	return
}
