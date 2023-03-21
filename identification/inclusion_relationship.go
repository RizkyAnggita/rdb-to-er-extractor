package identification

import (
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func IdentifyInclusionRelationship(allTable []model.Table, inclDepend []model.InclusionDependency) []model.Relationship {
	// Negative test ok, need case when there is an actual Inclusion Relationship (is-a) exist in the database
	relationships := []model.Relationship{}
	for _, id := range inclDepend {
		relationA := helper.GetTableByTableName(id.RelationAName, allTable)
		relationB := helper.GetTableByTableName(id.RelationBName, allTable)

		if relationA.Type == "STRONG" && relationB.Type == "STRONG" {
			isKeyAPK := helper.IsExistInPrimaryKeys(id.KeyA, relationA.PrimaryKeys)
			isKeyBPK := helper.IsExistInPrimaryKeys(id.KeyB, relationB.PrimaryKeys)

			if isKeyAPK && isKeyBPK {
				relationship := model.Relationship{
					Name:        relationA.Name + "-" + relationB.Name,
					Type:        "SPECIALIZATION",
					Cardinality: "1-1",
					EntityAName: relationA.Name,
					EntityBName: relationB.Name,
				}

				relationships = append(relationships, relationship)
			}
		}

	}
	return relationships
}
