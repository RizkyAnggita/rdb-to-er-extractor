package identification

import "rdb-to-er-extractor/model"

func IdentifyRelationshipByRegularRelationshipRelation(arrTable []model.Table, inclDepend []model.InclusionDependency) []model.Relationship {
	relationships := []model.Relationship{}

	for _, t := range arrTable {
		if t.Type == "REGULAR" {
			tInclDepend := []model.InclusionDependency{}
			for _, id := range inclDepend {
				if id.RelationAName == t.Name {
					tInclDepend = append(tInclDepend, id)
				}
			}

			if len(tInclDepend) == 2 {
				relationship := model.Relationship{
					Name:        t.Name,
					Type:        "BINARY",
					Cardinality: "N-N",
					EntityAName: tInclDepend[0].RelationBName,
					EntityBName: tInclDepend[1].RelationBName,
				}

				relationships = append(relationships, relationship)
			}

		}
	}

	return relationships
}
