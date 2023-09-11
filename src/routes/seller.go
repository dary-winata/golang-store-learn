package routes

import (
	"net/http"
	"online_shop_api/src/model"
	"online_shop_api/src/service"
)

func SellerRoute() {
	go http.HandleFunc("/v1/seller", createSeller)
}

func createSeller(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var seller model.Seller
		var value = seller.ModelToJson(req.Body)
		service.RegisterSeller(value)
	}
}
