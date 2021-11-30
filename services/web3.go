package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/utopia-eco/walletReader/abi/abigen/pancake"
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/utils"
	"log"
)

var (
	BscConn        *ethclient.Client
	PancakeRouter  *pancake.PancakeRouter
	PancakeFactory *pancake.PancakeFactory
)

func InitBsc() {
	var err error
	BscConn, err = ethclient.Dial(consts.BscBinanceUrl)
	if err != nil {
		utils.Logger.Fatal("Failed to connect to client: %v", err)
	}
}

func InitPancakes() {
	var err error
	PancakeRouter, err = pancake.NewPancakeRouter(common.HexToAddress(consts.PancakeRouterAddr), BscConn)
	if err != nil {
		log.Fatalf("Failed to instantiate PancakeRouter contract: %v", err)
	}

	PancakeFactory, err = pancake.NewPancakeFactory(common.HexToAddress(consts.PancakeFactoryAddr), BscConn)
	if err != nil {
		log.Fatalf("Failed to instantiate PancakeFactory contract: %v", err)
	}
}
