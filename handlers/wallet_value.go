package handlers

import (
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/services"
	"github.com/utopia-eco/walletReader/utils"
)

func GetWalletTokenValues(req *models.WalletTokensValueReq) (*models.WalletTokensValueResp, error) {
	var err error
	var walletValue float64
	var tokenValues []*models.TokenValue

	for _, token := range req.Tokens {
		value, _ := GetTokenValue(token.TokenAddress)
		totalValue := value * token.Amount
		tokenValue := &models.TokenValue{
			TokenAddress: token.TokenAddress,
			UnitPrice:    value,
			TotalValue:   totalValue,
		}
		tokenValues = append(tokenValues, tokenValue)
		walletValue += totalValue
	}

	resp := &models.WalletTokensValueResp{
		WalletAddress: req.WalletAddress,
		WalletValue:   walletValue,
		TokenValues:   tokenValues,
	}

	return resp, err
}

func GetTokenValue(tokenAddr string) (float64, error) {
	if chainlinkPriceAddr, ok := consts.TokenChainlinkProxyMap[tokenAddr]; ok {
		return services.GetKnownTokenValue(chainlinkPriceAddr)
	} else {
		//value, err = GetUnknownTokenValue(common.HexToAddress(tokenAddr), conn)
		utils.Logger.Warn("token not in known list: %s", tokenAddr)
		return 0, nil
	}
}
