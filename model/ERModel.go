package model

type ERModel struct {
	Class         string `json:"class"`
	NodeDataArray []Node `json:"nodeDataArray"`
	LinkDataArray []Link `json:"linkDataArray"`
}

type Node struct {
	Text                 string `json:"text"`
	Color                string `json:"color"`
	Figure               string `json:"figure"`
	Width                int    `json:"widht"`
	Height               int    `json:"height"`
	FromLinkable         bool   `json:"fromLinkable"`
	ToLinkableDuplicates bool   `json:"toLinkableDuplicates"`
	Key                  int    `json:"key"`
	Location             string `json:"location"`
}

type Link struct {
	From    int    `json:"from"`
	To      int    `json:"to"`
	Text    string `json:"text"`
	IsTotal bool   `json:"isTotal"`
	IsOne   bool   `json:"isOne"`
}
