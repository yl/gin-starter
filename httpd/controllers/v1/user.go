package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-trading/httpd/models"
	"go-trading/httpd/responses"
	"log"
	"strconv"
)

type UserController struct{}

func (c *UserController) List(context *gin.Context) {
	response := responses.NewResponse()

	userModel := &models.User{}
	users, err := userModel.All()
	if gorm.IsRecordNotFoundError(err) {
		response.NotFound(context)
		return
	}
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}

	response.Collection(context, users.Transformer())
	return
}

func (c *UserController) Show(context *gin.Context) {
	response := responses.NewResponse()

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}

	userModel := &models.User{}
	user, err := userModel.FindById(uint(id))
	if gorm.IsRecordNotFoundError(err) {
		response.NotFound(context)
		return
	}
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}

	response.Item(context, user.Transformer())
	return
}
