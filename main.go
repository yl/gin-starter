package main

import (
	"github.com/yangliulnn/gin-starter/configs"
	"github.com/yangliulnn/gin-starter/httpd/routers"
	"github.com/yangliulnn/gin-starter/httpd/utils/log"
	"github.com/yangliulnn/gin-starter/services/database"
	"github.com/yangliulnn/gin-starter/services/database/migrations"
	"github.com/yangliulnn/gin-starter/services/redis"
)

func main() {
	configs.Setup()
	log.Setup()
	database.Setup()
	defer database.Close()
	migrations.Setup()
	redis.Setup()

	router := routers.Setup()
	_ = router.Run(configs.App.Addr)
}
