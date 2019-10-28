package models

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"runtime"
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
	if runtime.GOOS == "windows" {
		repairString = repairString
	}
	/*if err := xml.Unmarshal([]byte(repairString), &history); err != nil {
		return err, history
	}*/
	d := xml.NewDecoder(bytes.NewReader([]byte(repairString)))
	d.CharsetReader = identReader
	if err := d.Decode(&history); err != nil {
		panic(err)
	}
	return nil, history
}
