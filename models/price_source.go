package models

type BakeApiInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		UpdateTime           int64  `json:"updateTime"`
		Address              string `json:"address"`
		TotalSupply          string `json:"totalSupply"`
		CirculationSupply    string `json:"circulationSupply"`
		CirculationMarketCap string `json:"circulationMarketCap"`
		Circulation          string `json:"circulation"`
		TotalLiquidityUSD    string `json:"totalLiquidityUSD"`
		Burned               string `json:"burned"`
		PriceUSD             string `json:"priceUSD"`
		PriceBNB             string `json:"priceBNB"`
		Volume24HUSD         string `json:"volume24hUSD"`
		TotalNft             string `json:"totalNft"`
		TotalNftTransaction  string `json:"totalNftTransaction"`
		TotalNftLockedBAKE   string `json:"totalNftLockedBAKE"`
		TotalNftVolumeBAKE   string `json:"totalNftVolumeBAKE"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
