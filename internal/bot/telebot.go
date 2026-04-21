package bot

import (
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

var Bot *telebot.Bot

func InitBot() (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}

	Bot = b
	return b, nil
}
