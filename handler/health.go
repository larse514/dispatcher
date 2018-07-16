package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping is a health check handler func
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
