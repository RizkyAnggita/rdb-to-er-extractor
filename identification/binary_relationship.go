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
				// Check if keyA is purely Foreign Key (not also a primary key which break the rule)
				if helper.IsExistInForeignKeys(id.KeyA, relationB.ForeignKeys) && !helper.IsExistInPrimaryKeys(id.KeyA, relationB.PrimaryKeys) {
					// Check if foreign key from relationB is a primary key in relationA
					isKeyBKeyInA := helper.IsExistInPrimaryKeys(id.KeyA, relationA.PrimaryKeys)
					if isKeyBKeyInA {
						relationship := model.Relationship{
							Name:          id.KeyA + "_" + relationA.Name,
							Type:          "BINARY",
							Cardinality:   "1-N",
							EntityAName:   relationA.Name,
							EntityBName:   relationB.Name,
							Identificator: id.KeyA,
						}

						relationships = append(relationships, relationship)
					} else {
						// Kasus ketika nama column Foreign Key berbeda dengan nama kolom primary key
						for _, fk := range relationB.ForeignKeys {
							if fk.ColumnName == id.KeyA {
								if fk.ReferencedTableName == relationA.Name && helper.IsExistInPrimaryKeys(fk.ReferencedColumnName, relationA.PrimaryKeys) {
									relationship := model.Relationship{
										Name:          fk.ColumnName + "_" + relationA.Name,
										Type:          "BINARY",
										Cardinality:   "1-N",
										EntityAName:   relationA.Name,
										EntityBName:   relationB.Name,
										Identificator: fk.ColumnName,
									}

									relationships = append(relationships, relationship)
								}
							}
						}
					}
				}

			}
		}
	}

	return relationships
}
