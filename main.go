/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:29:27
 * @LastEditors: Kasper
 * @LastEditTime: 2020-12-28 18:50:59
 * @Description: file content
 */

package main

import (
	"github.com/gin-gonic/gin"
	. "morse-telegram-bot/middleware"
	"morse-telegram-bot/util"
)

func main() {

	r := gin.Default()
	r.Use(ComputeCostTime, CORSMiddleware())
	r = CollectRoute(r)

	panic(r.Run(util.SeverPort))
}