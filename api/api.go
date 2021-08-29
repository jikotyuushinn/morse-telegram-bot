package api

import (
	"encoding/json"
	tele "gopkg.in/tucnak/telebot.v3"
	"io"
	tgbot "morse-telegram-bot"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tgbot.InitConfig()
	b, err := tgbot.NewBot()
	if err != nil {
		panic(err)
	}

	var u tele.Update
	resp, _ := io.ReadAll(r.Body)
	if err = json.Unmarshal(resp, &u); err == nil {
		b.ProcessUpdate(u)
	}
	return
}
