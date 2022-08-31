package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	UserId      string
	Description string
}

func (t *Task) AddTask() {
    executeOnCollection("tasks", func(c *mongo.Collection, ctx context.Context) {
        c.InsertOne(ctx, t)
    })
}
