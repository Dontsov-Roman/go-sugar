package db

import (
	"database/sql"
	"fmt"

	. "go-sugar/config"

	"github.com/go-sql-driver/mysql"
)

// SimpleRepo Simple Repository interface
type SimpleRepo interface {
	DeleteByID(int) bool
	GET()
}

// DB SQL DB Connect
var DB *sql.DB

// Connect main func for connect to DB
func Connect() *sql.DB {
	var err error
	dbConfig := mysql.NewConfig()
	dbConfig.User = Config.DB.Login
	dbConfig.Passwd = Config.DB.Password
	dbConfig.Addr = Config.DB.Addr
	dbConfig.DBName = Config.DB.Schema
	dbConfig.Net = Config.DB.Net
	DB, err = sql.Open("mysql", dbConfig.FormatDSN())
	// var str string = Config.DB.Login + ":" + Config.DB.Password + "@" + Config.DB.Net + "(" + Config.DB.Addr + ")" + "/" + Config.DB.Schema
	// DB, err = sql.Open("mysql", str)
	if err != nil {
		fmt.Println(err)
	} else {
		pimgErr := DB.Ping()
		if pimgErr == nil {
			fmt.Println("Succesfully connected to DB")
		} else {
			fmt.Println("Some problems appear while connecting to DB. ", pimgErr)
		}
	}
	return DB
}
func init() {
	Connect()
}
