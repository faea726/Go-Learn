package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() { // Username and password to login sql
	d, err := gorm.Open("mysql", "faea726:142536/simplerest?charset=utf8YparseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
