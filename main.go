package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"morse-telegram-bot/controller"
	"morse-telegram-bot/util"
)

func main() {
	
	webhook := &tb.Webhook{
		Listen: ":" + util.Port,
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: util.WebhookHost,
		},
	}
	
	b, err := tb.NewBot(tb.Settings{
		Token:  util.AccessToken,
		Poller: webhook,
	})
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/start", func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "不準開始。")
	})
	
	b.Handle("/help", func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "禁止幫助⛔。")
	})

	b.Handle("/decode", func(m *tb.Message) {
		if m.Payload == "" {
			_, _ = b.Send(m.Chat, "勸你最好有輸入。")
		}
		text, _ := controller.JsParser(util.StaticPath, "xmorse.decode", m.Payload)
		_, _ = b.Send(m.Chat, text)
	})
	
	b.Handle("/encode", func(m *tb.Message) {
		if m.Payload == "" {
			_, _ = b.Send(m.Chat, "勸你最好有輸入。")
		}
		text, _ := controller.JsParser(util.StaticPath, "xmorse.encode", m.Payload)
		_, _ = b.Send(m.Chat, text)
		//b.Delete(m.Text)
	})
	//b.Handle(tb.OnText, func(m *tb.Message) {
	//	_, _ = b.Send(m.Sender, "hello world")
	//})
	
	b.Start()
	
}