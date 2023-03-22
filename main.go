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
	dbName := "classicmodels"

	db, err := sql.Open(dbType, username+":"+password+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	tableNames := extract.GetAllTables(db, dbName)
	tables := []model.Table{}

	for _, tableName := range tableNames {
		table := model.Table{Name: tableName}
		table.PrimaryKeys = extract.GetPrimaryKeyFromRelation(db, dbName, tableName)
		table.ForeignKeys = extract.GetForeignKeyFromRelation(db, dbName, tableName)
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
			fmt.Println("MASUK: ", tables[i].Name)
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
	fmt.Println("SUPERTYPE: ")
	for _, r := range inclusionDependencies {
		r.Print()
	}

	inclusionDependencies2 := inclusion.HeuristicRelationshipByForeignKey(tables)
	fmt.Println("FOREIGN: ")
	for _, r := range inclusionDependencies2 {
		r.Print()
	}

	inclusionDependencies3 := inclusion.HeuristicRelationShipOwnerAndParticipatingEntity(tables)
	fmt.Println("REGULAR: ")
	for _, r := range inclusionDependencies3 {
		r.Print()
	}

	inclusionDependencies = append(inclusionDependencies, inclusionDependencies2...)
	inclusionDependencies = append(inclusionDependencies, inclusionDependencies3...)

	fmt.Println("Inclusion Dependencies Generated: ")
	for _, r := range inclusionDependencies {
		r.Print()
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

	// Remove Duplicate
	fmt.Println("Remove Duplicate: ")
	inclusionDependencies = inclusion.RemoveDuplicateInclDepend(inclusionDependencies)

	fmt.Println("Remove Redundancy:")
	inclusionDependencies = inclusion.RemoveRedundantInclDepend(db, inclusionDependencies)
	for _, r := range inclusionDependencies {
		r.Print()
	}

	fmt.Println("_____")

	strongEntities := identification.IdentifyStrongEntities(tables)
	weakEntities, dependentRelationship := identification.IdentifyWeakEntities(tables, inclusionDependencies)
	inclusionRelationship := identification.IdentifyInclusionRelationship(tables, inclusionDependencies)
	binaryRelationship := identification.IdentifyBinaryRelationship(tables, inclusionDependencies)
	binaryRelationship2 := identification.IdentifyRelationshipByRegularRelationshipRelation(tables, inclusionDependencies)

	for _, strong := range strongEntities {
		fmt.Println("S: ", strong)
	}

	for _, weak := range weakEntities {
		fmt.Println("W: ", weak)
	}

	for _, dr := range dependentRelationship {
		fmt.Println("DR: ", dr)
	}

	for _, ir := range inclusionRelationship {
		fmt.Println("IR: ", ir)
	}

	for _, br := range binaryRelationship {
		fmt.Println("BR: ", br)
	}

	for _, br2 := range binaryRelationship2 {
		fmt.Println("BR2: ", br2)
	}

	// fmt.Println("RES: ", res)

	// router.Run("localhost:8080")
}
