package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CommonController struct{}

func (c *CommonController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (c *CommonController) Time(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"time": time.Now().Format(time.RFC3339),
	})
}
