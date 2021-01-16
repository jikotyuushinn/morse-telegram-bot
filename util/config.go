/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:43:24
 * @LastEditors: Kasper
 * @LastEditTime: 2021-01-07 15:24:00
 * @Description: file content
 */

package util

import (
	"os"
)

var (
	AccessToken string
	StaticPath string
)


func init() {
	AccessToken = os.Getenv("ACCESS_TOKEN")
	StaticPath = os.Getenv("FILE_PATH")

}