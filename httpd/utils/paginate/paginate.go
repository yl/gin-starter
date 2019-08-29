package paginate

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yangliulnn/gin-starter/configs"
)

type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"page"`
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func Paginator(context *gin.Context, query *gorm.DB, data interface{}) *Pagination {
	pagination := NewPagination()
	pagination.CurrentPage = getPage(context)
	pagination.PerPage = getPerPage(context)

	query.Count(&pagination.Total)
	query.Offset((pagination.CurrentPage - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(data)

	return pagination
}

func getPage(context *gin.Context) int {
	page, err := strconv.Atoi(context.Query(configs.Paginate.PageField))
	if err != nil || page == 0 {
		return 1
	}
	return page
}

func getPerPage(context *gin.Context) int {
	perPage, err := strconv.Atoi(context.Query(configs.Paginate.PerPageField))
	if err != nil || perPage == 0 {
		return configs.Paginate.DefaultPerPage
	}
	return perPage
}
