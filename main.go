package main

import (
	"go-trading/configs"
	"go-trading/httpd/routers"
	"go-trading/services/database"
	"go-trading/services/database/migrations"
	"go-trading/services/redis"
)

func main() {
	database.Setup()
	defer database.Close()
	migrations.Setup()
	redis.Setup()

	router := routers.Setup()
	_ = router.Run(configs.App.Addr)
}
