package src

import (
	"log"
	"net/http"

	"online_shop_api/src/utils"
)

func Run() {
	log.Print("The is Server Running on localhost port 3000")
	err := http.ListenAndServe(":3000", nil)
	
	utils.ErrorException(err)
}
