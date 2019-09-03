package configs

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type app struct {
	Mode       string `ini:"mode"`
	Addr       string `ini:"addr"`
	Timezone   string `ini:"timezone"`
	TimeFormat string `ini:"time_format"`
	Local      string `ini:"local"`
}

type database struct {
	Driver   string `ini:"driver"`
	Host     string `ini:"host,omitempty"`
	Port     string `ini:"port,omitempty"`
	Database string `ini:"database"`
	Username string `ini:"username,omitempty"`
	Password string `ini:"password,omitempty"`
}

type redis struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password,omitempty"`
	DB       int    `ini:"db"`
}

type paginate struct {
	DefaultPerPage int `ini:"default_per_page"`
}

type jwt struct {
	TTL time.Duration `ini:"ttl"`
}

var (
	App      = &app{}
	Database = &database{}
	Redis    = &redis{}
	Paginate = &paginate{}
	JWT      = &jwt{}
)

func init() {
	config, err := ini.Load("./configs/debug.ini")
	if err != nil {
		log.Fatalln(err)
	}

	mapTo(config, App, "app")
	mapTo(config, Database, "database")
	mapTo(config, Redis, "redis")
	mapTo(config, Paginate, "paginate")
	mapTo(config, JWT, "jwt")
}

func mapTo(cfg *ini.File, p interface{}, section string) {
	err := cfg.Section(section).MapTo(p)
	if err != nil {
		log.Fatalln(err)
	}
}
