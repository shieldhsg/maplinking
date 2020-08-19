package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Db() *gorm.DB {
	db, err := gorm.Open("mysql", Config.Db)
	if err != nil {
		panic(err.Error())
	}
	return db
}
