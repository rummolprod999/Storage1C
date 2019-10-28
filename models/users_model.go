package models

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"runtime"
)

type Users struct {
	Records []UserRecord `xml:"Records>Record"`
}

type UserRecord struct {
	UserId string `xml:"USERID"`
	Name   string `xml:"NAME"`
}

func (t *Users) CreateModel(path string) (error, Users) {
	var users Users
	xmlFile, err := os.Open(path)
	if err != nil {
		return err, users
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err, users
	}
	repairString := ReplaceBadSymbols(string(byteValue))
	if runtime.GOOS == "windows" {
		repairString = repairString
	}
	d := xml.NewDecoder(bytes.NewReader([]byte(repairString)))
	d.CharsetReader = identReader
	if err := d.Decode(&users); err != nil {
		panic(err)
	}
	return nil, users
}
