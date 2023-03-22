package inclusion

import (
	"database/sql"
	"fmt"
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func RemoveRedundantInclDepend(db *sql.DB, inclDependencies []model.InclusionDependency) []model.InclusionDependency {

	for i := 0; i < len(inclDependencies); i++ {
		for j := 0; j < len(inclDependencies) && i >= 0; j++ {
			b := inclDependencies[i].RelationBName
			bx := inclDependencies[i].KeyB
			b2 := inclDependencies[j].RelationAName
			b2y := inclDependencies[j].KeyA

			if b == b2 && i != j {
				queryX := "SELECT DISTINCT(" + bx + ") FROM " + b
				queryY := "SELECT DISTINCT(" + b2y + ") FROM " + b2

				rowsA, err := db.Query(queryX)
				if err != nil {
					fmt.Println(err.Error())
					return inclDependencies
				}

				rowsB, err := db.Query(queryY)
				if err != nil {
					fmt.Println(err.Error())
					return inclDependencies
				}

				setValueX := []string{}
				setValueY := []string{}

				for rowsA.Next() {
					var valueX sql.NullString
					if err := rowsA.Scan(&valueX); err != nil {
						fmt.Println(err.Error())
						return inclDependencies
					}
					if valueX.Valid {
						setValueX = append(setValueX, valueX.String)
					}

				}

				for rowsB.Next() {
					var valueY sql.NullString
					if err := rowsB.Scan(&valueY); err != nil {
						fmt.Println(err.Error())
						return inclDependencies
					}
					if valueY.Valid {
						setValueY = append(setValueY, valueY.String)
					}
				}

				if helper.IsSubset(setValueY, setValueX) {
					// remove A.Y << C.Y
					a := inclDependencies[i].RelationAName
					ay := inclDependencies[j].KeyA
					c := inclDependencies[j].RelationBName
					cy := inclDependencies[j].KeyB

					for k := 0; k < len(inclDependencies); k++ {
						if inclDependencies[k].RelationAName == a && inclDependencies[k].KeyA == ay &&
							inclDependencies[k].RelationBName == c && inclDependencies[k].KeyB == cy {
							inclDependencies = append(inclDependencies[:k], inclDependencies[k+1:]...)
							i = i - 1
							j = 0
						}
					}
				}

			}
		}
	}

	return inclDependencies
}

func RemoveDuplicateInclDepend(inclDependencies []model.InclusionDependency) []model.InclusionDependency {
	for i := 0; i < len(inclDependencies); i++ {
		for j := i + 1; j < len(inclDependencies); j++ {
			if i != j && inclDependencies[i].IsEqualTo(inclDependencies[j]) {
				inclDependencies = append(inclDependencies[:j], inclDependencies[j+1:]...)
				i = i - 1
				j = 0
			}
		}
	}

	return inclDependencies
}
