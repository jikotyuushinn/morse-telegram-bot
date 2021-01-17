package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"morse-telegram-bot/controller"
	. "morse-telegram-bot/middleware"
	"morse-telegram-bot/util"
	"strings"
)

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
	
	router.POST("/" + bot.Token, func(c *gin.Context) {
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
		log.Printf("From: %+v Chat: %+v Text: %+v\n", update.Message.From, update.Message.Chat.ID,
			update.Message.Text)
		
		if update.Message.IsCommand() {
			switch command := update.Message.Command(); command {
			case "start":
				_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "不準開始"))
			case "help":
				_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "禁止幫助"))
			case "decode":
				morseCode := strings.TrimLeft(update.Message.Text, "/decode ")
				if morseCode == "" {
					_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "勸你最好有輸入"))
					break
				}
				res, _ := controller.JsParser(util.StaticPath, "xmorse.decode", morseCode)
				_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, res))
				_, _ = bot.Send(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
			case "encode":
				text := strings.TrimLeft(update.Message.Text, "/encode ")
				if text == "" {
					_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "勸你最好有輸入"))
					break
				}
				res, _ := controller.JsParser(util.StaticPath, "xmorse.encode", text)
				_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, res))
				_, _ = bot.Send(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
			default:
				_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "不要亂玩人家哦。"))
			}
		}
	})
	
	err = router.Run()
	if err != nil {
		log.Println(err)
	}

	//r = CollectRoute(r)
}