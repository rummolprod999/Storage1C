package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Objects struct {
	Records []ObjectRecord `xml:"Records>Record"`
}

type ObjectRecord struct {
	ObjId      string `xml:"OBJID"`
	ClassId    string `xml:"CLASSID"`
	SelfVerNum int    `xml:"SELFVERNUM"`
}

func (t *Objects) CreateModel(path string) (error, Objects) {
	var object Objects
	xmlFile, err := os.Open(path)
	if err != nil {
		return err, object
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err, object
	}
	repairString := ReplaceBadSymbols(string(byteValue))
	if err := xml.Unmarshal([]byte(repairString), &object); err != nil {
		return err, object
	}
	return nil, object
}
