package db

import (
	"context"
	"discord-bot/config"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var configuration config.Configuration

type dataBaseAction func(*mongo.Client, context.Context)
type CollectionAction func(*mongo.Collection, context.Context)

func init() {
	configuration = config.GetConfiguration()
}

func executeOnDB(action dataBaseAction) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configuration.MongoURI))

	if err != nil {
		fmt.Println("Error connecting with database", err)
	} else {
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				fmt.Println("Error disconnectig with database", err)
			}
		}()

		action(client, ctx)
	}
}

func executeOnCollection(collection string, action CollectionAction) {
	executeOnDB(func(c *mongo.Client, ctx context.Context) {
		collection := c.Database(configuration.DataBase).Collection(collection)
		action(collection, ctx)
	})
}
