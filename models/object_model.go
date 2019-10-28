package models

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"runtime"
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
	if runtime.GOOS == "windows" {
		repairString = repairString
	}
	/*if err := xml.Unmarshal([]byte(repairString), &history); err != nil {
		return err, history
	}*/
	d := xml.NewDecoder(bytes.NewReader([]byte(repairString)))
	d.CharsetReader = identReader
	if err := d.Decode(&object); err != nil {
		panic(err)
	}
	return nil, object
}
