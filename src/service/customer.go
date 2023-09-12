package service

import (
	"errors"
	"online_shop_api/src/controller"
	"online_shop_api/src/model"
	"online_shop_api/src/utils/auth"
)

func RegisterCustomer(customer model.Customer) error {
	var value = controller.FindCustomerByUsername(customer)

	var err error

	if (value == model.Customer{}) {
		var sellerByEmail = controller.FindCustomerByEmail(customer)
		if (sellerByEmail == model.Customer{}) {
			controller.CreateCustomer(customer)

			return nil
		} else {
			err = errors.New("Email is already used")
		}
	} else {
		err = errors.New("Username is already taken")
	}

	return err
}

func LoginCustomer(customer model.Customer) (string, error) {
	var value = controller.FindCustomerByUsername(customer)
	if (value != model.Customer{}) {
		if value.Password == customer.Password {
			token, error := auth.GenerateJwt(value)

			if error != nil {
				return "", error
			}

			return token, nil
		}
	}

	err := errors.New("Username or Password is not correct")

	return "", err
}
