package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

//ConfigFile struct for reading the configuration json file
type ConfigFile struct {
	TimeInterval time.Duration `json:"timeInterval"`
	Cookie       string        `json:"cookie"`
	URLs         []string      `json:"urls"`
}

//Config struct
var Config ConfigFile

//ReadConfig data
func ReadConfig() (configFile ConfigFile) {
	//get binary dir
	//os.Args doesn't work the way we want with "go run". You can use next line
	//for local dev, but use the original for production.
	dir, err := filepath.Abs("./") //local dev
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //production-binary
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading config file...")
	file := dir + "/config.json"
	configFileData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Can't read config.json, because: ", err)
	}

	log.Println("Parsing config file...")
	err = json.Unmarshal(configFileData, &configFile)
	if err != nil {
		log.Fatal("Cant parse json, because: ", err)
	}

	return
}
