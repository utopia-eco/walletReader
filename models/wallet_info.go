package models

type WalletTokensValueReq struct {
	WalletAddress string   `json:"wallet_address"`
	Tokens        []*Token `json:"tokens,omitempty"`
}

type Token struct {
	TokenAddress string  `json:"token_address"`
	Symbol       string  `json:"symbol,omitempty"`
	Amount       float64 `json:"amount"`
	IsWrapped    *bool   `json:"is_wrapped,omitempty"`
}

type WalletTokensValueResp struct {
	WalletAddress string        `json:"wallet_address"`
	WalletValue   float64       `json:"total_value"`
	TokenValues   []*TokenValue `json:"token_values"`
}

type TokenValue struct {
	TokenAddress string  `json:"token_address"`
	TokenSymbol  string  `json:"token_symbol"`
	UnitPrice    float64 `json:"unit_price"`
	Amount       float64 `json:"amount"`
	TotalValue   float64 `json:"token_value"`
}
