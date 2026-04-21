package main

import (
	middleware "duolingo-golang/internal/api"
	api "duolingo-golang/internal/api/handlers"
	"duolingo-golang/internal/database"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// MongoDB
	database.ConnectMongoDB()

	http.HandleFunc("/auth", api.AuthHandler)
	http.HandleFunc("/users", middleware.Auth(api.UsersHandler))

	http.ListenAndServe(":8080", nil)
}
