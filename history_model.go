package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type History struct {
	Records []HistoryRecord `xml:"Records>Record"`
}

type HistoryRecord struct {
	ObjId      string `xml:"OBJID"`
	VerNum     int    `xml:"VERNUM"`
	SelfVerNum int    `xml:"SELFVERNUM"`
	ObjVerId   string `xml:"OBJVERID"`
	ParentId   string `xml:"PARENTID"`
	ObjName    string `xml:"OBJNAME"`
	ObjPos     int    `xml:"OBJPOS"`
	Removed    bool   `xml:"REMOVED"`
}

func (t *History) CreateModel(path string) (error, History) {
	var history History
	xmlFile, err := os.Open(path)
	if err != nil {
		return err, history
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err, history
	}
	repairString := ReplaceBadSymbols(string(byteValue))
	if err := xml.Unmarshal([]byte(repairString), &history); err != nil {
		return err, history
	}
	return nil, history
}
