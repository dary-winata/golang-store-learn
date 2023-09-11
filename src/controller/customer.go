package controller

import (
	"context"
	"online_shop_api/src/model"
	"online_shop_api/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var databaseCustomer *mongo.Collection

func init() {
	databaseCustomer = utils.Controller.Collection("customers")
}

func CreateCustomer(customer model.Customer) {
	databaseCustomer.InsertOne(context.TODO(), customer)
}

func FindCustomerByUsername(customer model.Customer) model.Customer {
	var value model.Customer
	databaseCustomer.FindOne(context.TODO(), bson.M{"username": customer.Username}).Decode(&value)

	return value
}

func FindCustomerByEmail(customer model.Customer) model.Customer {
	var value model.Customer
	databaseCustomer.FindOne(context.TODO(), bson.M{"email": customer.Email}).Decode(&value)

	return value
}
