package controller

import (
	"context"
	"online_shop_api/src/model"
	"online_shop_api/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var databaseSeller *mongo.Collection

func init() {
	databaseSeller = utils.Controller.Collection("sellers")
}

func CreateSeller(seller model.Seller) {
	databaseSeller.InsertOne(context.TODO(), seller)
}

func FindSellerByUsernameandEmail(seller model.Seller) model.Seller {
	var value model.Seller
	databaseSeller.FindOne(context.TODO(), bson.M{"username": seller.Username, "email": seller.Email}).Decode(&value)

	return value
}

func FindSellerByUsername(seller model.Seller) model.Seller {
	var value model.Seller
	databaseSeller.FindOne(context.TODO(), bson.M{"username": seller.Username}).Decode(&value)

	return value
}

func FindSellerByEmail(seller model.Seller) model.Seller {
	var value model.Seller
	databaseSeller.FindOne(context.TODO(), bson.M{"email": seller.Email}).Decode(&value)

	return value
}
