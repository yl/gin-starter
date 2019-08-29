package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yangliulnn/gin-starter/configs"
)

var DB *gorm.DB

func Setup() {
	config := configs.Database

	dialect := config.Driver
	var args string
	switch dialect {
	case "sqlite3":
		args = config.Database
	case "mysql":
		args = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)
	default:
		log.Fatalln("不支持的数据库类型")
	}

	var err error
	DB, err = gorm.Open(dialect, args)
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {
	_ = DB.Close()
}
