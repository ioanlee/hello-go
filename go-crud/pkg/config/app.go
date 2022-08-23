package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var conn string = "ioanlee:password/tablename?charset=utf8&parseTime=True&loc=Local"

func Connect() {
	d, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
