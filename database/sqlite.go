package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func Connection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
