package auth

import (
	"fmt"
	"time"

	"online_shop_api/src/model"

	"github.com/golang-jwt/jwt"
)

var SECRET = []byte("XApHbc93o8")

func GenerateJwt(user model.User) (string, error) {
	//create token jwt
	token := jwt.New(jwt.SigningMethodHS256)

	//set token jwt
	claims := token.Claims.(jwt.MapClaims)
	// log.Print(time.Now().Add(24 * time.Hour).Unix())
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["authorized"] = "customer"
	claims["user"] = user.GetUsername()

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	return tokenStr, nil
}

// func ValidateJwt(token string) error {

// }
