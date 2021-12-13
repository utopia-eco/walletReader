package handlers

import (
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/services"
)

func GetWalletTokenValues(req *models.WalletTokensValueReq) (*models.WalletTokensValueResp, error) {
	var err error
	var walletValue float64
	var tokenValues []*models.TokenValue

	for _, token := range req.Tokens {
		var tokenSymbol string
		var ok bool
		value, _ := services.GetTokenValue(token.TokenAddress)
		totalValue := value * token.Amount
		if tokenSymbol, ok = consts.TokenSymbolMap[token.TokenAddress]; !ok {
			tokenSymbol = services.GetTokenSymbol(token.TokenAddress)
		}
		tokenValue := &models.TokenValue{
			TokenAddress: token.TokenAddress,
			TokenSymbol:  tokenSymbol,
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
