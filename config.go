package bot

import (
	"log"
	"os"
	"path/filepath"
)

var (
	AccessToken string
	StaticPath  string
)

func InitConfig() {
	AccessToken = os.Getenv("ACCESS_TOKEN")
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	StaticPath = filepath.Join(currentPath, os.Getenv("FILE_PATH"))
}
