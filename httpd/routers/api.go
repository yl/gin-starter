package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-trading/configs"
	"go-trading/httpd/controllers/v1"
	"go-trading/httpd/validators"
)

func Setup() *gin.Engine {
	gin.SetMode(configs.App.Mode)

	router := gin.New()

	binding.Validator = new(validators.DefaultValidator)

	router.Use(gin.Logger(), gin.Recovery())

	routerV1 := router.Group("/api/v1")
	{
		commonController := new(v1.CommonController)
		routerV1.GET("/ping", commonController.Ping)
		routerV1.GET("/time", commonController.Time)

		authController := new(v1.AuthController)
		auth := routerV1.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		userController := new(v1.UserController)
		users := routerV1.Group("/users")
		{
			users.GET("/", userController.List)
			users.GET("/:id", userController.Show)
		}
	}

	return router
}
