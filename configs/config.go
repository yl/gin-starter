package configs

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	App      *app
	Database *database
	Redis    *redis
	Paginate *paginate
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
	PageField      string `ini:"page_field"`
	PerPageField   string `ini:"per_page_field"`
	DefaultPerPage int    `ini:"default_per_page"`
}

func init() {
	cfg, err := ini.Load("./configs/debug.ini")
	if err != nil {
		log.Fatalln(err)
	}

	App = &app{}
	err = cfg.Section("app").MapTo(App)
	if err != nil {
		log.Fatalln(err)
	}

	Database = &database{}
	err = cfg.Section("database").MapTo(Database)
	if err != nil {
		log.Fatalln(err)
	}

	Redis = &redis{}
	err = cfg.Section("redis").MapTo(Redis)
	if err != nil {
		log.Fatalln(err)
	}

	Paginate = &paginate{}
	err = cfg.Section("paginate").MapTo(Paginate)
	if err != nil {
		log.Fatalln(err)
	}
}
