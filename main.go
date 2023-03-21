package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"rdb-to-er-extractor/extract"
	"rdb-to-er-extractor/identification"
	"rdb-to-er-extractor/inclusion"
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
		fmt.Println("KADIEU: ", table.ForeignKeys)
		tables = append(tables, table)
	}

	for i := 0; i < len(tables); i++ {
		if tables[i].Type == "" {
			extract.ClassifyStrongRelation(&tables[i], tables)
		}
	}

	for i := 0; i < len(tables); i++ {
		if tables[i].Type == "" {
			extract.ClassifyWeakRelation(&tables[i], tables)
		}
	}

	for i := 0; i < len(tables); i++ {
		if tables[i].Type == "" {
			extract.ClassifyRegularRelationshipRelation(&tables[i], tables)
		}
	}

	for i := 0; i < len(tables); i++ {
		fmt.Println("____________________________")
		fmt.Println("NAME: ", tables[i].Name)
		fmt.Println("TYPE: ", tables[i].Type)
		fmt.Println("PK: ", tables[i].PrimaryKeys)
		fmt.Println("FK: ", tables[i].ForeignKeys)
		fmt.Println("DK: ", tables[i].DanglingKeys)
		fmt.Println("____________________________")
		// if tables[i].Name == "employees" || tables[i].Name == "offices" {
		// tables[i].PrimaryKeys = []model.PrimaryKey{{ColumnName: "customerNumber"}}
		// }
	}

	inclusionDependencies := []model.InclusionDependency{}
	inclusionDependencies = append(inclusionDependencies, inclusion.HeuristicSupertypeRelationship(tables)...)
	fmt.Println("SUPERTYPE: ", inclusionDependencies)
	inclusionDependencies = append(inclusionDependencies, inclusion.HeuristicRelationshipByForeignKey(tables)...)
	inclusionDependencies = append(inclusionDependencies, inclusion.HeuristicRelationShipOwnerAndParticipatingEntity(tables)...)

	fmt.Println("Inclusion Dependencies Generated: ")
	for _, r := range inclusionDependencies {
		fmt.Println("H: ", r)
	}

	fmt.Println("Reject Invalid Inclusion Dependencies")
	k := 0
	for _, r := range inclusionDependencies {
		isRejected, err := inclusion.IsRejectInclusionDependency(db, r)
		if err != nil {
			panic(err.Error())
		}

		if !isRejected {
			inclusionDependencies[k] = r
			k++
		} else {
			fmt.Println("THIS IS REJECTED: ", r)
		}
	}
	inclusionDependencies = inclusionDependencies[:k]
	for _, r := range inclusionDependencies {
		fmt.Println("H: ", r)
	}

	// Remove Duplicate
	fmt.Println("Remove Duplicate: ")
	inclusionDependencies = inclusion.RemoveDuplicateInclDepend(inclusionDependencies)
	for _, r := range inclusionDependencies {
		fmt.Println("H: ", r)
	}

	fmt.Println("Remove Redundancy:")
	inclusionDependencies = inclusion.RemoveRedundantInclDepend(db, inclusionDependencies)
	for _, r := range inclusionDependencies {
		fmt.Println("H: ", r)
	}

	fmt.Println("_____")

	strongEntities := identification.IdentifyStrongEntities(tables)
	weakEntities, relationship := identification.IdentifyWeakEntities(tables, inclusionDependencies)
	relationship = append(relationship, identification.IdentifyInclusionRelationship(tables, inclusionDependencies)...)

	for _, strong := range strongEntities {
		fmt.Println("S: ", strong)
	}

	for _, weak := range weakEntities {
		fmt.Println("W: ", weak)
	}

	for _, r := range relationship {
		fmt.Println("R: ", r)
	}

	// fmt.Println("RES: ", res)

	// router.Run("localhost:8080")
}
