/*
 * @Author: Kasper
 * @Date: 2020-12-28 13:46:06
 * @LastEditors: Kasper
 * @LastEditTime: 2020-12-28 13:47:41
 * @Description: file content
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"morse-telegram-bot/util"
	"net/http"
)


func Decode(c *gin.Context) {
	morseCode := c.Query("morseCode")
	res, err := JsParser(util.StaticPath, "xmorse.decode", morseCode)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": res,
			"error": err,
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func Encode(c *gin.Context) {
	text := c.Query("text")
	res, err := JsParser(util.StaticPath, "xmorse.encode", text)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": res,
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func JsParser(filePath string, functionName string, args... interface{}) (result string, err error) {
	
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		panic(err)
	}
	
	return value.String(), nil
}