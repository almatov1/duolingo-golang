package api

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"crypto/sha512"

	"github.com/golang-jwt/jwt/v5"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		login := r.FormValue("login")
		pass := r.FormValue("password")

		h := sha512.Sum512([]byte(pass))
		passHash := hex.EncodeToString(h[:])

		if login != os.Getenv("ADMIN_LOGIN") || passHash != os.Getenv("ADMIN_PASSWORD") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(48 * time.Hour).Unix(),
			"sub": login,
		})
		secret := []byte(os.Getenv("JWT_SECRET"))
		tokenString, err := token.SignedString(secret)
		if err != nil {
			http.Error(w, "token signing error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"accessToken": tokenString})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
