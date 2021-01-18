package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"morse-telegram-bot/util"
)

func StartHandler(m *tb.Message) {
	_, _ = b.Send(m.Chat, "不準開始。")
}

func HelpHandler(m *tb.Message) {
	_, _ = b.Send(m.Chat, "禁止幫助⛔。")
}

func DecodeHandler(m *tb.Message) {
	if m.Payload == "" {
		_, _ = b.Send(m.Chat, "勸你最好有輸入。")
	}
	text, _ := util.JsParser(util.StaticPath, "xmorse.decode", m.Payload)
	_, _ = b.Reply(m, text)
	_ = b.Delete(m)
}

func EncodeHandler(m *tb.Message) {
	if m.Payload == "" {
		_, _ = b.Send(m.Chat, "勸你最好有輸入。")
	}
	morseCode, _ := util.JsParser(util.StaticPath, "xmorse.encode", m.Payload)
	_, _ = b.Reply(m, morseCode)
	_ = b.Delete(m)
}

func OnTextHandler(m *tb.Message) {
	_, _ = b.Send(m.Chat, "這位先生，本小姐不陪聊哦。")
}