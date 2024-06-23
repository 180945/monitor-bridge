// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// L2MetaData contains all meta data concerning the L2 contract.
var L2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_additionalData\",\"type\":\"bytes\"}],\"name\":\"WithdrawalWithMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1SharedBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"l1TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFromTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_additionalData\",\"type\":\"bytes\"}],\"name\":\"withdrawWithMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// L2ABI is the input ABI used to generate the binding from.
// Deprecated: Use L2MetaData.ABI instead.
var L2ABI = L2MetaData.ABI

// L2 is an auto generated Go binding around an Ethereum contract.
type L2 struct {
	L2Caller     // Read-only binding to the contract
	L2Transactor // Write-only binding to the contract
	L2Filterer   // Log filterer for contract events
}

// L2Caller is an auto generated read-only Go binding around an Ethereum contract.
type L2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2Session struct {
	Contract     *L2               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2CallerSession struct {
	Contract *L2Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TransactorSession struct {
	Contract     *L2Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2Raw is an auto generated low-level Go binding around an Ethereum contract.
type L2Raw struct {
	Contract *L2 // Generic contract binding to access the raw methods on
}

// L2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2CallerRaw struct {
	Contract *L2Caller // Generic read-only contract binding to access the raw methods on
}

// L2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TransactorRaw struct {
	Contract *L2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL2 creates a new instance of L2, bound to a specific deployed contract.
func NewL2(address common.Address, backend bind.ContractBackend) (*L2, error) {
	contract, err := bindL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2{L2Caller: L2Caller{contract: contract}, L2Transactor: L2Transactor{contract: contract}, L2Filterer: L2Filterer{contract: contract}}, nil
}

// NewL2Caller creates a new read-only instance of L2, bound to a specific deployed contract.
func NewL2Caller(address common.Address, caller bind.ContractCaller) (*L2Caller, error) {
	contract, err := bindL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2Caller{contract: contract}, nil
}

// NewL2Transactor creates a new write-only instance of L2, bound to a specific deployed contract.
func NewL2Transactor(address common.Address, transactor bind.ContractTransactor) (*L2Transactor, error) {
	contract, err := bindL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2Transactor{contract: contract}, nil
}

// NewL2Filterer creates a new log filterer instance of L2, bound to a specific deployed contract.
func NewL2Filterer(address common.Address, filterer bind.ContractFilterer) (*L2Filterer, error) {
	contract, err := bindL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2Filterer{contract: contract}, nil
}

// bindL2 binds a generic wrapper to an already deployed contract.
func bindL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2 *L2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2.Contract.L2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2 *L2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2.Contract.L2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2 *L2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2.Contract.L2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2 *L2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2 *L2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2 *L2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_L2 *L2Caller) BalanceOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_L2 *L2Session) BalanceOf(arg0 *big.Int) (*big.Int, error) {
	return _L2.Contract.BalanceOf(&_L2.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_L2 *L2CallerSession) BalanceOf(arg0 *big.Int) (*big.Int, error) {
	return _L2.Contract.BalanceOf(&_L2.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_L2 *L2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_L2 *L2Session) Decimals() (uint8, error) {
	return _L2.Contract.Decimals(&_L2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_L2 *L2CallerSession) Decimals() (uint8, error) {
	return _L2.Contract.Decimals(&_L2.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L2 *L2Caller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L2 *L2Session) L1Bridge() (common.Address, error) {
	return _L2.Contract.L1Bridge(&_L2.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L2 *L2CallerSession) L1Bridge() (common.Address, error) {
	return _L2.Contract.L1Bridge(&_L2.CallOpts)
}

// L1SharedBridge is a free data retrieval call binding the contract method 0xb852ad36.
//
// Solidity: function l1SharedBridge() view returns(address)
func (_L2 *L2Caller) L1SharedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "l1SharedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1SharedBridge is a free data retrieval call binding the contract method 0xb852ad36.
//
// Solidity: function l1SharedBridge() view returns(address)
func (_L2 *L2Session) L1SharedBridge() (common.Address, error) {
	return _L2.Contract.L1SharedBridge(&_L2.CallOpts)
}

// L1SharedBridge is a free data retrieval call binding the contract method 0xb852ad36.
//
// Solidity: function l1SharedBridge() view returns(address)
func (_L2 *L2CallerSession) L1SharedBridge() (common.Address, error) {
	return _L2.Contract.L1SharedBridge(&_L2.CallOpts)
}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address _l2Token) view returns(address)
func (_L2 *L2Caller) L1TokenAddress(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "l1TokenAddress", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address _l2Token) view returns(address)
func (_L2 *L2Session) L1TokenAddress(_l2Token common.Address) (common.Address, error) {
	return _L2.Contract.L1TokenAddress(&_L2.CallOpts, _l2Token)
}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address _l2Token) view returns(address)
func (_L2 *L2CallerSession) L1TokenAddress(_l2Token common.Address) (common.Address, error) {
	return _L2.Contract.L1TokenAddress(&_L2.CallOpts, _l2Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L2 *L2Caller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L2 *L2Session) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L2.Contract.L2TokenAddress(&_L2.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L2 *L2CallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L2.Contract.L2TokenAddress(&_L2.CallOpts, _l1Token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_L2 *L2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_L2 *L2Session) Name() (string, error) {
	return _L2.Contract.Name(&_L2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_L2 *L2CallerSession) Name() (string, error) {
	return _L2.Contract.Name(&_L2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_L2 *L2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_L2 *L2Session) Symbol() (string, error) {
	return _L2.Contract.Symbol(&_L2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_L2 *L2CallerSession) Symbol() (string, error) {
	return _L2.Contract.Symbol(&_L2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2 *L2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2 *L2Session) TotalSupply() (*big.Int, error) {
	return _L2.Contract.TotalSupply(&_L2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2 *L2CallerSession) TotalSupply() (*big.Int, error) {
	return _L2.Contract.TotalSupply(&_L2.CallOpts)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_L2 *L2Transactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "finalizeDeposit", _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_L2 *L2Session) FinalizeDeposit(_l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2.Contract.FinalizeDeposit(&_L2.TransactOpts, _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_L2 *L2TransactorSession) FinalizeDeposit(_l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2.Contract.FinalizeDeposit(&_L2.TransactOpts, _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_L2 *L2Transactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_L2 *L2Session) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Mint(&_L2.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_L2 *L2TransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Mint(&_L2.TransactOpts, _account, _amount)
}

// TransferFromTo is a paid mutator transaction binding the contract method 0x579952fc.
//
// Solidity: function transferFromTo(address _from, address _to, uint256 _amount) returns()
func (_L2 *L2Transactor) TransferFromTo(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "transferFromTo", _from, _to, _amount)
}

// TransferFromTo is a paid mutator transaction binding the contract method 0x579952fc.
//
// Solidity: function transferFromTo(address _from, address _to, uint256 _amount) returns()
func (_L2 *L2Session) TransferFromTo(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.TransferFromTo(&_L2.TransactOpts, _from, _to, _amount)
}

// TransferFromTo is a paid mutator transaction binding the contract method 0x579952fc.
//
// Solidity: function transferFromTo(address _from, address _to, uint256 _amount) returns()
func (_L2 *L2TransactorSession) TransferFromTo(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.TransferFromTo(&_L2.TransactOpts, _from, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_L2 *L2Transactor) Withdraw(opts *bind.TransactOpts, _l1Receiver common.Address) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "withdraw", _l1Receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_L2 *L2Session) Withdraw(_l1Receiver common.Address) (*types.Transaction, error) {
	return _L2.Contract.Withdraw(&_L2.TransactOpts, _l1Receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_L2 *L2TransactorSession) Withdraw(_l1Receiver common.Address) (*types.Transaction, error) {
	return _L2.Contract.Withdraw(&_L2.TransactOpts, _l1Receiver)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_L2 *L2Transactor) Withdraw0(opts *bind.TransactOpts, _l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "withdraw0", _l1Receiver, _l2Token, _amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_L2 *L2Session) Withdraw0(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Withdraw0(&_L2.TransactOpts, _l1Receiver, _l2Token, _amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_L2 *L2TransactorSession) Withdraw0(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Withdraw0(&_L2.TransactOpts, _l1Receiver, _l2Token, _amount)
}

// WithdrawWithMessage is a paid mutator transaction binding the contract method 0x84bc3eb0.
//
// Solidity: function withdrawWithMessage(address _l1Receiver, bytes _additionalData) payable returns()
func (_L2 *L2Transactor) WithdrawWithMessage(opts *bind.TransactOpts, _l1Receiver common.Address, _additionalData []byte) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "withdrawWithMessage", _l1Receiver, _additionalData)
}

// WithdrawWithMessage is a paid mutator transaction binding the contract method 0x84bc3eb0.
//
// Solidity: function withdrawWithMessage(address _l1Receiver, bytes _additionalData) payable returns()
func (_L2 *L2Session) WithdrawWithMessage(_l1Receiver common.Address, _additionalData []byte) (*types.Transaction, error) {
	return _L2.Contract.WithdrawWithMessage(&_L2.TransactOpts, _l1Receiver, _additionalData)
}

// WithdrawWithMessage is a paid mutator transaction binding the contract method 0x84bc3eb0.
//
// Solidity: function withdrawWithMessage(address _l1Receiver, bytes _additionalData) payable returns()
func (_L2 *L2TransactorSession) WithdrawWithMessage(_l1Receiver common.Address, _additionalData []byte) (*types.Transaction, error) {
	return _L2.Contract.WithdrawWithMessage(&_L2.TransactOpts, _l1Receiver, _additionalData)
}

// L2FinalizeDepositIterator is returned from FilterFinalizeDeposit and is used to iterate over the raw logs and unpacked data for FinalizeDeposit events raised by the L2 contract.
type L2FinalizeDepositIterator struct {
	Event *L2FinalizeDeposit // Event containing the contract specifics and raw log

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
func (it *L2FinalizeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2FinalizeDeposit)
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
		it.Event = new(L2FinalizeDeposit)
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
func (it *L2FinalizeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2FinalizeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2FinalizeDeposit represents a FinalizeDeposit event raised by the L2 contract.
type L2FinalizeDeposit struct {
	L1Sender   common.Address
	L2Receiver common.Address
	L2Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDeposit is a free log retrieval operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_L2 *L2Filterer) FilterFinalizeDeposit(opts *bind.FilterOpts, l1Sender []common.Address, l2Receiver []common.Address, l2Token []common.Address) (*L2FinalizeDepositIterator, error) {

	var l1SenderRule []interface{}
	for _, l1SenderItem := range l1Sender {
		l1SenderRule = append(l1SenderRule, l1SenderItem)
	}
	var l2ReceiverRule []interface{}
	for _, l2ReceiverItem := range l2Receiver {
		l2ReceiverRule = append(l2ReceiverRule, l2ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "FinalizeDeposit", l1SenderRule, l2ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2FinalizeDepositIterator{contract: _L2.contract, event: "FinalizeDeposit", logs: logs, sub: sub}, nil
}

// WatchFinalizeDeposit is a free log subscription operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_L2 *L2Filterer) WatchFinalizeDeposit(opts *bind.WatchOpts, sink chan<- *L2FinalizeDeposit, l1Sender []common.Address, l2Receiver []common.Address, l2Token []common.Address) (event.Subscription, error) {

	var l1SenderRule []interface{}
	for _, l1SenderItem := range l1Sender {
		l1SenderRule = append(l1SenderRule, l1SenderItem)
	}
	var l2ReceiverRule []interface{}
	for _, l2ReceiverItem := range l2Receiver {
		l2ReceiverRule = append(l2ReceiverRule, l2ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "FinalizeDeposit", l1SenderRule, l2ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2FinalizeDeposit)
				if err := _L2.contract.UnpackLog(event, "FinalizeDeposit", log); err != nil {
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

// ParseFinalizeDeposit is a log parse operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_L2 *L2Filterer) ParseFinalizeDeposit(log types.Log) (*L2FinalizeDeposit, error) {
	event := new(L2FinalizeDeposit)
	if err := _L2.contract.UnpackLog(event, "FinalizeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the L2 contract.
type L2MintIterator struct {
	Event *L2Mint // Event containing the contract specifics and raw log

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
func (it *L2MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Mint)
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
		it.Event = new(L2Mint)
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
func (it *L2MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Mint represents a Mint event raised by the L2 contract.
type L2Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_L2 *L2Filterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*L2MintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &L2MintIterator{contract: _L2.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_L2 *L2Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *L2Mint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Mint)
				if err := _L2.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_L2 *L2Filterer) ParseMint(log types.Log) (*L2Mint, error) {
	event := new(L2Mint)
	if err := _L2.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L2 contract.
type L2TransferIterator struct {
	Event *L2Transfer // Event containing the contract specifics and raw log

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
func (it *L2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Transfer)
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
		it.Event = new(L2Transfer)
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
func (it *L2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Transfer represents a Transfer event raised by the L2 contract.
type L2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2 *L2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2TransferIterator{contract: _L2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2 *L2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Transfer)
				if err := _L2.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2 *L2Filterer) ParseTransfer(log types.Log) (*L2Transfer, error) {
	event := new(L2Transfer)
	if err := _L2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the L2 contract.
type L2WithdrawalIterator struct {
	Event *L2Withdrawal // Event containing the contract specifics and raw log

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
func (it *L2WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Withdrawal)
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
		it.Event = new(L2Withdrawal)
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
func (it *L2WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Withdrawal represents a Withdrawal event raised by the L2 contract.
type L2Withdrawal struct {
	L2Sender   common.Address
	L1Receiver common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed _l2Sender, address indexed _l1Receiver, uint256 _amount)
func (_L2 *L2Filterer) FilterWithdrawal(opts *bind.FilterOpts, _l2Sender []common.Address, _l1Receiver []common.Address) (*L2WithdrawalIterator, error) {

	var _l2SenderRule []interface{}
	for _, _l2SenderItem := range _l2Sender {
		_l2SenderRule = append(_l2SenderRule, _l2SenderItem)
	}
	var _l1ReceiverRule []interface{}
	for _, _l1ReceiverItem := range _l1Receiver {
		_l1ReceiverRule = append(_l1ReceiverRule, _l1ReceiverItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Withdrawal", _l2SenderRule, _l1ReceiverRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawalIterator{contract: _L2.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed _l2Sender, address indexed _l1Receiver, uint256 _amount)
func (_L2 *L2Filterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *L2Withdrawal, _l2Sender []common.Address, _l1Receiver []common.Address) (event.Subscription, error) {

	var _l2SenderRule []interface{}
	for _, _l2SenderItem := range _l2Sender {
		_l2SenderRule = append(_l2SenderRule, _l2SenderItem)
	}
	var _l1ReceiverRule []interface{}
	for _, _l1ReceiverItem := range _l1Receiver {
		_l1ReceiverRule = append(_l1ReceiverRule, _l1ReceiverItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Withdrawal", _l2SenderRule, _l1ReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Withdrawal)
				if err := _L2.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed _l2Sender, address indexed _l1Receiver, uint256 _amount)
func (_L2 *L2Filterer) ParseWithdrawal(log types.Log) (*L2Withdrawal, error) {
	event := new(L2Withdrawal)
	if err := _L2.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the L2 contract.
type L2WithdrawalInitiatedIterator struct {
	Event *L2WithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *L2WithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawalInitiated)
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
		it.Event = new(L2WithdrawalInitiated)
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
func (it *L2WithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawalInitiated represents a WithdrawalInitiated event raised by the L2 contract.
type L2WithdrawalInitiated struct {
	L2Sender   common.Address
	L1Receiver common.Address
	L2Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_L2 *L2Filterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, l2Sender []common.Address, l1Receiver []common.Address, l2Token []common.Address) (*L2WithdrawalInitiatedIterator, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var l1ReceiverRule []interface{}
	for _, l1ReceiverItem := range l1Receiver {
		l1ReceiverRule = append(l1ReceiverRule, l1ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "WithdrawalInitiated", l2SenderRule, l1ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawalInitiatedIterator{contract: _L2.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_L2 *L2Filterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *L2WithdrawalInitiated, l2Sender []common.Address, l1Receiver []common.Address, l2Token []common.Address) (event.Subscription, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var l1ReceiverRule []interface{}
	for _, l1ReceiverItem := range l1Receiver {
		l1ReceiverRule = append(l1ReceiverRule, l1ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "WithdrawalInitiated", l2SenderRule, l1ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawalInitiated)
				if err := _L2.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_L2 *L2Filterer) ParseWithdrawalInitiated(log types.Log) (*L2WithdrawalInitiated, error) {
	event := new(L2WithdrawalInitiated)
	if err := _L2.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawalWithMessageIterator is returned from FilterWithdrawalWithMessage and is used to iterate over the raw logs and unpacked data for WithdrawalWithMessage events raised by the L2 contract.
type L2WithdrawalWithMessageIterator struct {
	Event *L2WithdrawalWithMessage // Event containing the contract specifics and raw log

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
func (it *L2WithdrawalWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawalWithMessage)
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
		it.Event = new(L2WithdrawalWithMessage)
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
func (it *L2WithdrawalWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawalWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawalWithMessage represents a WithdrawalWithMessage event raised by the L2 contract.
type L2WithdrawalWithMessage struct {
	L2Sender       common.Address
	L1Receiver     common.Address
	Amount         *big.Int
	AdditionalData []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalWithMessage is a free log retrieval operation binding the contract event 0xc405fe8958410bbaf0c73b7a0c3e20859e86ca168a4c9b0def9c54d2555a306b.
//
// Solidity: event WithdrawalWithMessage(address indexed _l2Sender, address indexed _l1Receiver, uint256 _amount, bytes _additionalData)
func (_L2 *L2Filterer) FilterWithdrawalWithMessage(opts *bind.FilterOpts, _l2Sender []common.Address, _l1Receiver []common.Address) (*L2WithdrawalWithMessageIterator, error) {

	var _l2SenderRule []interface{}
	for _, _l2SenderItem := range _l2Sender {
		_l2SenderRule = append(_l2SenderRule, _l2SenderItem)
	}
	var _l1ReceiverRule []interface{}
	for _, _l1ReceiverItem := range _l1Receiver {
		_l1ReceiverRule = append(_l1ReceiverRule, _l1ReceiverItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "WithdrawalWithMessage", _l2SenderRule, _l1ReceiverRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawalWithMessageIterator{contract: _L2.contract, event: "WithdrawalWithMessage", logs: logs, sub: sub}, nil
}

// WatchWithdrawalWithMessage is a free log subscription operation binding the contract event 0xc405fe8958410bbaf0c73b7a0c3e20859e86ca168a4c9b0def9c54d2555a306b.
//
// Solidity: event WithdrawalWithMessage(address indexed _l2Sender, address indexed _l1Receiver, uint256 _amount, bytes _additionalData)
func (_L2 *L2Filterer) WatchWithdrawalWithMessage(opts *bind.WatchOpts, sink chan<- *L2WithdrawalWithMessage, _l2Sender []common.Address, _l1Receiver []common.Address) (event.Subscription, error) {

	var _l2SenderRule []interface{}
	for _, _l2SenderItem := range _l2Sender {
		_l2SenderRule = append(_l2SenderRule, _l2SenderItem)
	}
	var _l1ReceiverRule []interface{}
	for _, _l1ReceiverItem := range _l1Receiver {
		_l1ReceiverRule = append(_l1ReceiverRule, _l1ReceiverItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "WithdrawalWithMessage", _l2SenderRule, _l1ReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawalWithMessage)
				if err := _L2.contract.UnpackLog(event, "WithdrawalWithMessage", log); err != nil {
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

// ParseWithdrawalWithMessage is a log parse operation binding the contract event 0xc405fe8958410bbaf0c73b7a0c3e20859e86ca168a4c9b0def9c54d2555a306b.
//
// Solidity: event WithdrawalWithMessage(address indexed _l2Sender, address indexed _l1Receiver, uint256 _amount, bytes _additionalData)
func (_L2 *L2Filterer) ParseWithdrawalWithMessage(log types.Log) (*L2WithdrawalWithMessage, error) {
	event := new(L2WithdrawalWithMessage)
	if err := _L2.contract.UnpackLog(event, "WithdrawalWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
