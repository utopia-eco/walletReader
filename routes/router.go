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
		"https://utopia-web-staging-3de7nawkiq-uw.a.run.app",
	}
	router.Use(cors.New(config))
	router.POST("/getWallet", views.GetWalletTokenValues)
	_ = router.Run()
}
