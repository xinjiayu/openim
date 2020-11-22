package model

type DataInfo struct {
	From string `orm:"from"  json:"from"`
	Data string `orm:"data"  json:"data"`
	Type string `orm:"type"  json:"type"`
}
