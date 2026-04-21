package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// --- ENUMS ---
type Language string
type Level string
type StudyFormat string
type LessonType string

const (
	LangKK Language = "kk"
	LangRU Language = "ru"
	LangEN Language = "en"

	LevelA1 Level = "A1"
	LevelA2 Level = "A2"
	LevelB1 Level = "B1"
	LevelB2 Level = "B2"
	LevelC1 Level = "C1"

	FormatOffline StudyFormat = "OFFLINE"
	FormatOnline  StudyFormat = "ONLINE"

	Writing   LessonType = "WRITING"
	Reading   LessonType = "READING"
	Listening LessonType = "LISTENING"
	Speaking  LessonType = "SPEAKING"
)

// --- COLLECTIONS ---

type UserProgress struct {
	TopicId    int        `bson:"topicId" json:"topicId"`
	LessonType LessonType `bson:"lessonType" json:"lessonType"`
	Content    *string    `bson:"content,omitempty" json:"content,omitempty"`
	CreatedAt  time.Time  `bson:"createdAt" json:"createdAt"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TelegramID int64              `bson:"telegramId" json:"telegramId"`

	Name        string `bson:"name" json:"name"`
	Birthday    string `bson:"birthday" json:"birthday"`
	Nationality string `bson:"nationality" json:"nationality"`
	Workplace   string `bson:"workplace" json:"workplace"`
	Address     string `bson:"address" json:"address"`
	Telephone   string `bson:"telephone" json:"telephone"`

	Language Language     `bson:"language" json:"language"`
	Level    *Level       `bson:"level,omitempty" json:"level,omitempty"`
	Format   *StudyFormat `bson:"format,omitempty" json:"format,omitempty"`

	Progress []UserProgress `bson:"progress" json:"progress"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
