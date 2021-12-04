package views

import (
	"github.com/gin-gonic/gin"
	"github.com/utopia-eco/walletReader/handlers"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/utils"
)

func GetWalletTokenValues(c *gin.Context) {
	req := &models.WalletTokensValueReq{}
	println("req")
	println(req)
	utils.Logger.Info("[views] GetWalletTokenValues req: %v", req)

	err := c.BindJSON(req)
	if err != nil {
		utils.Logger.Error("[views] GetWalletTokenValues unbind err: %v", err)
	}

	resp, err := handlers.GetWalletTokenValues(req)
	if err != nil {
		c.JSON(500, "error")
	}
	utils.Logger.Info("[views] GetWalletTokenValues resp: %v", resp)
	c.JSON(200, resp)
}
