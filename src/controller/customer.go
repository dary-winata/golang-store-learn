package controller

import (
	"context"
	"online_shop_api/src/model"
	"online_shop_api/src/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Collection

func init() {
	database = utils.Controller.Collection("customer")
}

func CreateCustomer(customer model.Customer) {
	database.InsertOne(context.TODO(), customer)
}
