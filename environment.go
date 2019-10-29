package main

import (
	"Storage1C/models"
	"flag"
	"fmt"
	"github.com/jinzhu/configor"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Filelog string

var DirLog = "log_storage1c"
var DirTemp = "temp_storage1c"
var PathTemp = ""
var FileLog Filelog
var SecretKey = "11111111" //change it
var mutex sync.Mutex

func GetArgs() {
	flag.Parse()
	ArgDebug := flag.Arg(0)
	if ArgDebug == "debug" {
		models.Debug = true
	}
}

func CreateEnv() {
	CreateLogFile()
	CreateTempDir()
	ReadConfig()
	GetArgs()
}

func CreateLogFile() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dirlog := filepath.FromSlash(fmt.Sprintf("%s/%s", dir, DirLog))
	if _, err := os.Stat(dirlog); os.IsNotExist(err) {
		err := os.MkdirAll(dirlog, 0711)

		if err != nil {
			fmt.Println("cannot create logs dir")
			os.Exit(1)
		}
	}
	t := time.Now()
	ft := t.Format("2006-01-02")
	FileLog = Filelog(filepath.FromSlash(fmt.Sprintf("%s/log_storage_1c_%v.log", dirlog, ft)))
}

func CreateTempDir() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	PathTemp = filepath.FromSlash(fmt.Sprintf("%s/%s", dir, DirTemp))
	if _, err := os.Stat(PathTemp); os.IsNotExist(err) {
		err := os.MkdirAll(PathTemp, 0711)

		if err != nil {
			fmt.Println("cannot create temp dir")
			os.Exit(1)
		}
	} else {
		err = os.RemoveAll(PathTemp)
		if err != nil {
			fmt.Println("cannot delete temp dir")
			os.Exit(1)
		}
		err := os.MkdirAll(PathTemp, 0711)
		if err != nil {
			fmt.Println("cannot create temp dir")
			os.Exit(1)
		}
	}
}

func ReadConfig() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	settingsPath := filepath.FromSlash(fmt.Sprintf("%s/%s", dir, "settings.json"))
	err := configor.Load(&Config, settingsPath)
	if err != nil {
		panic(err)
	}
}
func Logging(args ...interface{}) {
	mutex.Lock()
	file, err := os.OpenFile(string(FileLog), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("error writing to log file", err)
		return
	}
	fmt.Fprintf(file, "%v  ", time.Now())
	for _, v := range args {

		fmt.Fprintf(file, " %v", v)
	}
	//fmt.Fprintf(file, " %s", UrlXml)
	fmt.Fprintln(file, "")
	mutex.Unlock()
}
