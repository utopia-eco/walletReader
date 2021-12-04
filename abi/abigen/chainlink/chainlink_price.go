// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlink

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChainlinkPriceABI is the input ABI used to generate the binding from.
const ChainlinkPriceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"confirmAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"phaseAggregators\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phaseId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"proposeAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedAggregator\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"proposedGetRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedLatestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChainlinkPrice is an auto generated Go binding around an Ethereum contract.
type ChainlinkPrice struct {
	ChainlinkPriceCaller     // Read-only binding to the contract
	ChainlinkPriceTransactor // Write-only binding to the contract
	ChainlinkPriceFilterer   // Log filterer for contract events
}

// ChainlinkPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkPriceSession struct {
	Contract     *ChainlinkPrice   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainlinkPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkPriceCallerSession struct {
	Contract *ChainlinkPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChainlinkPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkPriceTransactorSession struct {
	Contract     *ChainlinkPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChainlinkPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkPriceRaw struct {
	Contract *ChainlinkPrice // Generic contract binding to access the raw methods on
}

// ChainlinkPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkPriceCallerRaw struct {
	Contract *ChainlinkPriceCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkPriceTransactorRaw struct {
	Contract *ChainlinkPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkPrice creates a new instance of ChainlinkPrice, bound to a specific deployed contract.
func NewChainlinkPrice(address common.Address, backend bind.ContractBackend) (*ChainlinkPrice, error) {
	contract, err := bindChainlinkPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPrice{ChainlinkPriceCaller: ChainlinkPriceCaller{contract: contract}, ChainlinkPriceTransactor: ChainlinkPriceTransactor{contract: contract}, ChainlinkPriceFilterer: ChainlinkPriceFilterer{contract: contract}}, nil
}

// NewChainlinkPriceCaller creates a new read-only instance of ChainlinkPrice, bound to a specific deployed contract.
func NewChainlinkPriceCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkPriceCaller, error) {
	contract, err := bindChainlinkPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceCaller{contract: contract}, nil
}

// NewChainlinkPriceTransactor creates a new write-only instance of ChainlinkPrice, bound to a specific deployed contract.
func NewChainlinkPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkPriceTransactor, error) {
	contract, err := bindChainlinkPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceTransactor{contract: contract}, nil
}

// NewChainlinkPriceFilterer creates a new log filterer instance of ChainlinkPrice, bound to a specific deployed contract.
func NewChainlinkPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkPriceFilterer, error) {
	contract, err := bindChainlinkPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFilterer{contract: contract}, nil
}

// bindChainlinkPrice binds a generic wrapper to an already deployed contract.
func bindChainlinkPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkPriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkPrice *ChainlinkPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkPrice.Contract.ChainlinkPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkPrice *ChainlinkPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.ChainlinkPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkPrice *ChainlinkPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.ChainlinkPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkPrice *ChainlinkPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkPrice *ChainlinkPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkPrice *ChainlinkPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.contract.Transact(opts, method, params...)
}

//// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
////
//// Solidity: function accessController() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCaller) AccessController(opts *bind.CallOpts) (common.Address, error) {
//	var (
//		ret0 = new(common.Address)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "accessController")
//	return *ret0, err
//}
//
//// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
////
//// Solidity: function accessController() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceSession) AccessController() (common.Address, error) {
//	return _ChainlinkPrice.Contract.AccessController(&_ChainlinkPrice.CallOpts)
//}
//
//// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
////
//// Solidity: function accessController() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) AccessController() (common.Address, error) {
//	return _ChainlinkPrice.Contract.AccessController(&_ChainlinkPrice.CallOpts)
//}
//
//// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
////
//// Solidity: function aggregator() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
//	var (
//		ret0 = new(common.Address)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "aggregator")
//	return *ret0, err
//}
//
//// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
////
//// Solidity: function aggregator() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceSession) Aggregator() (common.Address, error) {
//	return _ChainlinkPrice.Contract.Aggregator(&_ChainlinkPrice.CallOpts)
//}
//
//// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
////
//// Solidity: function aggregator() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) Aggregator() (common.Address, error) {
//	return _ChainlinkPrice.Contract.Aggregator(&_ChainlinkPrice.CallOpts)
//}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ChainlinkPrice *ChainlinkPriceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChainlinkPrice.contract.Call(opts, &out, "decimals")
	ret0 := abi.ConvertType(out[0], new(uint8)).(uint8)
	return ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ChainlinkPrice *ChainlinkPriceSession) Decimals() (uint8, error) {
	return _ChainlinkPrice.Contract.Decimals(&_ChainlinkPrice.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ChainlinkPrice *ChainlinkPriceCallerSession) Decimals() (uint8, error) {
	return _ChainlinkPrice.Contract.Decimals(&_ChainlinkPrice.CallOpts)
}

//// Description is a free data retrieval call binding the contract method 0x7284e416.
////
//// Solidity: function description() constant returns(string)
//func (_ChainlinkPrice *ChainlinkPriceCaller) Description(opts *bind.CallOpts) (string, error) {
//	var (
//		ret0 = new(string)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "description")
//	return *ret0, err
//}
//
//// Description is a free data retrieval call binding the contract method 0x7284e416.
////
//// Solidity: function description() constant returns(string)
//func (_ChainlinkPrice *ChainlinkPriceSession) Description() (string, error) {
//	return _ChainlinkPrice.Contract.Description(&_ChainlinkPrice.CallOpts)
//}
//
//// Description is a free data retrieval call binding the contract method 0x7284e416.
////
//// Solidity: function description() constant returns(string)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) Description() (string, error) {
//	return _ChainlinkPrice.Contract.Description(&_ChainlinkPrice.CallOpts)
//}
//
//// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
////
//// Solidity: function getAnswer(uint256 _roundId) constant returns(int256)
//func (_ChainlinkPrice *ChainlinkPriceCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
//	var (
//		ret0 = new(*big.Int)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "getAnswer", _roundId)
//	return *ret0, err
//}
//
//// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
////
//// Solidity: function getAnswer(uint256 _roundId) constant returns(int256)
//func (_ChainlinkPrice *ChainlinkPriceSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
//	return _ChainlinkPrice.Contract.GetAnswer(&_ChainlinkPrice.CallOpts, _roundId)
//}
//
//// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
////
//// Solidity: function getAnswer(uint256 _roundId) constant returns(int256)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
//	return _ChainlinkPrice.Contract.GetAnswer(&_ChainlinkPrice.CallOpts, _roundId)
//}

//// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
////
//// Solidity: function getRoundData(uint80 _roundId) constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	ret := new(struct {
//		RoundId         *big.Int
//		Answer          *big.Int
//		StartedAt       *big.Int
//		UpdatedAt       *big.Int
//		AnsweredInRound *big.Int
//	})
//	out := ret
//	err := _ChainlinkPrice.contract.Call(opts, out, "getRoundData", _roundId)
//	return *ret, err
//}
//
//// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
////
//// Solidity: function getRoundData(uint80 _roundId) constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceSession) GetRoundData(_roundId *big.Int) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.GetRoundData(&_ChainlinkPrice.CallOpts, _roundId)
//}
//
//// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
////
//// Solidity: function getRoundData(uint80 _roundId) constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) GetRoundData(_roundId *big.Int) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.GetRoundData(&_ChainlinkPrice.CallOpts, _roundId)
//}
//
//// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
////
//// Solidity: function getTimestamp(uint256 _roundId) constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
//	var (
//		ret0 = new(*big.Int)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "getTimestamp", _roundId)
//	return *ret0, err
//}
//
//// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
////
//// Solidity: function getTimestamp(uint256 _roundId) constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
//	return _ChainlinkPrice.Contract.GetTimestamp(&_ChainlinkPrice.CallOpts, _roundId)
//}
//
//// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
////
//// Solidity: function getTimestamp(uint256 _roundId) constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
//	return _ChainlinkPrice.Contract.GetTimestamp(&_ChainlinkPrice.CallOpts, _roundId)
//}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_ChainlinkPrice *ChainlinkPriceCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPrice.contract.Call(opts, &out, "latestAnswer")
	ret0 := abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return *ret0, err
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_ChainlinkPrice *ChainlinkPriceSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkPrice.Contract.LatestAnswer(&_ChainlinkPrice.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_ChainlinkPrice *ChainlinkPriceCallerSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkPrice.Contract.LatestAnswer(&_ChainlinkPrice.CallOpts)
}

//// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
////
//// Solidity: function latestRound() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
//	var (
//		ret0 = new(*big.Int)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "latestRound")
//	return *ret0, err
//}
//
//// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
////
//// Solidity: function latestRound() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceSession) LatestRound() (*big.Int, error) {
//	return _ChainlinkPrice.Contract.LatestRound(&_ChainlinkPrice.CallOpts)
//}
//
//// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
////
//// Solidity: function latestRound() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) LatestRound() (*big.Int, error) {
//	return _ChainlinkPrice.Contract.LatestRound(&_ChainlinkPrice.CallOpts)
//}
//
//// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
////
//// Solidity: function latestRoundData() constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	ret := new(struct {
//		RoundId         *big.Int
//		Answer          *big.Int
//		StartedAt       *big.Int
//		UpdatedAt       *big.Int
//		AnsweredInRound *big.Int
//	})
//	out := ret
//	err := _ChainlinkPrice.contract.Call(opts, out, "latestRoundData")
//	return *ret, err
//}
//
//// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
////
//// Solidity: function latestRoundData() constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceSession) LatestRoundData() (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.LatestRoundData(&_ChainlinkPrice.CallOpts)
//}
//
//// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
////
//// Solidity: function latestRoundData() constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) LatestRoundData() (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.LatestRoundData(&_ChainlinkPrice.CallOpts)
//}
//
//// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
////
//// Solidity: function latestTimestamp() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
//	var (
//		ret0 = new(*big.Int)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "latestTimestamp")
//	return *ret0, err
//}
//
//// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
////
//// Solidity: function latestTimestamp() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceSession) LatestTimestamp() (*big.Int, error) {
//	return _ChainlinkPrice.Contract.LatestTimestamp(&_ChainlinkPrice.CallOpts)
//}
//
//// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
////
//// Solidity: function latestTimestamp() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) LatestTimestamp() (*big.Int, error) {
//	return _ChainlinkPrice.Contract.LatestTimestamp(&_ChainlinkPrice.CallOpts)
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
//	var (
//		ret0 = new(common.Address)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "owner")
//	return *ret0, err
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceSession) Owner() (common.Address, error) {
//	return _ChainlinkPrice.Contract.Owner(&_ChainlinkPrice.CallOpts)
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) Owner() (common.Address, error) {
//	return _ChainlinkPrice.Contract.Owner(&_ChainlinkPrice.CallOpts)
//}

//// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
////
//// Solidity: function phaseAggregators(uint16 ) constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCaller) PhaseAggregators(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
//	var (
//		ret0 = new(common.Address)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "phaseAggregators", arg0)
//	return *ret0, err
//}
//
//// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
////
//// Solidity: function phaseAggregators(uint16 ) constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceSession) PhaseAggregators(arg0 uint16) (common.Address, error) {
//	return _ChainlinkPrice.Contract.PhaseAggregators(&_ChainlinkPrice.CallOpts, arg0)
//}
//
//// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
////
//// Solidity: function phaseAggregators(uint16 ) constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) PhaseAggregators(arg0 uint16) (common.Address, error) {
//	return _ChainlinkPrice.Contract.PhaseAggregators(&_ChainlinkPrice.CallOpts, arg0)
//}
//
//// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
////
//// Solidity: function phaseId() constant returns(uint16)
//func (_ChainlinkPrice *ChainlinkPriceCaller) PhaseId(opts *bind.CallOpts) (uint16, error) {
//	var (
//		ret0 = new(uint16)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "phaseId")
//	return *ret0, err
//}
//
//// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
////
//// Solidity: function phaseId() constant returns(uint16)
//func (_ChainlinkPrice *ChainlinkPriceSession) PhaseId() (uint16, error) {
//	return _ChainlinkPrice.Contract.PhaseId(&_ChainlinkPrice.CallOpts)
//}
//
//// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
////
//// Solidity: function phaseId() constant returns(uint16)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) PhaseId() (uint16, error) {
//	return _ChainlinkPrice.Contract.PhaseId(&_ChainlinkPrice.CallOpts)
//}
//
//// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
////
//// Solidity: function proposedAggregator() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCaller) ProposedAggregator(opts *bind.CallOpts) (common.Address, error) {
//	var (
//		ret0 = new(common.Address)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "proposedAggregator")
//	return *ret0, err
//}
//
//// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
////
//// Solidity: function proposedAggregator() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceSession) ProposedAggregator() (common.Address, error) {
//	return _ChainlinkPrice.Contract.ProposedAggregator(&_ChainlinkPrice.CallOpts)
//}
//
//// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
////
//// Solidity: function proposedAggregator() constant returns(address)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) ProposedAggregator() (common.Address, error) {
//	return _ChainlinkPrice.Contract.ProposedAggregator(&_ChainlinkPrice.CallOpts)
//}
//
//// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
////
//// Solidity: function proposedGetRoundData(uint80 _roundId) constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCaller) ProposedGetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	ret := new(struct {
//		RoundId         *big.Int
//		Answer          *big.Int
//		StartedAt       *big.Int
//		UpdatedAt       *big.Int
//		AnsweredInRound *big.Int
//	})
//	out := ret
//	err := _ChainlinkPrice.contract.Call(opts, out, "proposedGetRoundData", _roundId)
//	return *ret, err
//}
//
//// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
////
//// Solidity: function proposedGetRoundData(uint80 _roundId) constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceSession) ProposedGetRoundData(_roundId *big.Int) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.ProposedGetRoundData(&_ChainlinkPrice.CallOpts, _roundId)
//}
//
//// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
////
//// Solidity: function proposedGetRoundData(uint80 _roundId) constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) ProposedGetRoundData(_roundId *big.Int) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.ProposedGetRoundData(&_ChainlinkPrice.CallOpts, _roundId)
//}
//
//// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
////
//// Solidity: function proposedLatestRoundData() constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCaller) ProposedLatestRoundData(opts *bind.CallOpts) (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	ret := new(struct {
//		RoundId         *big.Int
//		Answer          *big.Int
//		StartedAt       *big.Int
//		UpdatedAt       *big.Int
//		AnsweredInRound *big.Int
//	})
//	out := ret
//	err := _ChainlinkPrice.contract.Call(opts, out, "proposedLatestRoundData")
//	return *ret, err
//}
//
//// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
////
//// Solidity: function proposedLatestRoundData() constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceSession) ProposedLatestRoundData() (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.ProposedLatestRoundData(&_ChainlinkPrice.CallOpts)
//}
//
//// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
////
//// Solidity: function proposedLatestRoundData() constant returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) ProposedLatestRoundData() (struct {
//	RoundId         *big.Int
//	Answer          *big.Int
//	StartedAt       *big.Int
//	UpdatedAt       *big.Int
//	AnsweredInRound *big.Int
//}, error) {
//	return _ChainlinkPrice.Contract.ProposedLatestRoundData(&_ChainlinkPrice.CallOpts)
//}
//
//// Version is a free data retrieval call binding the contract method 0x54fd4d50.
////
//// Solidity: function version() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
//	var (
//		ret0 = new(*big.Int)
//	)
//	out := ret0
//	err := _ChainlinkPrice.contract.Call(opts, out, "version")
//	return *ret0, err
//}
//
//// Version is a free data retrieval call binding the contract method 0x54fd4d50.
////
//// Solidity: function version() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceSession) Version() (*big.Int, error) {
//	return _ChainlinkPrice.Contract.Version(&_ChainlinkPrice.CallOpts)
//}
//
//// Version is a free data retrieval call binding the contract method 0x54fd4d50.
////
//// Solidity: function version() constant returns(uint256)
//func (_ChainlinkPrice *ChainlinkPriceCallerSession) Version() (*big.Int, error) {
//	return _ChainlinkPrice.Contract.Version(&_ChainlinkPrice.CallOpts)
//}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkPrice *ChainlinkPriceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPrice.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkPrice *ChainlinkPriceSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.AcceptOwnership(&_ChainlinkPrice.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkPrice *ChainlinkPriceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.AcceptOwnership(&_ChainlinkPrice.TransactOpts)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactor) ConfirmAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.contract.Transact(opts, "confirmAggregator", _aggregator)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkPrice *ChainlinkPriceSession) ConfirmAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.ConfirmAggregator(&_ChainlinkPrice.TransactOpts, _aggregator)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactorSession) ConfirmAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.ConfirmAggregator(&_ChainlinkPrice.TransactOpts, _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactor) ProposeAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.contract.Transact(opts, "proposeAggregator", _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkPrice *ChainlinkPriceSession) ProposeAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.ProposeAggregator(&_ChainlinkPrice.TransactOpts, _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactorSession) ProposeAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.ProposeAggregator(&_ChainlinkPrice.TransactOpts, _aggregator)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactor) SetController(opts *bind.TransactOpts, _accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.contract.Transact(opts, "setController", _accessController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkPrice *ChainlinkPriceSession) SetController(_accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.SetController(&_ChainlinkPrice.TransactOpts, _accessController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactorSession) SetController(_accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.SetController(&_ChainlinkPrice.TransactOpts, _accessController)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkPrice *ChainlinkPriceSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.TransferOwnership(&_ChainlinkPrice.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkPrice *ChainlinkPriceTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkPrice.Contract.TransferOwnership(&_ChainlinkPrice.TransactOpts, _to)
}

// ChainlinkPriceAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the ChainlinkPrice contract.
type ChainlinkPriceAnswerUpdatedIterator struct {
	Event *ChainlinkPriceAnswerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChainlinkPriceAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkPriceAnswerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChainlinkPriceAnswerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChainlinkPriceAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkPriceAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkPriceAnswerUpdated represents a AnswerUpdated event raised by the ChainlinkPrice contract.
type ChainlinkPriceAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkPrice *ChainlinkPriceFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ChainlinkPriceAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceAnswerUpdatedIterator{contract: _ChainlinkPrice.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkPrice *ChainlinkPriceFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkPriceAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkPriceAnswerUpdated)
				if err := _ChainlinkPrice.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkPrice *ChainlinkPriceFilterer) ParseAnswerUpdated(log types.Log) (*ChainlinkPriceAnswerUpdated, error) {
	event := new(ChainlinkPriceAnswerUpdated)
	if err := _ChainlinkPrice.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkPriceNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the ChainlinkPrice contract.
type ChainlinkPriceNewRoundIterator struct {
	Event *ChainlinkPriceNewRound // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChainlinkPriceNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkPriceNewRound)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChainlinkPriceNewRound)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChainlinkPriceNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkPriceNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkPriceNewRound represents a NewRound event raised by the ChainlinkPrice contract.
type ChainlinkPriceNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkPrice *ChainlinkPriceFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ChainlinkPriceNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceNewRoundIterator{contract: _ChainlinkPrice.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkPrice *ChainlinkPriceFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ChainlinkPriceNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkPriceNewRound)
				if err := _ChainlinkPrice.contract.UnpackLog(event, "NewRound", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkPrice *ChainlinkPriceFilterer) ParseNewRound(log types.Log) (*ChainlinkPriceNewRound, error) {
	event := new(ChainlinkPriceNewRound)
	if err := _ChainlinkPrice.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkPriceOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the ChainlinkPrice contract.
type ChainlinkPriceOwnershipTransferRequestedIterator struct {
	Event *ChainlinkPriceOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChainlinkPriceOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkPriceOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChainlinkPriceOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChainlinkPriceOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkPriceOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkPriceOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the ChainlinkPrice contract.
type ChainlinkPriceOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkPrice *ChainlinkPriceFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainlinkPriceOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceOwnershipTransferRequestedIterator{contract: _ChainlinkPrice.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkPrice *ChainlinkPriceFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ChainlinkPriceOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkPriceOwnershipTransferRequested)
				if err := _ChainlinkPrice.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkPrice *ChainlinkPriceFilterer) ParseOwnershipTransferRequested(log types.Log) (*ChainlinkPriceOwnershipTransferRequested, error) {
	event := new(ChainlinkPriceOwnershipTransferRequested)
	if err := _ChainlinkPrice.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkPriceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainlinkPrice contract.
type ChainlinkPriceOwnershipTransferredIterator struct {
	Event *ChainlinkPriceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChainlinkPriceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkPriceOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChainlinkPriceOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChainlinkPriceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkPriceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkPriceOwnershipTransferred represents a OwnershipTransferred event raised by the ChainlinkPrice contract.
type ChainlinkPriceOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkPrice *ChainlinkPriceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainlinkPriceOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceOwnershipTransferredIterator{contract: _ChainlinkPrice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkPrice *ChainlinkPriceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainlinkPriceOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkPrice.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkPriceOwnershipTransferred)
				if err := _ChainlinkPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkPrice *ChainlinkPriceFilterer) ParseOwnershipTransferred(log types.Log) (*ChainlinkPriceOwnershipTransferred, error) {
	event := new(ChainlinkPriceOwnershipTransferred)
	if err := _ChainlinkPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
