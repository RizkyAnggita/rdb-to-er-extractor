package identification

import (
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func IdentifyRelationshipByRegularRelationshipRelation(arrTable []model.Table,
	inclDepend []model.InclusionDependency) ([]model.Relationship, []model.AssociativeEntity) {
	relationships := []model.Relationship{}
	assocEntities := []model.AssociativeEntity{}

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
			} else if len(tInclDepend) == 3 {
				for _, r := range relationships {
					if r.Type == "BINARY" && r.Cardinality == "N-N" {
						sideAName := r.EntityAName
						sideBName := r.EntityBName

						setPossibilities := []string{tInclDepend[0].RelationBName, tInclDepend[1].RelationBName, tInclDepend[2].RelationBName}
						setToCheck := []string{sideAName, sideBName}

						if helper.IsSubset(setToCheck, setPossibilities) {
							// Create new Associative Entity with the name r.Name
							// create new relationship between new Associative Entity and r.Name
							// Remove the N-N binary relationship between sideAName & sideBName
							newAssocEntity := model.AssociativeEntity{
								Name:        r.Name,
								Type:        "ASSOCIATIVE",
								EntityAName: sideAName,
								EntityBName: sideBName,
							}
							assocEntities = append(assocEntities, newAssocEntity)

							newRelationship := model.Relationship{
								Name:        t.Name,
								Type:        "BINARY",
								Cardinality: "N-N",
								EntityAName: r.Name,
								EntityBName: helper.SetDifference(setPossibilities, setToCheck)[0],
							}
							relationships = RemoveRelationshipByName(r.Name, relationships)
							relationships = append(relationships, newRelationship)
						}
					}
				}
			}

		}
	}

	return relationships, assocEntities
}

func RemoveRelationshipByName(relationshipName string, relationships []model.Relationship) []model.Relationship {
	for i := 0; i < len(relationships); i++ {
		if relationships[i].Name == relationshipName {
			relationships[i] = relationships[len(relationships)-1]
			return relationships[:len(relationships)-1]
		}
	}
	return relationships
}
