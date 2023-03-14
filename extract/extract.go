package extract

import (
	"database/sql"
	"fmt"
	"rdb-to-er-extractor/model"
)

func GetAllTables(db *sql.DB, dbName string) (tables []string) {
	rows, err := db.Query("SELECT TABLE_NAME from information_schema.TABLES t where TABLE_SCHEMA = ?", dbName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			fmt.Println(err.Error())
			return
		}

		tables = append(tables, table)
	}

	return
}

func GetPrimaryKeyFromRelation(db *sql.DB, dbName, relationName string) (primaryKeyColumns []string) {
	rows, err := db.Query(`
		SELECT kcu.COLUMN_NAME 
		FROM information_schema.TABLE_CONSTRAINTS tc 
		JOIN information_schema.KEY_COLUMN_USAGE kcu USING(constraint_name, table_schema, table_name) 
		WHERE TABLE_SCHEMA = ? AND tc.TABLE_NAME = ? AND tc.CONSTRAINT_TYPE = "PRIMARY KEY";`,
		dbName, relationName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			fmt.Println(err.Error())
			return
		}

		primaryKeyColumns = append(primaryKeyColumns, column)
	}

	return
}

func GetForeignKeyFromRelation(db *sql.DB, dbName, relationName string) (foreignKeyColumns []model.ForeignKey) {
	rows, err := db.Query(`
		SELECT kcu.COLUMN_NAME, kcu.REFERENCED_TABLE_NAME, kcu.REFERENCED_COLUMN_NAME 
		FROM information_schema.KEY_COLUMN_USAGE kcu 
		INNER JOIN information_schema.REFERENTIAL_CONSTRAINTS rc USING(constraint_name) 
		WHERE kcu.REFERENCED_TABLE_NAME IS NOT NULL AND kcu.TABLE_SCHEMA = ? AND kcu.TABLE_NAME=?;`,
		dbName, relationName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var row model.ForeignKey
		if err := rows.Scan(&row.ColumnName, &row.ReferencedTableName, &row.ReferencedColumnName); err != nil {
			fmt.Println(err.Error())
			return
		}

		foreignKeyColumns = append(foreignKeyColumns, row)
	}

	return
}
