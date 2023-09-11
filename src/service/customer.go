package service

import (
	"log"
	"online_shop_api/src/controller"
	"online_shop_api/src/model"
	"online_shop_api/src/utils/auth"
)

func RegisterCustomer(customer model.Customer) {
	var value = controller.FindCustomerByUsername(customer)
	if (value == model.Customer{}) {
		value = controller.FindCustomerByEmail(value)
		if (value == model.Customer{}) {
			controller.CreateCustomer(customer)
		}
	}
}

func LoginCustomer(customer model.Customer) {
	var value = controller.FindCustomerByUsername(customer)
	if (value != model.Customer{}) {
		if value.Password == customer.Password {
			token, _ := auth.GenerateJwt(value)
			log.Print(token)
		}
	}
}
