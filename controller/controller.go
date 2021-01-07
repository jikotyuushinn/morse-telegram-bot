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
	"log"
	"morse-telegram-bot/util"
	"net/http"
)


func Decode(c *gin.Context) {
	morseCode := c.Query("morseCode")
	if morseCode == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": util.QueryError,
			"message": util.GetCodeMessage(util.QueryError),
		})
	}
	
	res, err := JsParser(util.StaticPath, "xmorse.decode", morseCode)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": util.ServerError,
			"message": util.GetCodeMessage(util.ServerError),
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": util.Success,
		"data": res,
	})
}

func Encode(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": util.QueryError,
			"message": util.GetCodeMessage(util.QueryError),
		})
	}
	
	res, err := JsParser(util.StaticPath, "xmorse.encode", text)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": util.ServerError,
			"message": util.GetCodeMessage(util.ServerError),
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": util.Success,
		"data": res,
	})
}

func JsParser(filePath string, functionName string, args... interface{}) (string, error) {
	
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("read js file error: %v", err)
		return "", err
	}
	
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		log.Fatalf("launch js file error: %v", err)
		return "", err
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		log.Fatalf("execute js file error: %v", err)
		return "", err
	}
	
	return value.String(), nil
}