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
	ToLinkable           bool   `json:"toLinkable"`
	ToLinkableDuplicates bool   `json:"toLinkableDuplicates"`
	Key                  int    `json:"key"`
	Location             string `json:"location"`
	FromMaxLinks         int    `json:"fromMaxLinks"`
	Underline            bool   `json:"underline"`
}

type Link struct {
	From       int    `json:"from"`
	To         int    `json:"to"`
	Text       string `json:"text"`
	IsTotal    bool   `json:"isTotal"`
	IsParent   bool   `json:"isParent"`
	IsOne      bool   `json:"isOne"`
	IsDisjoint bool   `json:"isDisjoint"`
}
