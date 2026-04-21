package bot

import (
	handlers "duolingo-golang/internal/bot/handlers"

	"gopkg.in/telebot.v4"
)

func RegisterHandlers(b *telebot.Bot) {
	b.Use(UserMiddleware)

	handlers.RegisterStartHandler(b)
	handlers.RegisterRegistrationHandler(b)
	handlers.RegisterTestHandler(b)
	handlers.RegisterFormatHandler(b)
}
