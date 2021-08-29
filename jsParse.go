package bot

import (
	"github.com/robertkrimen/otto"
	"os"
)

func JsParser(filePath string, functionName string, args ...interface{}) (string, error) {

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		return "", err
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		return "", err
	}

	return value.String(), nil
}
