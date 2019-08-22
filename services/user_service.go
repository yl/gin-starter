package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-trading/helpers"
	"go-trading/models"
	"go-trading/repositories"
	"go-trading/requests"
	"go-trading/responses"
	"net/http"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request requests.RegisterRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
		}

		var user models.User
		user.Mobile = request.Mobile
		user.Password = helpers.PasswordHash(request.Password)

		userRepository := repositories.NewUserRepository(db)
		userRepository.Save(&user)

		c.JSON(http.StatusOK, responses.UserResponse(&user))
	}
}
