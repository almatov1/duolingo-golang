package main

import (
	"duolingo-golang/internal/bot"
	"duolingo-golang/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// MongoDB
	database.ConnectMongoDB()

	// Telebot
	tgBot, err := bot.InitBot()
	if err != nil {
		log.Fatalf("Ошибка при запуске бота: %v", err)
	}
	bot.RegisterHandlers(tgBot)
	tgBot.Start()
}
