package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-trading/httpd/models"
	"go-trading/httpd/requests"
	"go-trading/httpd/responses"
	"go-trading/httpd/utils"
	"log"
	"net/http"
	"time"
)

type AuthController struct{}

func (c *AuthController) Register(context *gin.Context) {
	response := responses.NewResponse()

	request := requests.NewRegisterRequest()
	err := context.ShouldBindJSON(request)
	if err != nil {
		response.UnprocessableEntity(context, err)
		return
	}

	userModel := &models.User{}
	password, err := utils.NewPassword().Hash(request.Password)
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}
	user, err := userModel.Insert(request.Mobile, password)
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}

	response.Item(context, user.Transformer())
	return
}

func (c *AuthController) Login(context *gin.Context) {
	response := responses.NewResponse()

	request := requests.NewLoginRequest()
	err := context.ShouldBindJSON(request)
	if err != nil {
		response.UnprocessableEntity(context, err)
		return
	}

	userModel := &models.User{}
	user, err := userModel.FindOne(map[string]interface{}{
		"mobile": request.Mobile,
	})
	if gorm.IsRecordNotFoundError(err) {
		response.Error(context, http.StatusBadRequest, "账号不存在")
		return
	}
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}

	err = utils.NewPassword().Check(user.Password, request.Password)
	if err != nil {
		response.Error(context, http.StatusBadRequest, "登录密码错误")
		return
	}

	token, err := utils.NewToken().Generate(user.ID)
	if err != nil {
		log.Println(err)
		response.InternalServerError(context)
		return
	}

	response.Data(context, &gin.H{
		"token":      token,
		"expired_at": utils.NewTime().Format(time.Now().Add(24 * time.Hour)),
	})
	return
}
