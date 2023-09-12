package routes

import (
	"encoding/json"
	"net/http"
	"online_shop_api/src/model"
	"online_shop_api/src/service"
)

func SellerRoute() {
	go http.HandleFunc("/v1/seller", createSeller)
	go http.HandleFunc("/v1/seller/login", loginSeller)
}

func createSeller(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var seller model.Seller
		var value = seller.ModelToJson(req.Body)
		error := service.RegisterSeller(value)

		var response model.WebResponse
		if error != nil {
			response = model.WebResponse{
				Code:   400,
				Status: error.Error(),
			}
		} else {
			response = model.WebResponse{
				Code:   200,
				Status: "Seller is sucesfully created",
			}
		}

		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(response)
	}
}

func loginSeller(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var user model.Seller
		var seller = user.ModelToJson(req.Body)
		token, error := service.LoginSeller(seller)

		res.Header().Add("Content-Type", "application/json")

		if error != nil {
			var response = model.WebResponse{
				Code:   400,
				Status: error.Error(),
			}

			json.NewEncoder(res).Encode(response)
		} else {
			var response = model.WebResponse{
				Code:   200,
				Status: "Login Successfully",
				Data: model.UserLogin{
					Token: token,
				},
			}

			json.NewEncoder(res).Encode(response)
		}
	}
}
