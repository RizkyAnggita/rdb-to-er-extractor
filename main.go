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
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type ExtractERParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
	Port     string `json:"port"`
	URL      string `json:"url"`
	Driver   string `json:"driver"`
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.POST("/extract-eer", convertRDBtoEERModel)

	// fmt.Println("RES: ", res)

	router.Run("localhost:8080")
}

func convertRDBtoEERModel(c *gin.Context) {
	var params ExtractERParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := params.Username
	password := params.Password
	url := params.URL
	port := params.Port
	driverParams := params.Driver
	dbName := params.DBName

	driver := ""
	dataSourceName := ""
	if driverParams == "MySQL" {
		driver = "mysql"
		dataSourceName = username + ":" + password + "@tcp(" + url + ":" + port + ")/" + dbName
	} else if driverParams == "PostgreSQL" {
		driver = "postgres"
		dataSourceName = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, url, port, dbName)
	}

	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	tables := GenerateRelationsFromTables(db, driver, dbName)
	tables = GetNonKeyColumnsFromTables(db, driver, dbName, tables)

	fmt.Println("Inclusion Dependencies Generated: ")
	inclusionDependencies := GenerateInclusionDependencies(db, tables)

	fmt.Println("_____")

	entities, relationships := identification.IdentifyEntitiesAndRelationship(tables, inclusionDependencies)

	mapNameKey := map[string]int{}
	keyCounter := 1

	for _, strong := range entities.StrongEntities {
		fmt.Println("S: ", strong)
		mapNameKey[strong.Name] = keyCounter
		keyCounter += 1
		for _, pk := range strong.Keys {
			mapNameKey[strong.Name+pk.ColumnName] = keyCounter
			keyCounter += 1
		}

		for _, col := range strong.Columns {
			mapNameKey[strong.Name+col.Name] = keyCounter
			keyCounter += 1
		}
	}

	for _, weak := range entities.WeakEntities {
		fmt.Println("W: ", weak)
		mapNameKey[weak.Name] = keyCounter
		keyCounter += 1
		for _, dk := range weak.Keys {
			mapNameKey[weak.Name+dk.ColumnName] = keyCounter
			keyCounter += 1
		}
		for _, col := range weak.Columns {
			mapNameKey[weak.Name+col.Name] = keyCounter
			keyCounter += 1
		}
	}

	for _, asc := range entities.AssociativeEntities {
		fmt.Println("ASC: ", asc)
		mapNameKey[asc.Name] = keyCounter
		keyCounter += 1
		for _, pk := range asc.Keys {
			mapNameKey[asc.Name+pk.ColumnName] = keyCounter
			keyCounter += 1
		}
		for _, col := range asc.Columns {
			mapNameKey[asc.Name+col.Name] = keyCounter
			keyCounter += 1
		}
	}

	for _, dr := range relationships.DependentRelationships {
		fmt.Println("DR: ", dr)
		mapNameKey[dr.Name] = keyCounter
		keyCounter += 1
	}

	for _, ir := range relationships.InclusionRelationships {
		fmt.Println("IR: ", ir)
		mapNameKey[ir.Name] = keyCounter
		keyCounter += 1
	}

	for _, br := range relationships.BinaryRelationships {
		fmt.Println("BR: ", br)
		mapNameKey[br.Name] = keyCounter
		keyCounter += 1
		for _, col := range br.Columns {
			mapNameKey[br.Name+col.Name] = keyCounter
			keyCounter += 1
		}
	}

	ERModel := model.ERModel{}
	nodesData := []model.Node{}
	linkData := []model.Link{}

	for _, e := range entities.StrongEntities {
		entity := model.Node{
			Text:                 e.Name,
			Color:                "black",
			Figure:               "Rectangle",
			Width:                90,
			Height:               50,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[e.Name],
			Location:             "-100 -100.75775146484375",
		}
		nodesData = append(nodesData, entity)

		for _, pk := range e.Keys {
			pkAttrib := model.Node{
				Text:         pk.ColumnName,
				Color:        "black",
				Figure:       "Ellipse",
				FromMaxLinks: 1,
				Height:       30,
				Width:        10,
				Key:          mapNameKey[e.Name+pk.ColumnName],
				Location:     "-150.0127868652344 -100.75775146484375",
				Underline:    true,
			}

			link := model.Link{
				From: pkAttrib.Key,
				To:   entity.Key,
				Text: "",
			}
			linkData = append(linkData, link)

			nodesData = append(nodesData, pkAttrib)
		}

		for _, col := range e.Columns {
			colAttrib := model.Node{
				Text:         col.Name,
				Color:        "black",
				Figure:       "Ellipse",
				FromMaxLinks: 1,
				Height:       30,
				Width:        10,
				Key:          mapNameKey[e.Name+col.Name],
				Location:     "-150.0127868652344 -100.75775146484375",
			}

			link := model.Link{
				From: colAttrib.Key,
				To:   entity.Key,
				Text: "",
			}
			linkData = append(linkData, link)

			nodesData = append(nodesData, colAttrib)
		}

	}

	for _, w := range entities.WeakEntities {
		weakEntity := model.Node{
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
		nodesData = append(nodesData, weakEntity)

		for _, dk := range w.Keys {
			dkAttrib := model.Node{
				Text:         dk.ColumnName,
				Color:        "black",
				Figure:       "Ellipse",
				FromMaxLinks: 1,
				Height:       30,
				Width:        10,
				Key:          mapNameKey[w.Name+dk.ColumnName],
				Location:     "-431.0127868652344 -80.75775146484375",
				Underline:    true,
			}

			link := model.Link{
				From: dkAttrib.Key,
				To:   weakEntity.Key,
				Text: "",
			}
			nodesData = append(nodesData, dkAttrib)
			linkData = append(linkData, link)
		}

		for _, col := range w.Columns {
			colAttrib := model.Node{
				Text:         col.Name,
				Color:        "black",
				Figure:       "Ellipse",
				FromMaxLinks: 1,
				Height:       30,
				Width:        10,
				Key:          mapNameKey[w.Name+col.Name],
				Location:     "-150.0127868652344 -100.75775146484375",
			}

			link := model.Link{
				From: colAttrib.Key,
				To:   weakEntity.Key,
				Text: "",
			}
			linkData = append(linkData, link)

			nodesData = append(nodesData, colAttrib)
		}

	}

	for _, asc := range entities.AssociativeEntities {
		ascEntity := model.Node{
			Text:                 asc.Name,
			Color:                "black",
			Figure:               "AssociativeRectangle",
			Width:                90,
			Height:               50,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			Key:                  mapNameKey[asc.Name],
			Location:             "-200.0127868652344 -60.75775146484375",
		}
		nodesData = append(nodesData, ascEntity)

		entityA := helper.GetTableByTableName(asc.EntityAName, tables)
		entityB := helper.GetTableByTableName(asc.EntityBName, tables)
		link1 := model.Link{
			From: ascEntity.Key,
			To:   mapNameKey[entityA.Name],
			Text: "",
		}

		link2 := model.Link{
			From: ascEntity.Key,
			To:   mapNameKey[entityB.Name],
			Text: "",
		}

		linkData = append(linkData, link1, link2)
		for _, pk := range asc.Keys {
			pkAttrib := model.Node{
				Text:         pk.ColumnName,
				Color:        "black",
				Figure:       "Ellipse",
				FromMaxLinks: 1,
				Height:       30,
				Width:        10,
				Key:          mapNameKey[asc.Name+pk.ColumnName],
				Location:     "-431.0127868652344 -80.75775146484375",
				Underline:    true,
			}

			link := model.Link{
				From: pkAttrib.Key,
				To:   ascEntity.Key,
				Text: "",
			}
			nodesData = append(nodesData, pkAttrib)
			linkData = append(linkData, link)
		}
	}

	for _, dr := range relationships.DependentRelationships {
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
			From:  node.Key,
			To:    mapNameKey[ownerEntity.Name],
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

	for _, br := range relationships.BinaryRelationships {
		node := model.Node{
			Text:                 br.Name,
			Color:                "black",
			Figure:               "Diamond",
			Width:                130,
			Height:               70,
			FromLinkable:         false,
			ToLinkableDuplicates: true,
			FromMaxLinks:         2,
			Key:                  mapNameKey[br.Name],
			Location:             "-251.0127868652344 -20.75775146484375",
		}
		nodesData = append(nodesData, node)

		relationA := helper.GetTableByTableName(br.EntityAName, tables)
		relationB := helper.GetTableByTableName(br.EntityBName, tables)
		link1 := model.Link{
			From:  node.Key,
			To:    mapNameKey[relationA.Name],
			Text:  "",
			IsOne: br.Cardinality != "N-N",
		}

		link2 := model.Link{
			From: node.Key,
			To:   mapNameKey[relationB.Name],
			Text: "",
		}
		linkData = append(linkData, link1, link2)

		for _, col := range br.Columns {
			colAttrib := model.Node{
				Text:         col.Name,
				Color:        "black",
				Figure:       "Ellipse",
				FromMaxLinks: 1,
				Height:       30,
				Width:        10,
				Key:          mapNameKey[br.Name+col.Name],
				Location:     "-150.0127868652344 -100.75775146484375",
			}

			link := model.Link{
				From: colAttrib.Key,
				To:   node.Key,
				Text: "",
			}
			linkData = append(linkData, link)

			nodesData = append(nodesData, colAttrib)
		}

	}

	parentsInclusion := map[string][]string{}

	for _, ir := range relationships.InclusionRelationships {
		if _, isExist := parentsInclusion[ir.EntityBName]; isExist {
			parentsInclusion[ir.EntityBName] = append(parentsInclusion[ir.EntityBName], ir.EntityAName)
		} else {
			parentsInclusion[ir.EntityBName] = []string{ir.EntityAName}
		}
	}

	for k, v := range parentsInclusion {
		mapNameKey["Specialization"+k] = keyCounter
		keyCounter += 1
		node := model.Node{
			Text:       "Specialization",
			Color:      "black",
			Figure:     "TriangleDown",
			Width:      130,
			Height:     70,
			ToLinkable: false,
			Key:        mapNameKey["Specialization"+k],
			Location:   "-251.0127868652344 -20.75775146484375",
		}
		nodesData = append(nodesData, node)

		// connect new node, "specialization" to parent
		parentKey := mapNameKey[k]
		link := model.Link{
			From:     node.Key,
			To:       parentKey,
			Text:     "",
			IsParent: true,
		}
		linkData = append(linkData, link)

		for _, child := range v {
			childKey := mapNameKey[child]
			link := model.Link{
				From: node.Key,
				To:   childKey,
				Text: "",
			}
			linkData = append(linkData, link)
			// remove key from child
			nodesToRemove := []int{}

			for i := 0; i < len(linkData); i++ {
				if linkData[i].To == childKey {
					nodesToRemove = append(nodesToRemove, linkData[i].From)
				}
			}

			for _, node := range nodesToRemove {
				for i := 0; i < len(nodesData); i++ {
					if nodesData[i].Key == node && nodesData[i].Underline {
						for j := 0; j < len(linkData); j++ {
							if linkData[j].From == nodesData[i].Key && linkData[j].To == childKey {
								linkData = append(linkData[:j], linkData[j+1:]...)
								break
							}
						}
						nodesData = append(nodesData[:i], nodesData[i+1:]...)
						break
					}
				}
			}

		}
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

func GenerateRelationsFromTables(db *sql.DB, driver, dbName string) []model.Table {
	tableNames := extract.GetAllTables(db, driver, dbName)
	tables := []model.Table{}

	for _, tableName := range tableNames {
		table := model.Table{Name: tableName}
		table.PrimaryKeys = extract.GetPrimaryKeyFromRelation(db, dbName, driver, tableName)
		table.ForeignKeys = extract.GetForeignKeyFromRelation(db, dbName, driver, tableName)
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
	}

	return tables
}

func GenerateInclusionDependencies(db *sql.DB, tables []model.Table) []model.InclusionDependency {
	ids := []model.InclusionDependency{}

	ids = append(ids, inclusion.HeuristicSupertypeRelationship(tables)...)
	fmt.Println("SUPERTYPE: ")
	for _, r := range ids {
		r.Print()
	}

	ids2 := inclusion.HeuristicRelationshipByForeignKey(tables)
	fmt.Println("FOREIGN: ")
	for _, r := range ids2 {
		r.Print()
	}

	ids3 := inclusion.HeuristicRelationShipOwnerAndParticipatingEntity(tables)
	fmt.Println("REGULAR: ")
	for _, r := range ids3 {
		r.Print()
	}

	ids = append(ids, ids2...)
	ids = append(ids, ids3...)

	for _, r := range ids {
		r.Print()
	}

	fmt.Println("Reject Invalid Inclusion Dependencies")
	k := 0
	for _, r := range ids {
		isRejected, err := inclusion.IsRejectInclusionDependency(db, r)
		if err != nil {
			panic(err.Error())
		}

		if !isRejected {
			ids[k] = r
			k++
		} else {
			fmt.Println("THIS IS REJECTED: ", r)
		}
	}
	ids = ids[:k]

	// Remove Duplicate
	fmt.Println("Remove Duplicate: ")
	ids = inclusion.RemoveDuplicateInclDepend(ids)

	fmt.Println("Remove Redundancy:")
	ids = inclusion.RemoveRedundantInclDepend(db, ids)
	for _, r := range ids {
		r.Print()
	}
	return ids
}

func GetNonKeyColumnsFromTables(db *sql.DB, driver string, dbName string, tables []model.Table) []model.Table {
	for idx, table := range tables {
		columns := extract.GetColumnsFromRelation(db, dbName, driver, table.Name)
		tables[idx].Columns = append(tables[idx].Columns, columns...)
	}

	for i := 0; i < len(tables); i++ {
		for j := 0; j < len(tables[i].Columns); j++ {
			colToCheck := tables[i].Columns[j]
			if helper.IsExistInPrimaryKeys(colToCheck.Name, tables[i].PrimaryKeys) ||
				helper.IsExistInForeignKeys(colToCheck.Name, tables[i].ForeignKeys) {
				tables[i].Columns = append(tables[i].Columns[:j], tables[i].Columns[j+1:]...)
				j = j - 1
			}
		}
	}

	return tables
}
