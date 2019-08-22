package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-trading/services"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/ping", services.Ping())
	r.GET("/time", services.Time())
	users := r.Group("/users")
	{
		users.POST("/register", services.Register(db))
	}

	return r
}
