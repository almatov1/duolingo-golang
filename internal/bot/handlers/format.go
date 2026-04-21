package handlers

import (
	"context"
	"duolingo-golang/internal/configs"
	"duolingo-golang/internal/database"
	"duolingo-golang/internal/i18n"
	"duolingo-golang/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/telebot.v4"
)

func selectFormat(c telebot.Context) error {
	user := c.Get("user").(*models.User)

	formatSelector := &telebot.ReplyMarkup{}
	btnOnline := formatSelector.Data(
		i18n.T(string(user.Language), "format.online"),
		"format",
		string(models.FormatOnline),
	)
	btnOffline := formatSelector.Data(
		i18n.T(string(user.Language), "format.offline"),
		"format",
		string(models.FormatOffline),
	)

	formatSelector.Inline(
		formatSelector.Row(btnOnline, btnOffline),
	)

	return c.Send(i18n.T(string(user.Language), "format.start"), formatSelector)
}

func RegisterFormatHandler(b *telebot.Bot) {
	b.Handle(&telebot.Btn{Unique: "format"}, handleFormatSelect)
}

func handleFormatSelect(c telebot.Context) error {
	user, ok := c.Get("user").(*models.User)
	if !ok || user.State != models.StateFormat {
		return nil
	}

	format := c.Callback().Data

	collection := database.Client.Database(configs.DBName).Collection(configs.UserCollectionName)
	_, _ = collection.UpdateOne(
		context.TODO(),
		bson.M{"telegramId": user.TelegramID},
		bson.M{"$set": bson.M{"format": format}},
	)

	if format == string(models.FormatOffline) {
		return c.Send(i18n.T(string(user.Language), "format.location"))
	}

	return c.Send("lessons")
}
