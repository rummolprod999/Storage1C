package models

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

type Versions struct {
	Records []VersionRecord `xml:"Records>Record"`
}

type VersionRecord struct {
	VerNum    int        `xml:"VERNUM"`
	UserId    string     `xml:"USERID"`
	VerDate   customTime `xml:"VERDATE"`
	Code      string     `xml:"CODE"`
	Comment   string     `xml:"COMMENT"`
	VersionId string     `xml:"VERSIONID"`
}

type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	//parse, _ := time.Parse("2006-01-02T15:04:05", v)
	location, _ := time.LoadLocation("Europe/Moscow")
	parse, err := time.ParseInLocation("2006-01-02T15:04:05", v, location)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

func (t *Versions) CreateModel(path string) (error, Versions) {
	var versions Versions
	xmlFile, err := os.Open(path)
	if err != nil {
		return err, versions
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err, versions
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
	if err := d.Decode(&versions); err != nil {
		panic(err)
	}
	return nil, versions
}
