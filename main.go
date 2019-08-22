package main

import (
	"go-trading/database"
	"go-trading/models"
	"go-trading/routers"
)

func main() {
	db := database.Connection()
	db.AutoMigrate(&models.User{})
	defer db.Close()

	server := routers.Setup(db)
	_ = server.Run()
}
