package main

import (
	"github.com/utopia-eco/walletReader/routes"
	"github.com/utopia-eco/walletReader/services"
	"github.com/utopia-eco/walletReader/utils"
)

func main() {
	utils.InitLogger()
	services.InitBsc()
	services.InitPancakes()

	routes.InitRoutes()
}
