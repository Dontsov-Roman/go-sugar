package db

import (
	"database/sql"
	"fmt"

	. "../config"
	_ "github.com/go-sql-driver/mysql"
)

// DBConnect SQL DB Connect
var DBConnect *sql.DB

func init() {
	var err error
	var str string = Config.DB.Login + ":" + Config.DB.Password + "@/" + Config.DB.DBname
	fmt.Println(str)
	DBConnect, err = sql.Open("mysql", str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully connected to DB")
	}
}
