package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	. "morse-telegram-bot/util"
	"net/http"
)


func Decode(c *gin.Context) {
	morseCode := c.Query("morseCode")
	if morseCode == "" {
		Response(c, http.StatusUnprocessableEntity, QueryError, nil, GetCodeMessage(QueryError))
		return
	}
	
	res, err := JsParser(StaticPath, "xmorse.decode", morseCode)
	if err != nil {
		Response(c, http.StatusInternalServerError, ServerError, nil, GetCodeMessage(ServerError))
		return
	}
	
	SuccessResponse(c, gin.H{
		"text": res,
	})

}

func Encode(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		Response(c, http.StatusUnprocessableEntity, QueryError, nil, GetCodeMessage(QueryError))
		return
	}
	
	res, err := JsParser(StaticPath, "xmorse.encode", text)
	if err != nil {
		Response(c, http.StatusInternalServerError, ServerError, nil, GetCodeMessage(ServerError))
		return
	}
	
	SuccessResponse(c, gin.H{
		"morseCode": res,
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