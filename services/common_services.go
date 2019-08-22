package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func Time() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().Format(time.RFC3339),
		})
	}
}
