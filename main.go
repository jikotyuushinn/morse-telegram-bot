package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
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
	
	// to monitor changes run: heroku logs --tail
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
	
	
	// gin router
	router := gin.New()
	router.Use(gin.Logger())
	
	router.POST("/" + bot.Token, webhookHandler)
	
	err = router.Run()
	if err != nil {
		log.Println(err)
	}

	//r := gin.Default()
	//r.Use(LogMiddleware())
	//r = CollectRoute(r)
	//panic(r.Run())
}