package handlers

import (
	"gopkg.in/telebot.v4"
)

func RegisterStartHandler(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		userValue := c.Get("user")

		if userValue == nil {
			return selectLanguage(c)
		}

		return nil
	})
}
