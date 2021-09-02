package api

import (
	"bot/internal"
	"encoding/json"
	tele "gopkg.in/tucnak/telebot.v3"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	b, err := internal.NewBot()
	if err != nil {
		panic(err)
	}
	b.Start()

	var u tele.Update
	resp, _ := io.ReadAll(r.Body)
	if err = json.Unmarshal(resp, &u); err == nil {
		b.ProcessUpdate(u)
	}
	return
}
