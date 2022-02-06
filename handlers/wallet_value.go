package handlers

import (
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/services"
	"strings"
)

func GetWalletTokenValues(req *models.WalletTokensValueReq) (*models.WalletTokensValueResp, error) {
	var err error
	var walletValue float64
	var tokenValues []*models.TokenValue

	for _, token := range req.Tokens {
		if _, ok := consts.BlacklistTokens[strings.ToLower(token.TokenAddress)]; ok {
			tokenValue := &models.TokenValue{
				TokenAddress: token.TokenAddress,
				Blacklist:    true,
				TokenSymbol:  "",
				Amount:       token.Amount,
				UnitPrice:    0,
				TotalValue:   0,
			}
			tokenValues = append(tokenValues, tokenValue)
			continue
		}
		var tokenSymbol string
		var ok bool
		value, _ := services.GetTokenValueFromApi(token.TokenAddress)
		if value == 0 {
			value, _ = services.GetTokenValue(token.TokenAddress)
		}
		totalValue := value * token.Amount
		if tokenSymbol, ok = consts.TokenSymbolMap[token.TokenAddress]; !ok {
			tokenSymbol = services.GetTokenSymbol(token.TokenAddress)
		}
		tokenValue := &models.TokenValue{
			TokenAddress: token.TokenAddress,
			Blacklist:    false,
			TokenSymbol:  tokenSymbol,
			Amount:       token.Amount,
			UnitPrice:    value,
			TotalValue:   totalValue,
		}
		if token.IsWrapped != nil {
			if *token.IsWrapped {
				tokenValue.TokenSymbol = "W" + tokenValue.TokenSymbol
			}
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
