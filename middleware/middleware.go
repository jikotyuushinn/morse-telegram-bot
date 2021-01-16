package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		log.Printf("%s - %s - %s - %s - %d - %s", c.Request.UserAgent(), c.ClientIP(),
			c.Request.RequestURI, c.Request.Method, c.Writer.Status(), cost)
	}
}
