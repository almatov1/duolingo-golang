package database

import (
	"context"
	"duolingo-golang/internal/configs"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	uri := os.Getenv("MONGO_URI")

	var err error
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", user, password, uri)
	clientOptions := options.Client().ApplyURI(mongoURI)
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	collection := Client.Database(configs.DBName).Collection(configs.UserCollectionName)
	_, err = collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.M{"telegramId": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Ошибка пинга MongoDB: %v", err)
	}
	log.Println("Успешно подключено к MongoDB")
}
