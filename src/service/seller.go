package service

import (
	"errors"
	"online_shop_api/src/controller"
	"online_shop_api/src/model"
	"online_shop_api/src/utils/auth"
)

func RegisterSeller(seller model.Seller) error {
	var error error
	var dbSellerByUsername = controller.FindSellerByUsername(seller)

	if (dbSellerByUsername == model.Seller{}) {
		var dbSellerByEmail = controller.FindSellerByEmail(seller)

		if (dbSellerByEmail == model.Seller{}) {
			controller.CreateSeller(seller)

			return nil
		} else {
			error = errors.New("email is already taken")
		}
	} else {
		error = errors.New("username is already used")
	}

	return error
}

func LoginSeller(seller model.Seller) (string, error) {
	var dbSeller = controller.FindSellerByUsernameandEmail(seller)
	if (dbSeller != model.Seller{}) {
		if dbSeller.Password == seller.Password {
			token, error := auth.GenerateJwt(seller)

			if error != nil {
				return "", error
			}

			return token, nil
		}
	}

	err := errors.New("Username or Password is incorrect")

	return "", err
}
