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

	// HTTP
	http.HandleFunc("/auth", withCORS(api.AuthHandler))
	http.HandleFunc("/users", withCORS(middleware.Auth(api.UsersHandler)))

	http.ListenAndServe(":8080", nil)
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
