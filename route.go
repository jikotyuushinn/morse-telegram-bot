/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:41:53
 * @LastEditors: Kasper
 * @LastEditTime: 2020-12-28 13:42:17
 * @Description: file content
 */

package main

import (
	"github.com/gin-gonic/gin"
	. "morse-telegram-bot/controller"
	"morse-telegram-bot/util"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	MorseGroup := r.Group("/api")
	{
		MorseGroup.GET("/encode", Encode)
		MorseGroup.GET("/decode", Decode)
	}
	r.NoRoute(func(c *gin.Context) {
		util.FailResponse(c, gin.H{
			"data": "呵呵。",
		})
	})

	return r
}