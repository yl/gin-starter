package migrations

import (
	"github.com/yangliulnn/gin-starter/httpd/models"
	"github.com/yangliulnn/gin-starter/services/database"
)

func Setup() {
	database.DB.AutoMigrate(
		models.NewUser(),
	)
}
