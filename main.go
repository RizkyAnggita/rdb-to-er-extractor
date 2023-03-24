package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"rdb-to-er-extractor/extract"
	"rdb-to-er-extractor/helper"
	"rdb-to-er-extractor/identification"
	"rdb-to-er-extractor/inclusion"
	"rdb-to-er-extractor/model"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.POST("/albums", postAlbums)
	router.GET("/albums", convertRDBtoEERModel)

	// fmt.Println("RES: ", res)

	router.Run("localhost:8080")
}

func convertRDBtoEERModel(c *gin.Context) {
	username := "root"
	password := "regars2000"
	dbType := "mysql"
	dbName := "my_school"

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
	binaryRelationship2, associativeEntities := identification.IdentifyRelationshipByRegularRelationshipRelation(tables, inclusionDependencies)

	mapNameKey := map[string]int{}
	keyCounter := 1

	for _, strong := range strongEntities {
		fmt.Println("S: ", strong)
		mapNameKey[strong.Name] = keyCounter
		keyCounter += 1
	}

	for _, weak := range weakEntities {
		fmt.Println("W: ", weak)
		mapNameKey[weak.Name] = keyCounter
		keyCounter += 1
	}

	for _, asc := range associativeEntities {
		fmt.Println("ASC: ", asc)
		mapNameKey[asc.Name] = keyCounter
		keyCounter += 1
	}

	for _, dr := range dependentRelationship {
		fmt.Println("DR: ", dr)
		mapNameKey[dr.Name] = keyCounter
		keyCounter += 1
	}

	for _, ir := range inclusionRelationship {
		fmt.Println("IR: ", ir)
		mapNameKey[ir.Name] = keyCounter
		keyCounter += 1
	}

	for _, br := range binaryRelationship {
		fmt.Println("BR: ", br)
		mapNameKey[br.Name] = keyCounter
		keyCounter += 1
	}

	for _, br2 := range binaryRelationship2 {
		fmt.Println("BR2: ", br2)
		mapNameKey[br2.Name] = keyCounter
		keyCounter += 1
	}

	ERModel := model.ERModel{}
	nodesData := []model.Node{}
	linkData := []model.Link{}

	for _, e := range strongEntities {
		node := model.Node{
			Text:                 e.Name,
			Color:                "black",
			Figure:               "Rectangle",
			Width:                90,
			Height:               50,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[e.Name],
			Location:             "-431.0127868652344 -80.75775146484375",
		}
		nodesData = append(nodesData, node)
	}

	for _, w := range weakEntities {
		node := model.Node{
			Text:                 w.Name,
			Color:                "black",
			Figure:               "DoubleRectangle",
			Width:                90,
			Height:               50,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[w.Name],
			Location:             "-231.0127868652344 -60.75775146484375",
		}
		nodesData = append(nodesData, node)
	}

	for _, asc := range associativeEntities {
		node := model.Node{
			Text:                 asc.Name,
			Color:                "black",
			Figure:               "AssociativeRectangle",
			Width:                90,
			Height:               50,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[asc.Name],
			Location:             "-331.0127868652344 -60.75775146484375",
		}
		nodesData = append(nodesData, node)

		entityA := helper.GetTableByTableName(asc.EntityAName, tables)
		entityB := helper.GetTableByTableName(asc.EntityBName, tables)
		link1 := model.Link{
			From: mapNameKey[entityA.Name],
			To:   node.Key,
			Text: "",
		}

		link2 := model.Link{
			From: node.Key,
			To:   mapNameKey[entityB.Name],
			Text: "",
		}
		linkData = append(linkData, link1, link2)
	}

	for _, dr := range dependentRelationship {
		node := model.Node{
			Text:                 dr.Name,
			Color:                "black",
			Figure:               "DoubleDiamond",
			Width:                120,
			Height:               50,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[dr.Name],
			Location:             "-131.0127868652344 -40.75775146484375",
		}
		nodesData = append(nodesData, node)

		ownerEntity := helper.GetTableByTableName(dr.EntityAName, tables)
		weakEntity := helper.GetTableByTableName(dr.EntityBName, tables)
		link1 := model.Link{
			From:  mapNameKey[ownerEntity.Name],
			To:    node.Key,
			Text:  "",
			IsOne: true,
		}

		link2 := model.Link{
			From: node.Key,
			To:   mapNameKey[weakEntity.Name],
			Text: "",
		}
		linkData = append(linkData, link1, link2)
	}

	for _, br := range binaryRelationship {
		node := model.Node{
			Text:                 br.Name,
			Color:                "black",
			Figure:               "Diamond",
			Width:                130,
			Height:               70,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[br.Name],
			Location:             "-251.0127868652344 -20.75775146484375",
		}
		nodesData = append(nodesData, node)

		relationA := helper.GetTableByTableName(br.EntityAName, tables)
		relationB := helper.GetTableByTableName(br.EntityBName, tables)
		link1 := model.Link{
			From:  mapNameKey[relationA.Name],
			To:    node.Key,
			Text:  "",
			IsOne: true,
		}

		link2 := model.Link{
			From: node.Key,
			To:   mapNameKey[relationB.Name],
			Text: "",
		}
		linkData = append(linkData, link1, link2)
	}

	for _, br := range binaryRelationship2 {
		node := model.Node{
			Text:                 br.Name,
			Color:                "black",
			Figure:               "Diamond",
			Width:                130,
			Height:               70,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[br.Name],
			Location:             "-331.0127868652344 -50.75775146484375",
		}
		nodesData = append(nodesData, node)

		relationA := helper.GetTableByTableName(br.EntityAName, tables)
		relationB := helper.GetTableByTableName(br.EntityBName, tables)
		link1 := model.Link{
			From:  mapNameKey[relationA.Name],
			To:    node.Key,
			Text:  "",
			IsOne: false,
		}

		link2 := model.Link{
			From:  node.Key,
			To:    mapNameKey[relationB.Name],
			Text:  "",
			IsOne: false,
		}
		linkData = append(linkData, link1, link2)
	}

	for _, node := range nodesData {
		fmt.Println("N: ", node)
	}

	for _, link := range linkData {
		fmt.Println("LINK: ", link)
	}

	ERModel.Class = "GraphLinksModel"
	ERModel.NodeDataArray = nodesData
	ERModel.LinkDataArray = linkData

	c.JSON(http.StatusOK, ERModel)
}
