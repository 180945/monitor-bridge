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

// BridgeL1MetaData contains all meta data concerning the BridgeL1 contract.
var BridgeL1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1SharedBridge\",\"name\":\"_sharedBridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SHARED_BRIDGE\",\"outputs\":[{\"internalType\":\"contractIL1SharedBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"depositL2TxHash\",\"type\":\"bytes32\"}],\"name\":\"depositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2ToL1MessageNumber\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenProxyBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferTokenToSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeL1MetaData.ABI instead.
var BridgeL1ABI = BridgeL1MetaData.ABI

// BridgeL1 is an auto generated Go binding around an Ethereum contract.
type BridgeL1 struct {
	BridgeL1Caller     // Read-only binding to the contract
	BridgeL1Transactor // Write-only binding to the contract
	BridgeL1Filterer   // Log filterer for contract events
}

// BridgeL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeL1Session struct {
	Contract     *BridgeL1         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeL1CallerSession struct {
	Contract *BridgeL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BridgeL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeL1TransactorSession struct {
	Contract     *BridgeL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BridgeL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeL1Raw struct {
	Contract *BridgeL1 // Generic contract binding to access the raw methods on
}

// BridgeL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeL1CallerRaw struct {
	Contract *BridgeL1Caller // Generic read-only contract binding to access the raw methods on
}

// BridgeL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeL1TransactorRaw struct {
	Contract *BridgeL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeL1 creates a new instance of BridgeL1, bound to a specific deployed contract.
func NewBridgeL1(address common.Address, backend bind.ContractBackend) (*BridgeL1, error) {
	contract, err := bindBridgeL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeL1{BridgeL1Caller: BridgeL1Caller{contract: contract}, BridgeL1Transactor: BridgeL1Transactor{contract: contract}, BridgeL1Filterer: BridgeL1Filterer{contract: contract}}, nil
}

// NewBridgeL1Caller creates a new read-only instance of BridgeL1, bound to a specific deployed contract.
func NewBridgeL1Caller(address common.Address, caller bind.ContractCaller) (*BridgeL1Caller, error) {
	contract, err := bindBridgeL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeL1Caller{contract: contract}, nil
}

// NewBridgeL1Transactor creates a new write-only instance of BridgeL1, bound to a specific deployed contract.
func NewBridgeL1Transactor(address common.Address, transactor bind.ContractTransactor) (*BridgeL1Transactor, error) {
	contract, err := bindBridgeL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeL1Transactor{contract: contract}, nil
}

// NewBridgeL1Filterer creates a new log filterer instance of BridgeL1, bound to a specific deployed contract.
func NewBridgeL1Filterer(address common.Address, filterer bind.ContractFilterer) (*BridgeL1Filterer, error) {
	contract, err := bindBridgeL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeL1Filterer{contract: contract}, nil
}

// bindBridgeL1 binds a generic wrapper to an already deployed contract.
func bindBridgeL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeL1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeL1 *BridgeL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeL1.Contract.BridgeL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeL1 *BridgeL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeL1.Contract.BridgeL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeL1 *BridgeL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeL1.Contract.BridgeL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeL1 *BridgeL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeL1 *BridgeL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeL1 *BridgeL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeL1.Contract.contract.Transact(opts, method, params...)
}

// SHAREDBRIDGE is a free data retrieval call binding the contract method 0x546b6d2a.
//
// Solidity: function SHARED_BRIDGE() view returns(address)
func (_BridgeL1 *BridgeL1Caller) SHAREDBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "SHARED_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SHAREDBRIDGE is a free data retrieval call binding the contract method 0x546b6d2a.
//
// Solidity: function SHARED_BRIDGE() view returns(address)
func (_BridgeL1 *BridgeL1Session) SHAREDBRIDGE() (common.Address, error) {
	return _BridgeL1.Contract.SHAREDBRIDGE(&_BridgeL1.CallOpts)
}

// SHAREDBRIDGE is a free data retrieval call binding the contract method 0x546b6d2a.
//
// Solidity: function SHARED_BRIDGE() view returns(address)
func (_BridgeL1 *BridgeL1CallerSession) SHAREDBRIDGE() (common.Address, error) {
	return _BridgeL1.Contract.SHAREDBRIDGE(&_BridgeL1.CallOpts)
}

