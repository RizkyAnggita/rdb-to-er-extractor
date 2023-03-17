package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"rdb-to-er-extractor/extract"
	"rdb-to-er-extractor/model"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	username := "root"
	password := "regars2000"
	dbType := "mysql"

	db, err := sql.Open(dbType, username+":"+password+"@tcp(localhost:3306)/classicmodels")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	tableNames := extract.GetAllTables(db, "classicmodels")
	tables := []model.Table{}

	for _, tableName := range tableNames {
		table := model.Table{Name: tableName}
		table.PrimaryKeys = extract.GetPrimaryKeyFromRelation(db, "classicmodels", tableName)
		table.ForeignKeys = extract.GetForeignKeyFromRelation(db, "classicmodels", tableName)
		tables = append(tables, table)
	}

	for i := 0; i < len(tables); i++ {
		// fmt.Println("NAME: ", tables[i].Name)
		if isStrong := extract.IsStrongRelation(tables[i], tables); isStrong {
			tables[i].Type = "STRONG"
		}

		// fmt.Println("STRONG: ", tables[i].Type)
		// fmt.Println("------")
	}

	for i := 0; i < len(tables); i++ {
		// fmt.Println("NAME: ", tables[i].Name)
		if isWeak := extract.IsWeakRelation(tables[i], tables); isWeak {
			tables[i].Type = "WEAK"
		}

		fmt.Println("WEAK: ", tables[i].Type)
		fmt.Println("------")
	}

	for i := 0; i < len(tables); i++ {
		fmt.Println("NAME: ", tables[i].Name)
		if tables[i].Type == "" {
			if isRegular := extract.IsRegularRelationshipRelation(tables[i], tables); isRegular {
				tables[i].Type = "REGULAR"
			}
		}

		fmt.Println("TYPE: ", tables[i].Type)
		fmt.Println("------")
	}

	// fmt.Println(helper.GenerateProperSubsetPK([]model.PrimaryKey{{ColumnName: "A"}, {ColumnName: "B"}, {ColumnName: "C"}}))

	// for _, table := range tableNames {
	// 	primaryKeyColumns := extract.GetPrimaryKeyFromRelation(db, "classicmodels", table)
	// 	foreignKeyColumns := extract.GetForeignKeyFromRelation(db, "classicmodels", table)

	// 	fmt.Println("\nTABLE: ", table)
	// 	fmt.Println("PK: ", primaryKeyColumns)

	// 	for _, fk := range foreignKeyColumns {
	// 		fmt.Printf("FK: %+v\n", fk)
	// 	}

	// 	fmt.Println("--------------------------")
	// }

	// router.Run("localhost:8080")
}
