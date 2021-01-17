package util

import (
	"log"
	"os"
	"path/filepath"
)

var (
	AccessToken string
	WebhookHost string
	Port        string
	StaticPath  string
)


func init() {
	AccessToken = os.Getenv("ACCESS_TOKEN")
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	WebhookHost = os.Getenv("WEBHOOK_HOST")
	Port = os.Getenv("PORT")
	StaticPath = filepath.Join(currentPath, os.Getenv("FILE_PATH"))
}