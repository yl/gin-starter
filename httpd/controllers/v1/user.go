package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yangliulnn/gin-starter/httpd/models"
	"github.com/yangliulnn/gin-starter/httpd/responses"
	"github.com/yangliulnn/gin-starter/httpd/utils/paginate"
	"github.com/yangliulnn/gin-starter/services/database"
)

type UserController struct{}

func (c *UserController) List(context *gin.Context) {
	user := &models.User{}
	users := &models.Users{}

	query := database.DB.Model(user).Order("id DESC")
	pagination := paginate.Paginator(context, query, users)

	response := responses.NewResponse()
	response.Paginate(context, users.Transformer(), pagination)
	return
}

func (c *UserController) Show(context *gin.Context) {
	response := responses.NewResponse()

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		response.InternalServerError(context, err)
		return
	}

	user := models.NewUser()
	err = user.FirstBy("id", id)
	if gorm.IsRecordNotFoundError(err) {
		response.NotFound(context)
		return
	}
	if err != nil {
		response.InternalServerError(context, err)
		return
	}

	response.Item(context, user.Transformer())
	return
}
