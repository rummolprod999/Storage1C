package main

import (
	"Storage1C/models"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Storage1C struct {
	storageName  string
	storagePath  string
	pathUsers    string
	pathHistory  string
	pathObjects  string
	pathVersions string
}

func (t *Storage1C) Init() {
	for _, v := range Config.Storages {
		if v.Name == FirstArgument {
			t.storageName = v.Name
			t.storagePath = v.Path
			return
		}
	}
	panic("storage not found")
}

func (t *Storage1C) Run() {
	err := t.RunExternalProgram()
	if err != nil {
		Logging(err)
		return
	}
	err = t.CheckXml()
	if err != nil {
		Logging(err)
		return
	}
	err = t.createModels()
	if err != nil {
		Logging(err)
		return
	}
}

func (t *Storage1C) RunExternalProgram() error {
	command := exec.Command(Config.PathCTool1cd, t.storagePath, "-ne", "-ex", PathTemp, "USERS,HISTORY,VERSIONS,OBJECTS")
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		return err
	}
	//fmt.Println(out.String())
	return nil
}

func (t *Storage1C) CheckXml() error {
	t.pathUsers = filepath.FromSlash(fmt.Sprintf("%s/%s", PathTemp, "USERS.xml"))
	if _, err := os.Stat(t.pathUsers); os.IsNotExist(err) {
		return err
	}
	t.pathHistory = filepath.FromSlash(fmt.Sprintf("%s/%s", PathTemp, "HISTORY.xml"))
	if _, err := os.Stat(t.pathHistory); os.IsNotExist(err) {
		return err
	}
	t.pathVersions = filepath.FromSlash(fmt.Sprintf("%s/%s", PathTemp, "VERSIONS.xml"))
	if _, err := os.Stat(t.pathVersions); os.IsNotExist(err) {
		return err
	}
	t.pathObjects = filepath.FromSlash(fmt.Sprintf("%s/%s", PathTemp, "OBJECTS.xml"))
	if _, err := os.Stat(t.pathObjects); os.IsNotExist(err) {
		return err
	}
	return nil
}

func (t *Storage1C) createModels() error {
	var users models.Users
	err, users := users.CreateModel(t.pathUsers)
	if err != nil {
		return err
	}

	var history models.History
	err, history = history.CreateModel(t.pathHistory)
	if err != nil {
		return err
	}

	var objects models.Objects
	err, objects = objects.CreateModel(t.pathObjects)
	if err != nil {
		return err
	}

	var versions models.Versions
	err, versions = versions.CreateModel(t.pathVersions)
	if err != nil {
		return err
	}
	for _, v := range history.Records {
		fmt.Println("%+v", v)
	}
	return nil
}
