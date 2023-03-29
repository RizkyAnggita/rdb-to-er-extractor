package extract

import (
	"database/sql"
	"fmt"
	"rdb-to-er-extractor/model"
)

func GetAllTables(db *sql.DB, driver, dbName string) (tables []string) {
	query := ""
	param1 := ""
	if driver == "mysql" {
		query = `SELECT TABLE_NAME from information_schema.TABLES t where TABLE_SCHEMA = ?;`
		param1 = dbName
	} else if driver == "postgres" {
		query = `SELECT TABLE_NAME from information_schema.TABLES t WHERE table_schema = $1 AND table_type='BASE TABLE';`
		param1 = `public`
	}

	rows, err := db.Query(query, param1)
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

func GetPrimaryKeyFromRelation(db *sql.DB, dbName, driver, relationName string) (primaryKeyColumns []model.PrimaryKey) {
	whereQuery := ""
	if driver == "mysql" {
		whereQuery = ` WHERE TABLE_SCHEMA = ? AND tc.TABLE_NAME = ? AND tc.CONSTRAINT_TYPE = "PRIMARY KEY";`
	} else if driver == "postgres" {
		whereQuery = ` WHERE TABLE_SCHEMA = $1 AND tc.TABLE_NAME = $2 AND tc.CONSTRAINT_TYPE = 'PRIMARY KEY';`
		dbName = "public"
	}

	rows, err := db.Query(`
		SELECT kcu.COLUMN_NAME 
		FROM information_schema.TABLE_CONSTRAINTS tc 
		JOIN information_schema.KEY_COLUMN_USAGE kcu USING(constraint_name, table_schema, table_name)
		`+whereQuery,
		dbName, relationName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var row model.PrimaryKey
		if err := rows.Scan(&row.ColumnName); err != nil {
			fmt.Println(err.Error())
			return
		}

		primaryKeyColumns = append(primaryKeyColumns, row)
	}

	return
}

func GetForeignKeyFromRelation(db *sql.DB, dbName, driver, relationName string) (foreignKeyColumns []model.ForeignKey) {
	query := ""
	param1 := ""
	param2 := ""

	if driver == "mysql" {
		query = ` 
		SELECT 
			kcu.COLUMN_NAME, kcu.REFERENCED_TABLE_NAME, kcu.REFERENCED_COLUMN_NAME 
		FROM 
			information_schema.KEY_COLUMN_USAGE kcu 
		JOIN 
			information_schema.REFERENTIAL_CONSTRAINTS rc USING(constraint_name)
		WHERE kcu.REFERENCED_TABLE_NAME IS NOT NULL AND kcu.TABLE_SCHEMA = ? AND kcu.TABLE_NAME=?`
		param1 = dbName
		param2 = relationName

	} else if driver == "postgres" {
		query = `
			SELECT
				kcu.column_name, ccu.table_name AS REFERENCED_TABLE_NAME, ccu.column_name AS REFERENCED_COLUMN_NAME 
			FROM 
				information_schema.table_constraints AS tc 
				JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema
				JOIN information_schema.constraint_column_usage AS ccu ON ccu.constraint_name = tc.constraint_name AND ccu.table_schema = tc.table_schema
			WHERE tc.constraint_type = $1 AND tc.table_name= $2;
		`
		param1 = `FOREIGN KEY`
		param2 = relationName
	}

	rows, err := db.Query(query, param1, param2)

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
