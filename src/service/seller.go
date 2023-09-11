package service

import (
	"online_shop_api/src/controller"
	"online_shop_api/src/model"
)

func RegisterSeller(seller model.Seller) {
	var value = controller.FindSellerByUsernameandEmail(seller)
	if (value == model.Seller{}) {
		controller.CreateSeller(seller)
	}
}
