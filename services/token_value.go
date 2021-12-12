package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/utopia-eco/walletReader/abi/abigen/chainlink"
	"github.com/utopia-eco/walletReader/abi/abigen/pancake"
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/models"
	"github.com/utopia-eco/walletReader/utils"
	"math/big"
	"time"
)

const (
	UpdateThreshold = 30 // 30s interval for price update
	UsdBase         = 100000000
	BnbBase         = 1000000000000000000
)

var (
	//BnbPrice      float64
	TokenPriceMap map[string]*models.LastTokenPrice
)

// GetTokenValue checks if token is updated in last 30s
func GetTokenValue(tokenAddr string) (float64, error) {
	if lastTokenPrice, ok := TokenPriceMap[tokenAddr]; ok {
		if time.Now().UTC().Unix()-lastTokenPrice.Timestamp < UpdateThreshold {
			return lastTokenPrice.Price, nil
		}
	}
	if chainlinkPriceAddr, ok := consts.TokenChainlinkProxyMap[tokenAddr]; ok {
		return GetKnownTokenValue(chainlinkPriceAddr)
	} else {
		return GetUnknownTokenValue(tokenAddr)
	}
}

// GetKnownTokenValue checks chainlink price proxy oracle for price
func GetKnownTokenValue(chainlinkPriceAddr string) (float64, error) {
	chainlinkPriceCaller, err := chainlink.NewChainlinkPrice(common.HexToAddress(chainlinkPriceAddr), BscConn)
	if err != nil {
		utils.Logger.Error("GetKnownTokenValue NewChainlinkPrice err: %v", err)
		return 0, err
	}

	value, err := chainlinkPriceCaller.LatestAnswer(nil)
	if err != nil {
		utils.Logger.Error("GetKnownTokenValue LatestAnswer err: %v", err)
		return 0, err
	}

	valueFloat, _ := new(big.Float).SetInt(value).Float64()

	usdValue := valueFloat / UsdBase

	return usdValue, nil
}

// GetUnknownTokenValue checks token/WBNB
func GetUnknownTokenValue(tokenAddr string) (float64, error) {
	var tokenPoolInfo map[string]string
	var ok bool
	if tokenPoolInfo, ok = consts.WhitelistTokensPoolMap[tokenAddr]; !ok {
		utils.Logger.Warn("token not in known list or whitelist: %s", tokenAddr)
		return 0, nil
	}

	tokenPoolAddr := tokenPoolInfo["pool"]

	tokenPool, err := pancake.NewPancakePair(common.HexToAddress(tokenPoolAddr), BscConn)
	if err != nil {
		utils.Logger.Error("GetUnknownTokenValue NewPancakePair err: %v", err)
		return 0, err
	}

	value, err := calculateUnknownTokenValueFromPool(tokenAddr, tokenPoolInfo["pair"], tokenPool)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func calculateUnknownTokenValueFromPool(tokenAddr, pairAddr string, pancakePool *pancake.PancakePair) (float64, error) {
	resp, err := pancakePool.GetReserves(nil)
	if err != nil {
		utils.Logger.Error("GetUnknownTokenValue GetReserves err: %v", err)
		return 0, err
	}

	token0Addr, err := pancakePool.Token0(nil)
	if err != nil {
		utils.Logger.Error("GetUnknownTokenValue get Token0 err: %v", err)
		return 0, err
	}

	//bep20Token, err := token.NewTokenBEP20(common.HexToAddress(tokenAddr), BscConn)
	//if err != nil {
	//	utils.Logger.Error("GetUnknownTokenValue NewTokenBEP20 err: %v", err)
	//	return 0, err
	//}
	//
	//decimal1, err := bep20Token.Decimals(nil)
	//if err != nil {
	//	utils.Logger.Error("GetUnknownTokenValue get Decimals err: %v", err)
	//	return 0, err
	//}

	var pairReserve, tokenReserve float64

	utils.Logger.Info("reserves: %v", resp)

	if token0Addr.Hex() == pairAddr {
		pairReserve, _ = new(big.Float).SetInt(resp.Reserve0).Float64()
		tokenReserve, _ = new(big.Float).SetInt(resp.Reserve1).Float64()
	} else {
		pairReserve, _ = new(big.Float).SetInt(resp.Reserve1).Float64()
		tokenReserve, _ = new(big.Float).SetInt(resp.Reserve0).Float64()
	}
	if pairReserve/BnbBase < 1 {
		utils.Logger.Warn("reserves too small, bnb: %f, token1: %f", pairReserve, tokenReserve)
		return 0, nil
	}

	tokenValueInPair := tokenReserve / pairReserve

	utils.Logger.Info("tokenr: %f, pairr: %f", tokenReserve, pairReserve)

	pairValue, err := GetTokenValue(pairAddr)
	if err != nil {
		utils.Logger.Error("get pair value err: %+v", err)
		return 0, nil
	}

	return pairValue * tokenValueInPair, nil
}
