package consts

const (
	BnbAddr              = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
	PancakeRouterAddr    = "0x10ED43C718714eb63d5aA57B78B54704E256024E"
	PancakeFactoryAddr   = "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"
	PancakeRouterAbi     = `[{"inputs":[{"internalType":"address","name":"_factory","type":"address"},{"internalType":"address","name":"_WETH","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"WETH","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"amountADesired","type":"uint256"},{"internalType":"uint256","name":"amountBDesired","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amountTokenDesired","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountIn","outputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountOut","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsIn","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsOut","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"reserveA","type":"uint256"},{"internalType":"uint256","name":"reserveB","type":"uint256"}],"name":"quote","outputs":[{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETHSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermit","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityWithPermit","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapETHForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETHSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]`
	PancakeFactoryAbi    = `[{"inputs":[{"internalType":"address","name":"_feeToSetter","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"uint256","name":"","type":"uint256"}],"name":"PairCreated","type":"event"},{"constant":true,"inputs":[],"name":"INIT_CODE_PAIR_HASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"allPairs","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"allPairsLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"createPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"feeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"feeToSetter","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_feeTo","type":"address"}],"name":"setFeeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_feeToSetter","type":"address"}],"name":"setFeeToSetter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	BscChainlinkPriceAbi = `[{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"description","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint80","name":"_roundId","type":"uint80"}],"name":"getRoundData","outputs":[{"internalType":"uint80","name":"roundId","type":"uint80"},{"internalType":"int256","name":"answer","type":"int256"},{"internalType":"uint256","name":"startedAt","type":"uint256"},{"internalType":"uint256","name":"updatedAt","type":"uint256"},{"internalType":"uint80","name":"answeredInRound","type":"uint80"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"latestRoundData","outputs":[{"internalType":"uint80","name":"roundId","type":"uint80"},{"internalType":"int256","name":"answer","type":"int256"},{"internalType":"uint256","name":"startedAt","type":"uint256"},{"internalType":"uint256","name":"updatedAt","type":"uint256"},{"internalType":"uint80","name":"answeredInRound","type":"uint80"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
)

var (
	StableCoinsMap = map[string]string{
		"DAI":  "0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3", //DAI
		"USDT": "0x55d398326f99059ff775485246999027b3197955", //USDT
		"BUSD": "0xe9e7cea3dedca5984780bafc599bd69add087d56", //BUSD
		"USDC": "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d", //USDC
		"TUSD": "0x14016e85a25aeb13065688cafb43044c2ef86784", //TUSD
		"USDN": "0x03ab98f5dc94996f8c33e15cd4468794d12d41f9", //USDN
		"UST":  "0x23396cf899ca06c4472205fc903bdb4de249d6fc", //UST
	}

	TokenChainlinkProxyMap = map[string]string{
		"0xfb6115445bff7b52feb98650c87f44907e58f802": "0xA8357BF572460fC40f4B0aCacbB2a6A61c89f475", //AAVE / USD
		"0x3ee2200efb3400fabb9aacf31297cbdd1d435d47": "0xa767f745331D267c7751297D982b050c93985627", //ADA / USD
		"0x8f0528ce5ef7b51152a59745befdd91d97091d2f": "0xe0073b60833249ffd1bb2af809112c2fbf221DF6", //ALPACA / USD
		"0x0eb3a705fc54725037cc9e008bdede697f62f335": "0xb056B7C804297279A9a673289264c17E6Dc6055d", //ATOM / USD
		"0xa184088a740c695e156f91f5cc086a06bb78b827": "0x88E71E6520E5aC75f5338F5F0c9DeD9d4f692cDA", //AUTO / USD
		"0x715d400f88c167884bbcc41c5fea407ed4d2f8a0": "0x7B49524ee5740c99435f52d731dFC94082fE61Ab", //AXS / USD
		"0x83085138ed8a96ec0cbd2013ddfe1ebb975940cc": "0x368b7ab0a0Ff94E23fF5e4A7F04327dF7079E174", //BAC / USD
		"0xad6caeb32cd2c308980a548bd0bc5aa4306c6c18": "0xC78b99Ae87fF43535b0C782128DB3cB49c74A4d3", //BAND / USD
		"0x8ff795a6f4d97e7887c79bea79aba5cc76444adf": "0x43d80f616DAf0b0B42a928EeD32147dC59027D41", //BCH / USD
		"0x250632378e573c6be1ac2f97fcdf00515d0aa91b": "0x2A3796273d47c4eD363b361D3AEFb7F7E2A13782", //BETH / USD
		"0xca3f508b8e4dd382ee878a314789373d80a5190a": "0xaB827b69daCd586A37E80A7d552a4395d576e645", //BIFI / USD
		"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c": "0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE", //BNB / USD
		"0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c": "0x264990fbd0A4796A3E3d8E37C4d5F87a3aCa5Ebf", //BTC / USD
		"0xe9e7cea3dedca5984780bafc599bd69add087d56": "0xcBb98864Ef56E9042e7d2efef76141f15731B82f", //BUSD / USD
		"0x4b87642aedf10b642be4663db842ecc5a88bf5ba": "0xFc362828930519f236ad0c8f126B7996562a695A", //BZRX / USD
		"0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82": "0xB6064eD41d4f67e353768aA239cA86f4F73665a1", //CAKE / USD
		"0x045c4324039dA91c52C55DF5D785385Aab073DcF": "0xe3cA2f3Bad1D8327820f648C759f17162b5383ae", //CFX / USD
		"0x52ce071bd9b1c4b00a0b92d298c512478cad67e8": "0x0Db8945f9aEf5651fa5bd52314C5aAe78DfDe540", //COMP / USD
		"0xd4cb328a82bdf5f03eb737f37fa6b370aef3e888": "0xa12Fc27A873cf114e6D8bBAf8BD9b8AC56110b39", //CREAM / USD
		"0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3": "0x132d3C0B1D2cEa0BC552588063bdBb210FDeecfA", //DAI / USD
		"0x3fda9383a84c05ec8f7630fe10adf1fac13241cc": "0x39F1275366D130eB677D4F47D40F9296B62D877A", //DEGO / USD
		"0x4a9a2b2b04549c3927dd2c9668a5ef3fca473623": "0x1b816F5E122eFa230300126F97C018716c4e47F5", //DF / USD
		"0x67ee3cb086f8a16f34bee3ca72fad36f7db929e2": "0x87701B15C08687341c2a847ca44eCfBc8d7873E1", //DODO / USD
		"0xba2ae424d960c26247dd6c32edc70b295c744c43": "0x3AB0A0d137D4F946fBB19eecc6e92E64660231C8", //DOGE / USD
		"0x7083609fce4d1d8dc0c979aab8c869ea2c873402": "0xC333eb0086309a16aa7c8308DfD32c8BBA0a2592", //DOT / USD
		"0xab671ea900a8e3c959f5bd29ceb5e370ba75bb9e": "0x7ee7E7847FFC93F8Cf67BCCc0002afF9C52DE524", //DPI / USD
		"0x56b6fb708fc5732dec1afc8d8556423a2edccbd6": "0xd5508c8Ffdb8F15cE336e629fD4ca9AdB48f50F0", //EOS / USD
		"0x2170ed0880ac9a755fd29b2688956bd959f933f8": "0x9ef1B8c0E4F7dc8bF5719Ea496883DC6401d5b2e", //ETH / USD
		"0x0d8ce2a99bb6e3b7db580ed848240e4a0f9ae153": "0xE5dbFD9003bFf9dF5feB2f4F445Ca00fb121fb83", //FIL / USD
		"0xa2b726b1145a4773f68593cf171187d8ebe4d495": "0x63A9133cd7c611d6049761038C16f238FddA71d7", //INJ / USD
		"0x762539b45a1dcce3d36d080f74d1aed37844b878": "0x38393201952f2764E04B290af9df427217D56B41", //LINA / USD
		"0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd": "0xca236E327F629f9Fc2c30A4E95775EbF0B89fac8", //LINK / USD
		"0xb59490ab09a0f526cc7305822ac65f2ab12f9723": "0x83766bA8d964fEAeD3819b145a69c947Df9Cb035", //LIT / USD
		"0x4338665cbb7b2485a8855a139b75d5e34ab0db94": "0x74E72F37A8c415c8f1a98Ed42E78Ff997435791D", //LTC / USD
		"0x2ed9a5c8c13b93955103b9a7c167b67ef4d568a3": "0x4978c0abE6899178c1A74838Ee0062280888E2Cf", //MASK / USD
		"0xcc42724c6683b7e57334c4e856f4c9965ed682bd": "0x7CA57b0cA6367191c94C8914d7Df09A57655905f", //MATIC / USD
		"0x5b6dcf557e2abe2323c48445e8cc948910d8c2c9": "0x291B2983b995870779C36A102Da101f8765244D6", //MIR / USD
		"0x8cd6e29d3686d24d3c2018cee54621ea0f89313b": "0xaCFBE73231d312AC6954496b3f786E892bF0f7e5", //NULS / USD
		"0xfd7b3a77848f1c2d67e05e54d78d174a0c850335": "0x887f177CBED2cf555a64e7bF125E1825EB69dB82", //ONT / USD
		"0x7950865a9140cb519342433146ed5b40c6f210f7": "0x7F8caD4690A38aC28BDA3D132eF83DB1C17557Df", //PAXG / USD
		"0x8519ea49c997f50ceffa444d240fb655e89248aa": "0xD1225da5FC21d17CaE526ee4b6464787c6A71b4C", //RAMP / USD
		"0xf21768ccbc73ea5b6fd3c687208a7c2def2d966e": "0x46f13472A4d4FeC9E07E8A00eE52f4Fa77810736", //REEF / USD
		"0x947950bcc74888a40ffa2593c5798f11fc9124c4": "0xa679C72a97B654CFfF58aB704de3BA15Cde89B07", //SUSHI / USD
		"0x47bead2563dcbf3bf2c9407fea4dc236faba485a": "0xE188A9875af525d25334d75F3327863B2b8cd0F1", //SXP / USD
		"0x85eac5ac2f758618dfa09bdbe0cf174e7d574d5b": "0xF4C5e535756D11994fCBB12Ba8adD0192D9b88be", //TRX / USD
		"0x14016e85a25aeb13065688cafb43044c2ef86784": "0xa3334A9762090E827413A7495AfeCE76F41dFc06", //TUSD / USD
		"0xbf5140a22578168fd562dccf235e5d43a02ce9b1": "0xb57f259E7C24e56a1dA00F66b55A5640d9f9E7e4", //UNI / USD
		"0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d": "0x51597f405303C4377E36123cBc172b13269EA163", //USDC / USD
		"0x03ab98f5dc94996f8c33e15cd4468794d12d41f9": "0x7C0BC703Dc56645203CFeBE1928E34B8e885ae37", //USDN / USD
		"0x55d398326f99059ff775485246999027b3197955": "0xB97Ad0E74fa7d920791E90258A6E2085088b4320", //USDT / USD
		"0x23396cf899ca06c4472205fc903bdb4de249d6fc": "0xcbf8518F8727B8582B22837403cDabc53463D462", //UST / USD
		"0x4bd17003473389a42daf6a0a729f6fdb328bbbd7": "0x058316f8Bb13aCD442ee7A216C7b60CFB4Ea1B53", //VAI / USD
		"0x1d2f0da169ceb9fc7b3144628db156f3f6c60dbe": "0x93A67D414896A280bF8FFB3b389fE3686E014fda", //XRP / USD
		"0x16939ef78684453bfdfb47825f8a5f714f12623a": "0x9A18137ADCF7b05f033ad26968Ed5a9cf0Bf8E6b", //XTZ / USD
		"0xcf6bb5389c92bdda8a3747ddb454cb7a64626c63": "0xBF63F430A79D4036A5900C19818aFf1fa710f206", //XVS / USD
		"0x88f1a5ae2a3bf98aeaf342d26b30a79438c9142e": "0xD7eAa5Bf3013A96e3d515c055Dbd98DbdC8c620D", //YFI / USD
		"0x7f70642d88cf1c4a3a7abb072b53b929b653eda5": "0xC94580FAaF145B2FD0ab5215031833c98D3B77E4", //YFII / USD
		"0xb86abcb37c3a4b64f74f59301aff131a1becc787": "0x3e3aA4FC329529C8Ab921c810850626021dbA7e6", //ZIL / USD
	}

	WhitelistTokensPoolMap = map[string]map[string]string{
		"0xe02df9e3e622debdd69fb838bb799e3f168902c5": {
			"pool": "0xc2Eed0F5a0dc28cfa895084bC0a9B8B8279aE492",
			"pair": BnbAddr,
		}, // BAKE
		"0xc9849e6fdb743d08faee3e34dd2d1bc69ea11a51": {
			"pool": "0x5aFEf8567414F29f0f927A0F2787b188624c10E2",
			"pair": BnbAddr,
		}, // BUNNY
		"0xaec945e04baf28b135fa7c640f624f8d90f1c3a6": {
			"pool": "0x92247860A03F48d5c6425c7CA35CDcFCB1013AA1",
			"pair": BnbAddr,
		}, // COIN98
		"0xda6802bbec06ab447a68294a63de47ed4506acaa": {
			"pool": "0x7625aae940c35505d069deF82D5469b74a36550B",
			"pair": BnbAddr,
		}, // CRYPT
		"0xa7f552078dcc247c2684336020c03648500c6d9f": {
			"pool": "0xf9045866e7b372DeF1EFf3712CE55FAc1A98dAF0",
			"pair": BnbAddr,
		}, // EPS (also on binance)
		"0x42f6f551ae042cbe50c739158b4f0cac0edb9096": {
			"pool": "0x401479091d0F7b8AE437Ee8B054575cd33ea72Bd",
			"pair": StableCoinsMap["BUSD"],
		}, // NRV (BUSD)
		"0x8076c74c5e3f5852037f31ff0093eeb8c8add8d3": {
			"pool": "0x9adc6Fb78CEFA07E13E9294F150C1E8C1Dd566c0",
			"pair": BnbAddr,
		}, // SAFEMOON
		"0x55b53855eae06c4744841dbfa06fce335db4355b": {
			"pool": "0x16cc94254b282C6B4FBcD671aA9E5E6Aebf52Fa3",
			"pair": BnbAddr,
		}, // SSB
		"0x9f589e3eabe42ebc94a44727b3f3531c0c877809": {
			"pool": "0xFFd4B200d3C77A0B691B5562D804b3bd54294e6e",
			"pair": BnbAddr,
		}, // TKO
		"0x6169b3b23e57de79a6146a2170980ceb1f83b9e0": {
			"pool": "0x39D26BcA778a528eBDD57b136f777820Be5eDd1a",
			"pair": BnbAddr,
		}, // VETTER
	}
)
