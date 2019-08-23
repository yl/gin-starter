package responses

import (
	"github.com/gin-gonic/gin"
	"go-trading/httpd/validators"
	"gopkg.in/go-playground/validator.v9"
	//en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	//zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
	"net/http"
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

//NotFound 400
func (r *Response) NotFound(context *gin.Context) {
	r.Error(context, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

//UnprocessableEntity 422
func (r *Response) UnprocessableEntity(context *gin.Context, err error) {
	errs := err.(validator.ValidationErrors)
	trans,_:= validators.UT.GetTranslator("zh")
	for _, e := range errs.Translate(trans) {
		r.Error(context, http.StatusUnprocessableEntity, e)
		return
	}
}

//InternalServerError 500
func (r *Response) InternalServerError(context *gin.Context) {
	r.Error(context, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}
