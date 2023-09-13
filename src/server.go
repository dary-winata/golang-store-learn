package src

import (
	"log"
	"net/http"
	"os"

	"online_shop_api/src/routes"
	"online_shop_api/src/utils"
)

func Run() {
	port := os.Getenv("PORT")

	log.Print("Portnya adalah : " + port)

	log.Print("The is Server Running on localhost port " + port)
	routes.CustomerRoute()
	routes.SellerRoute()
	err := http.ListenAndServe(":"+port, nil)
	utils.ErrorException(err)
}
