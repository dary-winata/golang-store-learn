package model

import (
	"encoding/json"
	"io"
	"online_shop_api/src/utils"

	"github.com/google/uuid"
)

type Seller struct {
	Id       string `json:"id", binding:"required"`
	Username string `json:"username", binding:"required"`
	Email    string `json:"email", binding:"required"`
	Password string `json:"password", binding:"requiired"`
	Name     string `json:"name"`
}

func (seller Seller) ModelToJson(body io.ReadCloser) Seller {
	var value Seller

	err := json.NewDecoder(body).Decode(&value)
	utils.ErrorException(err)
	value.Id = uuid.New().String()

	return value
}

func (seller Seller) GetUsername() string {
	return seller.Username
}