// DepositAmount is a free data retrieval call binding the contract method 0x01eae183.
//
// Solidity: function depositAmount(address account, address l1Token, bytes32 depositL2TxHash) view returns(uint256 amount)
func (_BridgeL1 *BridgeL1Caller) DepositAmount(opts *bind.CallOpts, account common.Address, l1Token common.Address, depositL2TxHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "depositAmount", account, l1Token, depositL2TxHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositAmount is a free data retrieval call binding the contract method 0x01eae183.
//
// Solidity: function depositAmount(address account, address l1Token, bytes32 depositL2TxHash) view returns(uint256 amount)
func (_BridgeL1 *BridgeL1Session) DepositAmount(account common.Address, l1Token common.Address, depositL2TxHash [32]byte) (*big.Int, error) {
	return _BridgeL1.Contract.DepositAmount(&_BridgeL1.CallOpts, account, l1Token, depositL2TxHash)
}

// DepositAmount is a free data retrieval call binding the contract method 0x01eae183.
//
// Solidity: function depositAmount(address account, address l1Token, bytes32 depositL2TxHash) view returns(uint256 amount)
func (_BridgeL1 *BridgeL1CallerSession) DepositAmount(account common.Address, l1Token common.Address, depositL2TxHash [32]byte) (*big.Int, error) {
	return _BridgeL1.Contract.DepositAmount(&_BridgeL1.CallOpts, account, l1Token, depositL2TxHash)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_BridgeL1 *BridgeL1Caller) IsWithdrawalFinalized(opts *bind.CallOpts, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "isWithdrawalFinalized", l2BatchNumber, l2ToL1MessageNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_BridgeL1 *BridgeL1Session) IsWithdrawalFinalized(l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _BridgeL1.Contract.IsWithdrawalFinalized(&_BridgeL1.CallOpts, l2BatchNumber, l2ToL1MessageNumber)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_BridgeL1 *BridgeL1CallerSession) IsWithdrawalFinalized(l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _BridgeL1.Contract.IsWithdrawalFinalized(&_BridgeL1.CallOpts, l2BatchNumber, l2ToL1MessageNumber)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_BridgeL1 *BridgeL1Caller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_BridgeL1 *BridgeL1Session) L2Bridge() (common.Address, error) {
	return _BridgeL1.Contract.L2Bridge(&_BridgeL1.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_BridgeL1 *BridgeL1CallerSession) L2Bridge() (common.Address, error) {
	return _BridgeL1.Contract.L2Bridge(&_BridgeL1.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_BridgeL1 *BridgeL1Caller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_BridgeL1 *BridgeL1Session) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _BridgeL1.Contract.L2TokenAddress(&_BridgeL1.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_BridgeL1 *BridgeL1CallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _BridgeL1.Contract.L2TokenAddress(&_BridgeL1.CallOpts, _l1Token)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_BridgeL1 *BridgeL1Caller) L2TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "l2TokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_BridgeL1 *BridgeL1Session) L2TokenBeacon() (common.Address, error) {
	return _BridgeL1.Contract.L2TokenBeacon(&_BridgeL1.CallOpts)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_BridgeL1 *BridgeL1CallerSession) L2TokenBeacon() (common.Address, error) {
	return _BridgeL1.Contract.L2TokenBeacon(&_BridgeL1.CallOpts)
}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_BridgeL1 *BridgeL1Caller) L2TokenProxyBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeL1.contract.Call(opts, &out, "l2TokenProxyBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_BridgeL1 *BridgeL1Session) L2TokenProxyBytecodeHash() ([32]byte, error) {
	return _BridgeL1.Contract.L2TokenProxyBytecodeHash(&_BridgeL1.CallOpts)
}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_BridgeL1 *BridgeL1CallerSession) L2TokenProxyBytecodeHash() ([32]byte, error) {
	return _BridgeL1.Contract.L2TokenProxyBytecodeHash(&_BridgeL1.CallOpts)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_BridgeL1 *BridgeL1Transactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _BridgeL1.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_BridgeL1 *BridgeL1Session) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _BridgeL1.Contract.ClaimFailedDeposit(&_BridgeL1.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_BridgeL1 *BridgeL1TransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _BridgeL1.Contract.ClaimFailedDeposit(&_BridgeL1.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_BridgeL1 *BridgeL1Transactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _BridgeL1.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_BridgeL1 *BridgeL1Session) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _BridgeL1.Contract.Deposit(&_BridgeL1.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_BridgeL1 *BridgeL1TransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _BridgeL1.Contract.Deposit(&_BridgeL1.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_BridgeL1 *BridgeL1Transactor) Deposit0(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _BridgeL1.contract.Transact(opts, "deposit0", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_BridgeL1 *BridgeL1Session) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _BridgeL1.Contract.Deposit0(&_BridgeL1.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_BridgeL1 *BridgeL1TransactorSession) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _BridgeL1.Contract.Deposit0(&_BridgeL1.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_BridgeL1 *BridgeL1Transactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _BridgeL1.contract.Transact(opts, "finalizeWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_BridgeL1 *BridgeL1Session) FinalizeWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _BridgeL1.Contract.FinalizeWithdrawal(&_BridgeL1.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_BridgeL1 *BridgeL1TransactorSession) FinalizeWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _BridgeL1.Contract.FinalizeWithdrawal(&_BridgeL1.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BridgeL1 *BridgeL1Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeL1.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BridgeL1 *BridgeL1Session) Initialize() (*types.Transaction, error) {
	return _BridgeL1.Contract.Initialize(&_BridgeL1.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BridgeL1 *BridgeL1TransactorSession) Initialize() (*types.Transaction, error) {
	return _BridgeL1.Contract.Initialize(&_BridgeL1.TransactOpts)
}

// TransferTokenToSharedBridge is a paid mutator transaction binding the contract method 0xb7b080ab.
//
// Solidity: function transferTokenToSharedBridge(address _token) returns()
func (_BridgeL1 *BridgeL1Transactor) TransferTokenToSharedBridge(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _BridgeL1.contract.Transact(opts, "transferTokenToSharedBridge", _token)
}

// TransferTokenToSharedBridge is a paid mutator transaction binding the contract method 0xb7b080ab.
//
// Solidity: function transferTokenToSharedBridge(address _token) returns()
func (_BridgeL1 *BridgeL1Session) TransferTokenToSharedBridge(_token common.Address) (*types.Transaction, error) {
	return _BridgeL1.Contract.TransferTokenToSharedBridge(&_BridgeL1.TransactOpts, _token)
}

// TransferTokenToSharedBridge is a paid mutator transaction binding the contract method 0xb7b080ab.
//
// Solidity: function transferTokenToSharedBridge(address _token) returns()
func (_BridgeL1 *BridgeL1TransactorSession) TransferTokenToSharedBridge(_token common.Address) (*types.Transaction, error) {
	return _BridgeL1.Contract.TransferTokenToSharedBridge(&_BridgeL1.TransactOpts, _token)
}

// BridgeL1ClaimedFailedDepositIterator is returned from FilterClaimedFailedDeposit and is used to iterate over the raw logs and unpacked data for ClaimedFailedDeposit events raised by the BridgeL1 contract.
type BridgeL1ClaimedFailedDepositIterator struct {
	Event *BridgeL1ClaimedFailedDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeL1ClaimedFailedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeL1ClaimedFailedDeposit)
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
		it.Event = new(BridgeL1ClaimedFailedDeposit)
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
func (it *BridgeL1ClaimedFailedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeL1ClaimedFailedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeL1ClaimedFailedDeposit represents a ClaimedFailedDeposit event raised by the BridgeL1 contract.
type BridgeL1ClaimedFailedDeposit struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDeposit is a free log retrieval operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) FilterClaimedFailedDeposit(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*BridgeL1ClaimedFailedDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _BridgeL1.contract.FilterLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeL1ClaimedFailedDepositIterator{contract: _BridgeL1.contract, event: "ClaimedFailedDeposit", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDeposit is a free log subscription operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) WatchClaimedFailedDeposit(opts *bind.WatchOpts, sink chan<- *BridgeL1ClaimedFailedDeposit, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _BridgeL1.contract.WatchLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeL1ClaimedFailedDeposit)
				if err := _BridgeL1.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
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

// ParseClaimedFailedDeposit is a log parse operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) ParseClaimedFailedDeposit(log types.Log) (*BridgeL1ClaimedFailedDeposit, error) {
	event := new(BridgeL1ClaimedFailedDeposit)
	if err := _BridgeL1.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeL1DepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the BridgeL1 contract.
type BridgeL1DepositInitiatedIterator struct {
	Event *BridgeL1DepositInitiated // Event containing the contract specifics and raw log

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
func (it *BridgeL1DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeL1DepositInitiated)
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
		it.Event = new(BridgeL1DepositInitiated)
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
func (it *BridgeL1DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeL1DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeL1DepositInitiated represents a DepositInitiated event raised by the BridgeL1 contract.
type BridgeL1DepositInitiated struct {
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiated is a free log retrieval operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) FilterDepositInitiated(opts *bind.FilterOpts, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (*BridgeL1DepositInitiatedIterator, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeL1.contract.FilterLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeL1DepositInitiatedIterator{contract: _BridgeL1.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *BridgeL1DepositInitiated, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeL1.contract.WatchLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeL1DepositInitiated)
				if err := _BridgeL1.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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

// ParseDepositInitiated is a log parse operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) ParseDepositInitiated(log types.Log) (*BridgeL1DepositInitiated, error) {
	event := new(BridgeL1DepositInitiated)
	if err := _BridgeL1.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeL1WithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the BridgeL1 contract.
type BridgeL1WithdrawalFinalizedIterator struct {
	Event *BridgeL1WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *BridgeL1WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeL1WithdrawalFinalized)
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
		it.Event = new(BridgeL1WithdrawalFinalized)
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
func (it *BridgeL1WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeL1WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeL1WithdrawalFinalized represents a WithdrawalFinalized event raised by the BridgeL1 contract.
type BridgeL1WithdrawalFinalized struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*BridgeL1WithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _BridgeL1.contract.FilterLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeL1WithdrawalFinalizedIterator{contract: _BridgeL1.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *BridgeL1WithdrawalFinalized, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _BridgeL1.contract.WatchLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeL1WithdrawalFinalized)
				if err := _BridgeL1.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_BridgeL1 *BridgeL1Filterer) ParseWithdrawalFinalized(log types.Log) (*BridgeL1WithdrawalFinalized, error) {
	event := new(BridgeL1WithdrawalFinalized)
	if err := _BridgeL1.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
