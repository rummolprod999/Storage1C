package models

import (
	"bytes"
	"encoding/xml"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
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
	var byteValue []byte
	if runtime.GOOS == "windows" {
		O := transform.NewReader(xmlFile, unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder())
		byteValue, err = ioutil.ReadAll(O)
		if err != nil {
			if Debug {
				panic(err)
			}
			return err, users
		}
		repairString := string(byteValue)
		repairString = ReplaceBadSymbols(repairString)
		d := xml.NewDecoder(bytes.NewReader([]byte(repairString)))
		d.CharsetReader = identReader
		if err := d.Decode(&users); err != nil {
			if Debug {
				panic(err)
			}
			panic(err)
		}
		return nil, users
	} else {
		byteValue, err = ioutil.ReadAll(xmlFile)
		if err != nil {
			return err, users
		}
		repairString := ReplaceBadSymbols(string(byteValue))
		d := xml.NewDecoder(bytes.NewReader([]byte(repairString)))
		d.CharsetReader = identReader
		if err := d.Decode(&users); err != nil {
			if Debug {
				panic(err)
			}
			panic(err)
		}
		return nil, users
	}

}
