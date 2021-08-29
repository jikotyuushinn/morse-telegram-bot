package api

import (
	"bot"
	"encoding/json"
	tele "gopkg.in/tucnak/telebot.v3"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	bot.InitConfig()
	b, err := bot.NewBot()
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
