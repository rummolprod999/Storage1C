package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
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
	if err := xml.Unmarshal([]byte(repairString), &users); err != nil {
		return err, users
	}
	return nil, users
}
