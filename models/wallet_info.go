package models

type WalletTokensValueReq struct {
	WalletAddress string   `json:"wallet_address"`
	Tokens        []*Token `json:"tokens"`
}

type Token struct {
	TokenAddress string  `json:"token_address"`
	Symbol       string  `json:"symbol"`
	Amount       float64 `json:"amount"`
}

type WalletTokensValueResp struct {
	WalletAddress string        `json:"wallet_address"`
	TokenValues   []*TokenValue `json:"token_values"`
}

type TokenValue struct {
	TokenAddress string  `json:"token_address"`
	Value        float64 `json:"value"`
}
