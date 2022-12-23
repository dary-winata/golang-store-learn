package model

import (
	"encoding/json"
	"io"

	"online_shop_api/src/utils"

	"github.com/google/uuid"
)

type Customer struct {
	Id       string `json:"id", binding:"required"`
	Username string `json:"username", binding:"required"`
	Email    string `json:"email", binding:"required"`
	Password string `json:"password", binding:"requiired"`
	Name     string `json:"name"`
}

func (customer Customer) CustomerToJson(body io.ReadCloser) Customer {
	var value Customer

	err := json.NewDecoder(body).Decode(&value)
	utils.ErrorException(err)
	value.Id = uuid.New().String()

	return value
}
