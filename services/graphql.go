package services

//func QueryGraphQL() {
//	query := `{ ethereum(network: bsc) { dexTrades( options: {limit: 100, desc: ["block.timestamp.time"]} baseCurrency: {is: "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"} quoteCurrency: {is: "${address}"} date: {since: "2021-09-28"} ) { block { height timestamp { time(format: "%Y-%m-%d %H:%M:%S") } } tradeIndex protocol buyAmount buyCurrency { address symbol } sellAmount sellCurrency { address symbol } transaction { hash } } } } `,
//
//}
