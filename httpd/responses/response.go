package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangliulnn/gin-starter/configs"
	. "github.com/yangliulnn/gin-starter/httpd/utils/log"
	"github.com/yangliulnn/gin-starter/httpd/utils/paginate"
	"github.com/yangliulnn/gin-starter/httpd/validators"
	"gopkg.in/go-playground/validator.v9"
)

type Item *map[string]interface{}

type itemResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Item   `json:"data"`
}

type Collection []Item

type collectionResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    Collection `json:"data"`
}

type Response struct{}

type paginateResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    Collection           `json:"data"`
	Meta    *paginate.Pagination `json:"meta"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Item(context *gin.Context, item Item) {
	context.JSON(http.StatusOK, &itemResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    item,
	})
}

func (r *Response) Collection(context *gin.Context, collection Collection) {
	context.JSON(http.StatusOK, &collectionResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    collection,
	})
}

func (r *Response) Paginate(context *gin.Context, collection Collection, paginate *paginate.Pagination) {
	context.JSON(http.StatusOK, &paginateResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    collection,
		Meta:    paginate,
	})
}

func (r *Response) Data(context *gin.Context, data *gin.H) {
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"data":    data,
	})
}

func (r *Response) Error(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//NotFound 404
func (r *Response) NotFound(context *gin.Context) {
	r.Error(context, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

//UnprocessableEntity 422
func (r *Response) UnprocessableEntity(context *gin.Context, err error) {
	lang := context.GetHeader("Accept-Language")
	trans, _ := validators.UT.GetTranslator(lang)
	errs := err.(validator.ValidationErrors)
	for _, e := range errs.Translate(trans) {
		r.Error(context, http.StatusUnprocessableEntity, e)
		break
	}
}

//InternalServerError 500
func (r *Response) InternalServerError(context *gin.Context, err error) {
	Log.Error(err)

	var message string
	if configs.App.Mode == "debug" {
		message = err.Error()
	} else {
		message = http.StatusText(http.StatusInternalServerError)
	}
	r.Error(context, http.StatusInternalServerError, message)
}
