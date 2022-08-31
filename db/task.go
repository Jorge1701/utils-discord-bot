package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	UserId      string
	Description string
}

func (t *Task) SaveTask() {
	executeOnCollection("tasks", func(c *mongo.Collection, ctx context.Context) {
		c.InsertOne(ctx, t)
	})
}

func ListTasks(userId string) []Task {
	var tasks []Task
	executeOnCollection("tasks", func(c *mongo.Collection, ctx context.Context) {
		cursor, err := c.Find(ctx, bson.D{primitive.E{Key: "userid", Value: userId}})
		if err != nil {
			fmt.Println("Error when listing tasks", err)
		}

		if err = cursor.All(ctx, &tasks); err != nil {
			fmt.Println("Error getting tasks result", err)
		}
	})
	return tasks
}
