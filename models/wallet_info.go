package models

type WalletTokensValueReq struct {
	WalletAddress string   `json:"wallet_address"`
	Tokens        []*Token `json:"tokens,omitempty"`
}

type Token struct {
	TokenAddress string  `json:"token_address"`
	Symbol       string  `json:"symbol,omitempty"`
	Amount       float64 `json:"amount"`
}

type WalletTokensValueResp struct {
	WalletAddress string        `json:"wallet_address"`
	WalletValue   float64       `json:"total_value"`
	TokenValues   []*TokenValue `json:"token_values"`
}

type TokenValue struct {
	TokenAddress string  `json:"token_address"`
	UnitPrice    float64 `json:"unit_price"`
	TotalValue   float64 `json:"value"`
}
