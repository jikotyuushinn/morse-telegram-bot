package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	. "morse-telegram-bot/middleware"
	"morse-telegram-bot/util"
)

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()
	
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}
	if update.Message.IsCommand() {
		switch command := update.Message.Command(); command  {
		case "help":
			fmt.Println(tgbotapi.NewMessage(update.Message.Chat.ID, "禁止幫助"))
		case "decode":
			fmt.Println(tgbotapi.NewMessage(update.Message.Chat.ID, "不準解碼"))
		case "encode":
			fmt.Println(tgbotapi.NewMessage(update.Message.Chat.ID, "不準編碼"))
		default:
			fmt.Println(tgbotapi.NewMessage(update.Message.Chat.ID, "錯誤命令"))
		}
	}
	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(util.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(util.WebhookHost + bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	
	router := gin.Default()
	router.Use(LogMiddleware())
	
	router.POST("/" + bot.Token, webhookHandler)
	
	err = router.Run()
	if err != nil {
		log.Println(err)
	}

	//r = CollectRoute(r)
}