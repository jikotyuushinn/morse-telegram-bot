/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:43:24
 * @LastEditors: Kasper
 * @LastEditTime: 2021-01-07 15:24:00
 * @Description: file content
 */

package util

import (
	"github.com/spf13/viper"
	"log"
)

var (
	ServerIP string
	SeverPort string
	StaticPath string
)


func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("init config file error: %v", err)
	}
	
	initConfig()
}

func initConfig() {
	ServerIP = viper.GetString("Server.ip")
	SeverPort = viper.GetString("Server.port")
	StaticPath = viper.GetString("Static.filePath")
}