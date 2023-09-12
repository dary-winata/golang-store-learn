package routes

import (
	"encoding/json"
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
		error := service.RegisterCustomer(value)

		res.Header().Add("Content-Type", "application/json")

		var response model.WebResponse

		if error != nil {
			response = model.WebResponse{
				Code:   400,
				Status: error.Error(),
			}
		} else {
			response = model.WebResponse{
				Code:   200,
				Status: "Customer User is Succesfully Created",
			}
		}

		json.NewEncoder(res).Encode(response)
	}
}

func loginCustomer(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var user model.Customer
		var value = user.ModelToJson(req.Body)
		tokens, errors := service.LoginCustomer(value)

		res.Header().Add("Content-Type", "application/json")

		if errors != nil {
			var response = model.WebResponse{
				Code:   400,
				Status: errors.Error(),
			}
			json.NewEncoder(res).Encode(response)
		} else {
			var response = model.WebResponse{
				Code:   200,
				Status: "Login successfully",
				Data: model.UserLogin{
					Token: tokens,
				},
			}
			json.NewEncoder(res).Encode(response)
		}
	}
}
