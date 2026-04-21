package bot

import "gopkg.in/telebot.v4"

func RegisterHandlers(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		return c.Send("Привет! Я модульный бот на Go 🚀")
	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("Echo: " + c.Text())
	})
}
