package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func init() {
	user := os.Getenv("MYSQL_USERNAME")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB_NAME")
	var err error
	if DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)); err != nil {
		log.Fatal("failed to connect db: " + dbname)
	}
}

func clearDB() {
	models := []interface{}{User{}, Customer{}, Store{}}
	for _, m := range models {
		DB.Exec(fmt.Sprintf("TRUNCATE TABLE %v;", DB.NewScope(m).TableName()))
	}
}

func CloseDB() {
	DB.Close()
}
