package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	QueryError  = 4001
	ServerError = 5000
)

var codeMessage = map[int]string {
	QueryError:  "QUERY PARAM ERROR",
	ServerError: "SERVER ERROR",
}

func GetCodeMessage(code int) string {
	return codeMessage[code]
}

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"message": msg,
	})
}

func SuccessResponse(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, 200, data, "request succeed")
}

func FailResponse(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, 400, data, "request failed")
}