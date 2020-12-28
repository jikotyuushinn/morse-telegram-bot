/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:29:27
 * @LastEditors: Kasper
 * @LastEditTime: 2020-12-28 18:49:30
 * @Description: file content
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	. "morse-telegram-bot/middleware"
	"morse-telegram-bot/util"
)

func main() {
	util.InitConfig()

	r := gin.Default()
	r.Use(ComputeCostTime, CORSMiddleware())
	r = CollectRoute(r)

	panic(r.Run())
}