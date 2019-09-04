package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yangliulnn/gin-starter/configs"
	"github.com/yangliulnn/gin-starter/httpd/models"
	"github.com/yangliulnn/gin-starter/httpd/requests"
	"github.com/yangliulnn/gin-starter/httpd/responses"
	"github.com/yangliulnn/gin-starter/httpd/utils"
	"github.com/yangliulnn/gin-starter/httpd/utils/jwt"
	. "github.com/yangliulnn/gin-starter/httpd/utils/log"
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

	user := models.NewUser()
	err = user.FirstBy("mobile", request.Mobile)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		Log.Error(err)
		response.InternalServerError(context)
		return
	}
	if user.ID > 0 {
		response.Error(context, http.StatusBadRequest, "手机号已注册")
		return
	}

	user.Mobile = request.Mobile
	user.Password, err = utils.NewPassword().Hash(request.Password)
	if err != nil {
		Log.Error(err)
		response.InternalServerError(context)
		return
	}

	err = user.Save()
	if err != nil {
		Log.Error(err)
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

	user := models.NewUser()
	err = user.FirstBy("mobile", request.Mobile)
	if gorm.IsRecordNotFoundError(err) {
		response.Error(context, http.StatusBadRequest, "账号不存在")
		return
	}
	if err != nil {
		Log.Error(err)
		response.InternalServerError(context)
		return
	}

	err = utils.NewPassword().Check(user.Password, request.Password)
	if err != nil {
		response.Error(context, http.StatusBadRequest, "登录密码错误")
		return
	}

	token, err := jwt.Generate(user)
	if err != nil {
		Log.Error(err)
		response.InternalServerError(context)
		return
	}

	response.Data(context, &gin.H{
		"token":      token,
		"expired_at": utils.NewTime().Format(time.Now().Add(configs.JWT.TTL * time.Second)),
	})
	return
}
