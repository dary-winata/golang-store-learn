package routes

import (
	"net/http"
	"online_shop_api/src/model"
)

func CreateCustomer(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var customer model.Customer
		customer.CustomerToJson(req.Body)
	}
}

func CustomerRoute() {
	go http.HandleFunc("/v1/customer", CreateCustomer)
}