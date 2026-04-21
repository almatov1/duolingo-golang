package handlers

import (
	"context"
	"duolingo-golang/internal/configs"
	"duolingo-golang/internal/database"
	"duolingo-golang/internal/i18n"
	"duolingo-golang/internal/models"
	"time"

	"gopkg.in/telebot.v4"
)

var (
	userDrafts = make(map[int64]*models.User)

	languageSelector = &telebot.ReplyMarkup{}
	btnKk            = languageSelector.Data("🇰🇿 Қазақ тілі", "lang", "kk")
	btnRu            = languageSelector.Data("🇷🇺 Русский язык", "lang", "ru")
	btnEn            = languageSelector.Data("🇺🇸 English", "lang", "en")
)

func selectLanguage(c telebot.Context) error {
	languageSelector.Inline(
		languageSelector.Row(btnKk, btnRu, btnEn),
	)

	return c.Send("👋 Ақтөбе облысының Тілдерді дамыту басқармасы ММ-нің «Тілдерді оқыту орталығы» КММ әзірлеген тіл үйрену ботына қош келдіңіз!\n\nҚызметті бастау үшін өзіңізге ыңғайлы тілді таңдаңыз:", languageSelector)
}

func RegisterRegistrationHandler(b *telebot.Bot) {
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		userID := c.Sender().ID

		draft := userDrafts[userID]
		if draft == nil {
			return nil
		}

		if draft.Name == "" {
			return handleRegistrationName(c, draft)
		}

		if draft.Birthday == "" {
			return handleRegistrationBirthday(c, draft)
		}

		if draft.Nationality == "" {
			return handleRegistrationNationality(c, draft)
		}

		if draft.Workplace == "" {
			return handleRegistrationWorkplace(c, draft)
		}

		if draft.Address == "" {
			return handleRegistrationAddress(c, draft)
		}

		if draft.Telephone == "" {
			return handleRegistrationTelephone(c, draft)
		}

		return nil
	})

	b.Handle(&btnKk, handleLanguageSelect)
	b.Handle(&btnRu, handleLanguageSelect)
	b.Handle(&btnEn, handleLanguageSelect)
}

func handleLanguageSelect(c telebot.Context) error {
	userValue := c.Get("user")
	if userValue != nil {
		return nil
	}

	userID := c.Sender().ID
	languageCode := c.Callback().Data

	draft := &models.User{
		TelegramID: userID,
		Language:   models.Language(languageCode),
	}
	userDrafts[userID] = draft

	return c.Send(
		i18n.T(languageCode, "registration", "start") +
			"\n\n" +
			i18n.T(languageCode, "registration", "ask_name"),
	)
}

func handleRegistrationName(c telebot.Context, user *models.User) error {
	draft := userDrafts[user.TelegramID]
	draft.Name = c.Text()

	return c.Send(i18n.T(string(draft.Language), "registration", "ask_birthday"))
}

func handleRegistrationBirthday(c telebot.Context, user *models.User) error {
	draft := userDrafts[user.TelegramID]
	draft.Birthday = c.Text()

	return c.Send(i18n.T(string(draft.Language), "registration", "ask_nationality"))
}

func handleRegistrationNationality(c telebot.Context, user *models.User) error {
	draft := userDrafts[user.TelegramID]
	draft.Nationality = c.Text()

	return c.Send(i18n.T(string(draft.Language), "registration", "ask_workplace"))
}

func handleRegistrationWorkplace(c telebot.Context, user *models.User) error {
	draft := userDrafts[user.TelegramID]
	draft.Workplace = c.Text()

	return c.Send(i18n.T(string(draft.Language), "registration", "ask_address"))
}

func handleRegistrationAddress(c telebot.Context, user *models.User) error {
	draft := userDrafts[user.TelegramID]
	draft.Address = c.Text()

	return c.Send(i18n.T(string(draft.Language), "registration", "ask_telephone"))
}

func handleRegistrationTelephone(c telebot.Context, user *models.User) error {
	draft := userDrafts[user.TelegramID]
	draft.Telephone = c.Text()

	collection := database.Client.Database(configs.DBName).Collection(configs.UserCollectionName)

	draft.CreatedAt = time.Now()
	_, err := collection.InsertOne(
		context.TODO(),
		draft,
	)
	if err != nil {
		return nil
	}

	delete(userDrafts, user.TelegramID)

	return c.Send("Successfully registered")
}
