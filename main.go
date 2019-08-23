package main

import (
	"go-trading/httpd/routers"
	"go-trading/services/database"
	"go-trading/services/database/migrations"
	"go-trading/services/redis"
)

func main() {
	database.Setup()
	migrations.Setup()
	redis.Setup()
	routers.Setup()

	defer database.Close()
}
