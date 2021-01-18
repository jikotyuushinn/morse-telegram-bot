package util

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
)

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
