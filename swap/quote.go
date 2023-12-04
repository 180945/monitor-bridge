// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// QuoteMetaData contains all meta data concerning the Quote contract.
var QuoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QuoteABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoteMetaData.ABI instead.
var QuoteABI = QuoteMetaData.ABI

// Quote is an auto generated Go binding around an Ethereum contract.
type Quote struct {
	QuoteCaller     // Read-only binding to the contract
	QuoteTransactor // Write-only binding to the contract
	QuoteFilterer   // Log filterer for contract events
}

// QuoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoteSession struct {
	Contract     *Quote            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoteCallerSession struct {
	Contract *QuoteCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoteTransactorSession struct {
	Contract     *QuoteTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoteRaw struct {
	Contract *Quote // Generic contract binding to access the raw methods on
}

// QuoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoteCallerRaw struct {
	Contract *QuoteCaller // Generic read-only contract binding to access the raw methods on
}

// QuoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoteTransactorRaw struct {
	Contract *QuoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuote creates a new instance of Quote, bound to a specific deployed contract.
func NewQuote(address common.Address, backend bind.ContractBackend) (*Quote, error) {
	contract, err := bindQuote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quote{QuoteCaller: QuoteCaller{contract: contract}, QuoteTransactor: QuoteTransactor{contract: contract}, QuoteFilterer: QuoteFilterer{contract: contract}}, nil
}

// NewQuoteCaller creates a new read-only instance of Quote, bound to a specific deployed contract.
func NewQuoteCaller(address common.Address, caller bind.ContractCaller) (*QuoteCaller, error) {
	contract, err := bindQuote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoteCaller{contract: contract}, nil
}

// NewQuoteTransactor creates a new write-only instance of Quote, bound to a specific deployed contract.
func NewQuoteTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoteTransactor, error) {
	contract, err := bindQuote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoteTransactor{contract: contract}, nil
}

// NewQuoteFilterer creates a new log filterer instance of Quote, bound to a specific deployed contract.
func NewQuoteFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoteFilterer, error) {
	contract, err := bindQuote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoteFilterer{contract: contract}, nil
}

// bindQuote binds a generic wrapper to an already deployed contract.
func bindQuote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quote *QuoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quote.Contract.QuoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quote *QuoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quote.Contract.QuoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quote *QuoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quote.Contract.QuoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quote *QuoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quote *QuoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quote *QuoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quote.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quote *QuoteCaller) QuoteExactInput(opts *bind.CallOpts, path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	var out []interface{}
	err := _Quote.contract.Call(opts, &out, "quoteExactInput", path, amountIn)

	outstruct := new(struct {
		AmountOut                   *big.Int
		SqrtPriceX96AfterList       []*big.Int
		InitializedTicksCrossedList []uint32
		GasEstimate                 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96AfterList = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.InitializedTicksCrossedList = *abi.ConvertType(out[2], new([]uint32)).(*[]uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quote *QuoteSession) QuoteExactInput(path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _Quote.Contract.QuoteExactInput(&_Quote.CallOpts, path, amountIn)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quote *QuoteCallerSession) QuoteExactInput(path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _Quote.Contract.QuoteExactInput(&_Quote.CallOpts, path, amountIn)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quote *QuoteCaller) QuoteExactInputSingle(opts *bind.CallOpts, params IQuoterV2QuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	var out []interface{}
	err := _Quote.contract.Call(opts, &out, "quoteExactInputSingle", params)

	outstruct := new(struct {
		AmountOut               *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96After = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitializedTicksCrossed = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quote *QuoteSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quote.Contract.QuoteExactInputSingle(&_Quote.CallOpts, params)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quote *QuoteCallerSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quote.Contract.QuoteExactInputSingle(&_Quote.CallOpts, params)
}

// QuoteExactOutput is a free data retrieval call binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) view returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quote *QuoteCaller) QuoteExactOutput(opts *bind.CallOpts, path []byte, amountOut *big.Int) (struct {
	AmountIn                    *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	var out []interface{}
	err := _Quote.contract.Call(opts, &out, "quoteExactOutput", path, amountOut)

	outstruct := new(struct {
		AmountIn                    *big.Int
		SqrtPriceX96AfterList       []*big.Int
		InitializedTicksCrossedList []uint32
		GasEstimate                 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96AfterList = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.InitializedTicksCrossedList = *abi.ConvertType(out[2], new([]uint32)).(*[]uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactOutput is a free data retrieval call binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) view returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quote *QuoteSession) QuoteExactOutput(path []byte, amountOut *big.Int) (struct {
	AmountIn                    *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _Quote.Contract.QuoteExactOutput(&_Quote.CallOpts, path, amountOut)
}

// QuoteExactOutput is a free data retrieval call binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) view returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quote *QuoteCallerSession) QuoteExactOutput(path []byte, amountOut *big.Int) (struct {
	AmountIn                    *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _Quote.Contract.QuoteExactOutput(&_Quote.CallOpts, path, amountOut)
}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quote *QuoteCaller) QuoteExactOutputSingle(opts *bind.CallOpts, params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	var out []interface{}
	err := _Quote.contract.Call(opts, &out, "quoteExactOutputSingle", params)

	outstruct := new(struct {
		AmountIn                *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96After = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitializedTicksCrossed = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quote *QuoteSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quote.Contract.QuoteExactOutputSingle(&_Quote.CallOpts, params)
}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quote *QuoteCallerSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quote.Contract.QuoteExactOutputSingle(&_Quote.CallOpts, params)
}
