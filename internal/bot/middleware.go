package bot

import (
	"context"
	"duolingo-golang/internal/configs"
	"duolingo-golang/internal/database"
	"duolingo-golang/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/telebot.v4"
)

func UserMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		collection := database.Client.Database(configs.DBName).Collection(configs.UserCollectionName)

		var user models.User
		err := collection.FindOne(context.TODO(), bson.M{"telegramId": c.Sender().ID}).Decode(&user)

		if err == nil {
			c.Set("user", &user)
		}

		return next(c)
	}
}
