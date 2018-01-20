package goBotUtils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ContextJsonParam = "jsonParam" //параметры в web запросах
)

func HttpSuccess(c *gin.Context, res interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"result": res,
	})
}

func HttpError(c *gin.Context, status int, message string) {
	println("httpError", status, message)
	c.JSON(status, gin.H{
		"ok":      false,
		"message": message,
	})
	c.Abort()
}
