package utils

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Controller *mongo.Database
var ctx = context.Background()

func init() {
	mongodbUrl := os.Getenv("MONGODB_URL")

	if mongodbUrl == "" {
		fmt.Println("no mongodb url")
	}

	clientOptions := options.Client().ApplyURI(mongodbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		ErrorException(err)
	}
	fmt.Println("database mongo connected")

	Controller = client.Database("go-olshop")
}
