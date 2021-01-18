package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"morse-telegram-bot/util"
)

var b *tb.Bot

func main() {
	b = util.InitBot()
	
	b.Handle("/start", StartHandler)
	b.Handle("/help", HelpHandler)
	b.Handle("/decode", DecodeHandler)
	b.Handle("/encode", EncodeHandler)
	b.Handle(tb.OnText, OnTextHandler)
	
	b.Start()
}