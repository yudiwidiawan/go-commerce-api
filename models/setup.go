package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := "yudi_jcc_go_2021"
	PASS := "PasswordKu2021"
	HOST := "www.db4free.net"
	PORT := "3306"
	DBNAME := "db_commerce_jcc"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
