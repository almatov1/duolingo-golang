package handlers

import (
	"duolingo-golang/internal/models"

	"gopkg.in/telebot.v4"
)

func RegisterStartHandler(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		user, ok := c.Get("user").(*models.User)

		if !ok || user == nil {
			return selectLanguage(c)
		}

		if user.Level == nil {
			return startTest(c)
		}

		if user.Format == nil {
			return selectLanguage(c)
		}

		if *user.Format == models.FormatOffline {
			return handleOfflineFormat(c, *user)
		}

		if *user.Format == models.FormatOnline {
			return showLessons(c)
		}

		return nil
	})
}
