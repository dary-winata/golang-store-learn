package service

import (
	"online_shop_api/src/model"
)

func CreateCustomer(customer model.Customer) {
	channel := make(chan model.Customer)
	defer close(channel)
	channel <- customer
	
}
