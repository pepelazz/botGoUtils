package goBotUtils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ContextJsonParam         = "jsonParam"         //параметры в web запросах
	ContextJsonParamFldParam = "jsonParamFldParam" //поле params в параметры в web запросах
	ContextUserRole          = "userRole"
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

func CheckIsAdmin(c *gin.Context) bool {
	userRole, ok := c.Get(ContextUserRole)
	if !ok {
		HttpError(c, http.StatusMethodNotAllowed, "not found current user")
		return false
	}
	if userRole != "admin" {
		HttpError(c, http.StatusMethodNotAllowed, "not enough rights")
		return false
	}
	return true
}
