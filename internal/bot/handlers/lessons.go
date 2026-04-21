package handlers

import "gopkg.in/telebot.v4"

func showLessons(c telebot.Context) error {
	return c.Send("10 тақырып ұсынылады")
}
