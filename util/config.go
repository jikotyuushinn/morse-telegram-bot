/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:43:24
 * @LastEditors: Kasper
 * @LastEditTime: 2020-12-28 13:43:32
 * @Description: file content
 */

package util

import (
	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}