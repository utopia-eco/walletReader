package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/utopia-eco/walletReader/views"
)

var router *gin.Engine

func InitRoutes() {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"https://localhost:3000",
		"https://utopia.cc/",
	}
	router.Use(cors.New(config))
	router.POST("/getWallet", views.GetWalletTokenValues)
	_ = router.Run()
}
