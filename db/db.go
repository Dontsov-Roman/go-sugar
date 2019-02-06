package db

import (
	"database/sql"
	"fmt"

	. "../config"
	_ "github.com/go-sql-driver/mysql"
)

// DBConnect SQL DB Connect
var DB *sql.DB

func Connect() *sql.DB {
	var err error
	var str string = Config.DB.Login + ":" + Config.DB.Password + "@/" + Config.DB.Schema
	DB, err = sql.Open("mysql", str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully connected to DB")
	}
	return DB
}
func init() {
	Connect()
}
