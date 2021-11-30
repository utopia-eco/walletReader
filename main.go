package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/utopia-eco/walletReader/consts"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func main()  {
	ctx := context.Background()
	bscConn, err := ethclient.Dial(consts.BscBinanceUrl)
	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}
	quickswap, err := NewQuickSwap(common.HexToAddress(quickswapAddr), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	polyWalletAddr := common.HexToAddress(polyWallet)
	nonce, _ := conn.NonceAt(ctx, polyWalletAddr, nil)
	fmt.Printf("nonce at: %v\n", nonce)

	startTime := time.Now()
	fmt.Printf("start time: %v\n", startTime)

	// start one flow
	event := &models.Event{}
	if err = json.Unmarshal([]byte(test.Input), event); err != nil {
		log.Panicf("Failed to unmarshal event json: %v", err)
	}

	log.Printf("event: %+v", event)

	path := []common.Address{
		common.HexToAddress(wmaticAddr),
		common.HexToAddress(usdcAddr),
	}
	amountsOut, err := quickswap.GetAmountsOut(nil, big.NewInt(amountIn), path)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}

	fmt.Println("Amounts out:", amountsOut)
	//fmt.Printf("start time: %v\n", time.Now())

	pendingNonce, _ := conn.PendingNonceAt(ctx, polyWalletAddr)
	fmt.Printf("pending nonce: %v\n", pendingNonce)
	fmt.Printf("start time: %v\n", time.Now())

	fmt.Println(ctx.Deadline())

	//txOpts := &bind.TransactOpts{
	//	From:      polyWalletAddr,
	//	Nonce:     nil,
	//	Signer:    nil,
	//	Value:     nil,
	//	GasPrice:  nil,
	//	GasFeeCap: nil,
	//	GasTipCap: nil,
	//	GasLimit:  0,
	//	Context:   nil,
	//	NoSend:    false,
	//}
	//tx := quickswap.SwapExactETHForTokens(
	//
	//	)
}
