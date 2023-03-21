package identification

import (
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func IdentifyBinaryRelationship(allTable []model.Table, inclDepend []model.InclusionDependency) []model.Relationship {
	relationships := []model.Relationship{}
	for _, id := range inclDepend {
		relationB := helper.GetTableByTableName(id.RelationAName, allTable)
		relationA := helper.GetTableByTableName(id.RelationBName, allTable)

		if relationA.Type == "STRONG" || relationA.Type == "WEAK" {
			if relationB.Type == "STRONG" || relationB.Type == "WEAK" {
				isKeyBKeyInA := helper.IsExistInPrimaryKeys(id.KeyA, relationA.PrimaryKeys)
				if isKeyBKeyInA {

					relationship := model.Relationship{
						Name:        id.KeyA + "-" + relationA.Name + "-" + relationB.Name,
						Type:        "BINARY",
						Cardinality: "1-N",
						EntityAName: relationA.Name,
						EntityBName: relationB.Name,
					}

					relationships = append(relationships, relationship)
				} else {
					// Kasus ketika nama column Foreign Key berbeda dengan nama kolom primary key
					for _, fk := range relationB.ForeignKeys {
						if fk.ColumnName == id.KeyA {
							if fk.ReferencedTableName == relationA.Name && helper.IsExistInPrimaryKeys(fk.ReferencedColumnName, relationA.PrimaryKeys) {
								relationship := model.Relationship{
									Name:        fk.ColumnName + "-" + relationA.Name + "-" + relationB.Name,
									Type:        "BINARY",
									Cardinality: "1-N",
									EntityAName: relationA.Name,
									EntityBName: relationB.Name,
								}

								relationships = append(relationships, relationship)
							}
						}
					}
				}

			}
		}
	}

	return relationships
}
