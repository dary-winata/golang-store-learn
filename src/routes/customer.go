package routes

import (
	"net/http"
	"online_shop_api/src/model"
	"online_shop_api/src/service"
)

func CustomerRoute() {
	go http.HandleFunc("/v1/customer", createCustomer)
	go http.HandleFunc("/v1/customer/login", loginCustomer)
}

func createCustomer(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var customer model.Customer
		var value = customer.ModelToJson(req.Body)
		service.RegisterCustomer(value)
	}
}

func loginCustomer(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var user model.Customer
		var value = user.ModelToJson(req.Body)
		service.LoginCustomer(value)
	}
}