package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Controller *mongo.Database
var ctx = context.Background()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://darywinata:zabuza0920@oneforall.ar0ve.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		ErrorException(err)
	}
	fmt.Println("database mongo connected")

	Controller = client.Database("go-olshop")
}
