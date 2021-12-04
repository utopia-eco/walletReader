package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/utopia-eco/walletReader/abi/abigen/chainlink"
	"github.com/utopia-eco/walletReader/utils"
)

const (
	UsdBase = 100000000
	BnbBase = 1000000000000000000
)

var (
	BnbPrice float64
)

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

	usdValue := float64(value.Int64()) / UsdBase

	return usdValue, nil
}

// TODO: add calculate unknown tokens
//func GetUnknownTokenValue(tokenAddr common.Address, conn *bind.ContractBackend) (float64, error) {
//	pairAddr, err := PancakeFactory.GetPair(nil, common.HexToAddress(consts.BnbAddr), tokenAddr)
//	if err != nil {
//		utils.Logger.Error("GetUnknownTokenValue GetPair err: %v", err)
//		return 0, err
//	}
//
//	pancakePool, err := pancake.NewPancakePair(pairAddr, *conn)
//	if err != nil {
//		utils.Logger.Error("GetUnknownTokenValue NewPancakePair err: %v", err)
//		return 0, err
//	}
//
//	value, err := calculateUnknownTokenValueFromPool(tokenAddr, pancakePool)
//	if err != nil {
//		return 0, err
//	}
//
//	return value, nil
//}
//
//

//func calculateUnknownTokenValueFromPool(tokenAddr common.Address, pancakePool *pancake.PancakePair) (float64, error) {
//	resp, err := pancakePool.GetReserves(nil)
//	if err != nil {
//		utils.Logger.Error("GetUnknownTokenValue GetReserves err: %v", err)
//		return 0, err
//	}
//
//	token0Addr, err := pancakePool.Token0(nil)
//	if err != nil {
//		utils.Logger.Error("GetUnknownTokenValue get Token0 err: %v", err)
//		return 0, err
//	}
//
//	bep20Token, err := token.NewTokenBEP20(tokenAddr, BscConn)
//	if err != nil {
//		utils.Logger.Error("GetUnknownTokenValue NewTokenBEP20 err: %v", err)
//		return 0, err
//	}
//
//	decimal1, err := bep20Token.Decimals(nil)
//	if err != nil {
//		utils.Logger.Error("GetUnknownTokenValue get Decimals err: %v", err)
//		return 0, err
//	}
//
//	var reserve0, reserve1 float64
//
//	if token0Addr.Hex() == consts.BnbAddr {
//		reserve0 = float64(resp.Reserve0.Int64())
//		reserve1 = float64(resp.Reserve1.Int64())
//	} else {
//		reserve0 = float64(resp.Reserve1.Int64())
//		reserve1 = float64(resp.Reserve0.Int64())
//	}
//
//	return calculateValueFromReserves(reserve0, reserve1, float64(decimal1.Int64())), nil
//}
//
//func calculateValueFromReserves(reserve0, reserve1, decimal1 float64) float64 {
//	k := reserve0 * reserve1
//	if (k / BnbBase / decimal1) < 1 {
//		utils.Logger.Warn("reserves too small, bnb: %f, token1: %f", reserve0, reserve1)
//		return 0
//	}
//}
