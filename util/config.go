package util

import (
	tb "gopkg.in/tucnak/telebot.v2"
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

func InitBot() *tb.Bot {
	webhook := &tb.Webhook{
		Listen: ":" + Port,
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: WebhookHost,
		},
	}
	
	b, err := tb.NewBot(tb.Settings{
		Token:  AccessToken,
		Poller: webhook,
	})
	if err != nil {
		log.Fatal(err)
	}
	
	return b
}