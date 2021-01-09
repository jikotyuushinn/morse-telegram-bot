/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:44:50
 * @LastEditors: Kasper
 * @LastEditTime: 2020-12-28 13:45:00
 * @Description: file content
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"morse-telegram-bot/util"
	"net/http"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		log.Printf("%s - %s - %s - %s - %d - %s", c.Request.UserAgent(), c.ClientIP(),
			c.Request.RequestURI, c.Request.Method, c.Writer.Status(), cost)
	}
	
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", util.ServerIP)
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}