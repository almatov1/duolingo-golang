package handlers

import (
	"context"
	"duolingo-golang/internal/configs"
	"duolingo-golang/internal/database"
	"duolingo-golang/internal/i18n"
	"duolingo-golang/internal/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/telebot.v4"
)

type (
	TestQuestion struct {
		Question     string   `bson:"question" json:"question"`
		Answers      []string `bson:"answers" json:"answers"`
		CorrectIndex int      `bson:"correctIndex" json:"correctIndex"`
	}

	TestSession struct {
		CurrentIdx int
		Score      int
		MessageID  int
	}
)

var (
	testSessions = make(map[int64]*TestSession)

	TEST_QUESTIONS = []TestQuestion{
		{
			Question:     "What is 2+2?",
			Answers:      []string{"3", "4", "5", "6"},
			CorrectIndex: 1,
		},
		{
			Question:     "Capital of Kazakhstan?",
			Answers:      []string{"Astana", "Almaty", "Shymkent", "Aqtobe"},
			CorrectIndex: 0,
		},
	}

	TEST_RESULT = []struct {
		CorrectCount int
		Level        string
	}{
		{0, "A1"},
		{2, "A2"},
		{3, "B1"},
		{4, "B2"},
		{5, "C1"},
	}
)

func startTest(c telebot.Context) error {
	user := c.Get("user").(*models.User)
	testSessions[user.TelegramID] = &TestSession{CurrentIdx: 0, Score: 0}
	return sendQuestion(c, user.TelegramID, true)
}

func RegisterTestHandler(b *telebot.Bot) {
	b.Handle(&telebot.Btn{Unique: "answer"}, func(c telebot.Context) error {
		user := c.Get("user").(*models.User)
		session := testSessions[user.TelegramID]
		if session == nil && user.State != models.StateTest {
			return nil
		}

		data := c.Callback().Data
		var qIdx, aIdx int
		_, err := fmt.Sscanf(data, "%d-%d", &qIdx, &aIdx)
		if err != nil {
			return nil
		}

		if qIdx != session.CurrentIdx {
			return c.Respond()
		}

		question := TEST_QUESTIONS[session.CurrentIdx]
		if aIdx == question.CorrectIndex {
			session.Score++
		}

		session.CurrentIdx++
		if session.CurrentIdx < len(TEST_QUESTIONS) {
			return sendQuestion(c, user.TelegramID, false)
		}

		level := "A1"
		for _, res := range TEST_RESULT {
			if session.Score >= res.CorrectCount {
				level = res.Level
			}
		}

		collection := database.Client.Database(configs.DBName).Collection(configs.UserCollectionName)
		_, _ = collection.UpdateOne(
			context.TODO(),
			bson.M{"telegramId": user.TelegramID},
			bson.M{"$set": bson.M{
				"level": level,
				"state": models.StateFormat,
			}},
		)

		delete(testSessions, user.TelegramID)

		c.Send(i18n.T(
			string(user.Language),
			"test.result",
			level,
			i18n.T(string(user.Language), "test."+level),
		))

		return selectFormat(c)
	})
}

func sendQuestion(c telebot.Context, userID int64, isFirst bool) error {
	session := testSessions[userID]
	q := TEST_QUESTIONS[session.CurrentIdx]

	keyboard := &telebot.ReplyMarkup{}
	rows := []telebot.Row{}
	for i, ans := range q.Answers {
		btn := keyboard.Data(ans, "answer", fmt.Sprintf("%d-%d", session.CurrentIdx, i))
		rows = append(rows, keyboard.Row(btn))
	}
	keyboard.Inline(rows...)

	text := fmt.Sprintf("%d. %s", session.CurrentIdx+1, q.Question)

	if isFirst {
		return c.Send(text, keyboard)
	} else {
		if err := c.Edit("⏳"); err != nil {
			return err
		}

		time.Sleep(500 * time.Millisecond)
		return c.Edit(text, keyboard)
	}
}
