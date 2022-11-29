package main

import (
	"fmt"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// global var so that we use same var in Handlers also
var DataBase *gorm.DB

// var DataBase *sql.DB

// url from where the connection has to be established "DRIVER"

// var urlDSN = "root:Test@123@tcp(localhost:3306)/db_cognologix"

var urlDSN = "root:Test@123@tcp(localhost:3306)/db_cognologix?charset=utf8&parseTime=True"

var err error // for error

func DataMigration() {

	DataBase, err = gorm.Open("mysql", urlDSN) // open and close
	if err != nil {
		fmt.Println(err)
		panic("connection failed")
	}

	// defer DataBase.Close()

	DataBase.AutoMigrate(&Employee{}) // Creates table
	fmt.Println("Connect!")
}
