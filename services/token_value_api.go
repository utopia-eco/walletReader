package services

import (
	"encoding/json"
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/utils"
	"strconv"
	"time"
)

// GetTokenValueFromApi checks if token is updated in last 30s
func GetTokenValueFromApi(tokenAddr string) (float64, error) {
	if lastTokenPrice, ok := TokenPriceMap[tokenAddr]; ok {
		if time.Now().UTC().Unix()-lastTokenPrice.Timestamp < UpdateThreshold {
			return lastTokenPrice.Price, nil
		}
	}
	return GetTokenValueFromPancakeswap(tokenAddr)
}

// GetTokenValueFromPancakeswap checks pancake price api for price
func GetTokenValueFromPancakeswap(tokenAddr string) (float64, error) {
	var err error

	url := consts.PancakeswapApiUrl + "/v2/tokens/" + tokenAddr
	pancakePriceHttpBody, err := SendHttpGetRequest(url)
	if err != nil {
		return 0, err
	}

	pancakePriceResp := &models.PancakePriceResp{}
	if err := json.Unmarshal(*pancakePriceHttpBody, pancakePriceResp); err != nil {
		utils.Logger.Error("GetTokenValueFromPancakeswap unmarshal token: %s, err: %v", tokenAddr, err)
		return 0, err
	}

	var price float64
	price, err = strconv.ParseFloat(pancakePriceResp.Data.Price, 64)
	if err != nil {
		utils.Logger.Error("GetTokenValueFromPancakeswap price to float: %s, err: %v",
			pancakePriceResp.Data.Price, err)
		return 0, err
	}
	if price <= 0 {
		priceBnb, err := strconv.ParseFloat(pancakePriceResp.Data.PriceBNB, 64)
		if err != nil {
			utils.Logger.Error("GetTokenValueFromPancakeswap priceBnb to float: %s, err: %v",
				pancakePriceResp.Data.Price, err)
			return 0, err
		}
		if priceBnb <= 0 {
			return 0, nil
		}
		bnbPrice, err := GetTokenValue(consts.BnbAddr)
		if err != nil {
			utils.Logger.Error("get pair value err: %+v", err)
			return 0, nil
		}
		price = priceBnb * bnbPrice
	}

	TokenPriceMap[tokenAddr] = &models.LastTokenPrice{
		Price:     price,
		Timestamp: time.Now().UTC().Unix(),
	}

	return price, nil
}
