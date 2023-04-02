package identification

import "rdb-to-er-extractor/model"

func IdentifyEntitiesAndRelationship(tables []model.Table, inclusionDependencies []model.InclusionDependency) (entities model.Entities, relationships model.Relationships) {
	strongEntities := IdentifyStrongEntities(tables)
	weakEntities, dependentRelationship := IdentifyWeakEntities(tables, inclusionDependencies)
	inclusionRelationship := IdentifyInclusionRelationship(tables, inclusionDependencies)
	binaryRelationship := IdentifyBinaryRelationship(tables, inclusionDependencies)
	binaryRelationship2, associativeEntities := IdentifyRelationshipByRegularRelationshipRelation(tables, inclusionDependencies)

	entities.StrongEntities = strongEntities
	entities.WeakEntities = weakEntities
	entities.AssociativeEntities = associativeEntities

	relationships.DependentRelationships = dependentRelationship
	relationships.InclusionRelationships = inclusionRelationship
	relationships.BinaryRelationships = binaryRelationship
	relationships.BinaryRelationships = append(relationships.BinaryRelationships, binaryRelationship2...)

	return
}
