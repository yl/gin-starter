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
	config   *ini.File
	App      = &app{}
	Database = &database{}
	Redis    = &redis{}
	Paginate = &paginate{}
	JWT      = &jwt{}
)

func Setup() {
	var err error
	config, err = ini.Load("./configs/debug.ini")
	if err != nil {
		log.Fatalln(err)
	}

	mapTo("app", App)
	mapTo("database", Database)
	mapTo("redis", Redis)
	mapTo("paginate", Paginate)
	mapTo("jwt", JWT)
}

func mapTo(section string, p interface{}) {
	err := config.Section(section).MapTo(p)
	if err != nil {
		log.Fatalln(err)
	}
}
