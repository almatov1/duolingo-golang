package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"duolingo-golang/internal/configs"
	"duolingo-golang/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 20
		}

		collection := database.Client.Database(configs.DBName).Collection(configs.UserCollectionName)
		opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
		cur, err := collection.Find(context.TODO(), bson.M{}, opts)
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
		defer cur.Close(context.TODO())

		var users []bson.M
		if err = cur.All(context.TODO(), &users); err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		collection := database.Client.Database(configs.DBName).Collection(configs.UserCollectionName)
		_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("user deleted"))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
