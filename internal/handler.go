package internal

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b Bot) onStart(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	_ = c.Notify(tele.Typing)
	return c.Send("🈲️开始")
}

func (b Bot) onHelp(c tele.Context) error {
	_ = c.Notify(tele.Typing)
	return c.Send("🈲️帮助")
}

func (b Bot) onDecode(c tele.Context) error {
	_ = c.Notify(tele.Typing)

	if c.Message().Payload == "" {
		return c.Send("勸你最好有輸入。")
	}

	text, err := jsParser(StaticPath, "xmorse.decode", c.Message().Payload)
	if err != nil {
		return err
	}
	return c.Reply(text)
}

func (b Bot) onEncode(c tele.Context) error {
	_ = c.Notify(tele.Typing)

	if c.Message().Payload == "" {
		return c.Send("勸你最好有輸入。")
	}

	morseCode, err := jsParser(StaticPath, "xmorse.encode", c.Message().Payload)
	if err != nil {
		return err
	}
	return c.Reply(morseCode)
}

func (b Bot) onText(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	_ = c.Notify(tele.Typing)
	return c.Reply("這位先生，本小姐不陪聊哦。")
}
