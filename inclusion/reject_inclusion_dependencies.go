package inclusion

import (
	"database/sql"
	"fmt"
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/model"
)

func IsRejectInclusionDependency(db *sql.DB, inclDepend model.InclusionDependency) (bool, error) {

	queryA := "SELECT DISTINCT(" + inclDepend.KeyA + ") FROM " + inclDepend.RelationAName
	queryB := "SELECT DISTINCT(" + inclDepend.KeyB + ") FROM " + inclDepend.RelationBName
	rowsA, err := db.Query(queryA)
	if err != nil {
		fmt.Println(err.Error())
		return true, err
	}

	rowsB, err := db.Query(queryB)
	if err != nil {
		fmt.Println(err.Error())
		return true, err
	}

	setValueAX := []string{}
	setValueBX := []string{}

	for rowsA.Next() {
		var valueA sql.NullString
		if err := rowsA.Scan(&valueA); err != nil {
			fmt.Println(err.Error())
			return true, err
		}
		if valueA.Valid {
			setValueAX = append(setValueAX, valueA.String)
		}

	}

	for rowsB.Next() {
		var valueB sql.NullString
		if err := rowsB.Scan(&valueB); err != nil {
			fmt.Println(err.Error())
			return true, err
		}
		if valueB.Valid {
			setValueBX = append(setValueBX, valueB.String)
		}
	}

	if len(setValueAX) != 0 && helper.IsSubset(setValueAX, setValueBX) {
		return false, nil
	}

	return true, nil
}
