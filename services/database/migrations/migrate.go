package migrations

import (
	"go-trading/httpd/models"
	"go-trading/services/database"
)

func Setup() {
	database.DB.AutoMigrate(
		&models.User{},
	)
}
