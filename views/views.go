package views

import (
	"github.com/gin-gonic/gin"
	"github.com/utopia-eco/walletReader/handlers"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/utils"
)

func GetWalletTokenValues(c *gin.Context) {
	req := &models.WalletTokensValueReq{}
	err := c.BindJSON(req)
	if err != nil {
		utils.Logger.Error("GetWalletTokenValues unbind err: %v", err)
	}

	handlers.GetWalletTokenValues()

}
