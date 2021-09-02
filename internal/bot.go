package internal

import (
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/middleware"
)

type Bot struct {
	*tele.Bot
}

func NewBot() (*Bot, error) {
	initConfig()
	b, err := tele.NewBot(tele.Settings{Token: AccessToken, Synchronous: true})
	if err != nil {
		return nil, err
	}

	return &Bot{Bot: b}, nil
}

func (b *Bot) Start() {
	b.Use(middleware.DefaultLogger(), middleware.AutoRespond())
	b.Handle("/start", b.onStart)
	b.Handle("/help", b.onHelp)
	b.Handle("/decode", b.onDecode)
	b.Handle("/encode", b.onEncode)
	b.Handle(tele.OnText, b.onText)
}
