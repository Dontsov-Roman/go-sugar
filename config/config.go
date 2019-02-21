package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig struct
type dBConfig struct {
	Login    string
	Password string
	Schema   string
	Addr     string
	Net      string
}

// ConfigStruct Main config struct
type configStruct struct {
	Title    string
	HTTPPort string
	DB       dBConfig
}

// Config App's Configuration
var Config configStruct

func loadFromJSON() {
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
}
func loadFromEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	Config.DB.Login = os.Getenv("MYSQL_USER")
	Config.DB.Password = os.Getenv("MYSQL_PASSWORD")
	Config.DB.Schema = os.Getenv("MYSQL_DATABASE")
	Config.DB.Addr = os.Getenv("MYSQL_ADDR")
	Config.DB.Net = os.Getenv("MYSQL_NET")
	Config.HTTPPort = os.Getenv("HTTP_PORT")
	fmt.Println(Config)
}
func init() {
	loadFromEnv()
}
