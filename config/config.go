package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// DBConfig struct
type dBConfig struct {
	Login    string
	Password string
	DBname   string
}

// ConfigStruct Main config struct
type configStruct struct {
	Title string
	DB    dBConfig
}

// Config App's Configuration
var Config configStruct

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened config.json")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	errJSON := json.Unmarshal([]byte(byteValue), &Config)
	if errJSON != nil {
		fmt.Println(errJSON)
	}
	fmt.Println(Config)
}
