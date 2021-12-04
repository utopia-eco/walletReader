package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/utopia-eco/walletReader/views"
)

var router *gin.Engine

func InitRoutes() {
	router = gin.Default()
	router.POST("/getWallet", views.GetWalletTokenValues)
	_ = router.Run(":8000")
}
