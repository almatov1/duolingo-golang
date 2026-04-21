package bot

import (
	"context"
	"duolingo-golang/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type State string

const (
	StateNone         State = ""
	StateWaitName     State = "wait_name"
	StateWaitBirthday State = "wait_birthday"
	StateWaitLevel    State = "wait_level"
)

// SetState сохраняет текущий шаг пользователя в отдельную коллекцию 'states'
func SetState(tgID int64, state State) {
	col := database.Client.Database("duolingo").Collection("states")
	_, _ = col.UpdateOne(
		context.TODO(),
		bson.M{"tgId": tgID},
		bson.M{"$set": bson.M{"state": string(state)}},
		options.Update().SetUpsert(true),
	)
}

// GetState получает текущий шаг
func GetState(tgID int64) State {
	col := database.Client.Database("duolingo").Collection("states")
	var res struct {
		State string `bson:"state"`
	}
	err := col.FindOne(context.TODO(), bson.M{"tgId": tgID}).Decode(&res)
	if err != nil {
		return StateNone
	}
	return State(res.State)
}
