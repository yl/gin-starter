package paginate

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-trading/configs"
	"strconv"
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
	s, ok := context.GetQuery(configs.Paginate.PageField)
	if !ok {
		return 1
	}
	page, err := strconv.Atoi(s)
	if err != nil || page == 0 {
		return 1
	}
	return page
}

func getPerPage(context *gin.Context) int {
	s, ok := context.GetQuery(configs.Paginate.PerPageField)
	if !ok {
		return configs.Paginate.DefaultPerPage
	}
	perPage, err := strconv.Atoi(s)
	if err != nil || perPage == 0 {
		return configs.Paginate.DefaultPerPage
	}
	return perPage
}
