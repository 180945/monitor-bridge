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

// L1ERC20MetaData contains all meta data concerning the L1ERC20 contract.
var L1ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIZkSync\",\"name\":\"_zkSync\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_l2TokenBeacon\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deployBridgeImplementationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deployBridgeProxyFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2ToL1MessageNumber\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenProxyBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC20MetaData.ABI instead.
var L1ERC20ABI = L1ERC20MetaData.ABI

// L1ERC20 is an auto generated Go binding around an Ethereum contract.
type L1ERC20 struct {
	L1ERC20Caller     // Read-only binding to the contract
	L1ERC20Transactor // Write-only binding to the contract
	L1ERC20Filterer   // Log filterer for contract events
}

// L1ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC20Session struct {
	Contract     *L1ERC20          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC20CallerSession struct {
	Contract *L1ERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// L1ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC20TransactorSession struct {
	Contract     *L1ERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// L1ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC20Raw struct {
	Contract *L1ERC20 // Generic contract binding to access the raw methods on
}

// L1ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC20CallerRaw struct {
	Contract *L1ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// L1ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC20TransactorRaw struct {
	Contract *L1ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC20 creates a new instance of L1ERC20, bound to a specific deployed contract.
func NewL1ERC20(address common.Address, backend bind.ContractBackend) (*L1ERC20, error) {
	contract, err := bindL1ERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC20{L1ERC20Caller: L1ERC20Caller{contract: contract}, L1ERC20Transactor: L1ERC20Transactor{contract: contract}, L1ERC20Filterer: L1ERC20Filterer{contract: contract}}, nil
}

// NewL1ERC20Caller creates a new read-only instance of L1ERC20, bound to a specific deployed contract.
func NewL1ERC20Caller(address common.Address, caller bind.ContractCaller) (*L1ERC20Caller, error) {
	contract, err := bindL1ERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Caller{contract: contract}, nil
}

// NewL1ERC20Transactor creates a new write-only instance of L1ERC20, bound to a specific deployed contract.
func NewL1ERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC20Transactor, error) {
	contract, err := bindL1ERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Transactor{contract: contract}, nil
}

// NewL1ERC20Filterer creates a new log filterer instance of L1ERC20, bound to a specific deployed contract.
func NewL1ERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC20Filterer, error) {
	contract, err := bindL1ERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Filterer{contract: contract}, nil
}

// bindL1ERC20 binds a generic wrapper to an already deployed contract.
func bindL1ERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20 *L1ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20.Contract.L1ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20 *L1ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20.Contract.L1ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20 *L1ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20.Contract.L1ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20 *L1ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20 *L1ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20 *L1ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20.Contract.contract.Transact(opts, method, params...)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_L1ERC20 *L1ERC20Caller) IsWithdrawalFinalized(opts *bind.CallOpts, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _L1ERC20.contract.Call(opts, &out, "isWithdrawalFinalized", l2BatchNumber, l2ToL1MessageNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_L1ERC20 *L1ERC20Session) IsWithdrawalFinalized(l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _L1ERC20.Contract.IsWithdrawalFinalized(&_L1ERC20.CallOpts, l2BatchNumber, l2ToL1MessageNumber)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_L1ERC20 *L1ERC20CallerSession) IsWithdrawalFinalized(l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _L1ERC20.Contract.IsWithdrawalFinalized(&_L1ERC20.CallOpts, l2BatchNumber, l2ToL1MessageNumber)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20 *L1ERC20Caller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20 *L1ERC20Session) L2Bridge() (common.Address, error) {
	return _L1ERC20.Contract.L2Bridge(&_L1ERC20.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20 *L1ERC20CallerSession) L2Bridge() (common.Address, error) {
	return _L1ERC20.Contract.L2Bridge(&_L1ERC20.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20 *L1ERC20Caller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20 *L1ERC20Session) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1ERC20.Contract.L2TokenAddress(&_L1ERC20.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20 *L1ERC20CallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1ERC20.Contract.L2TokenAddress(&_L1ERC20.CallOpts, _l1Token)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L1ERC20 *L1ERC20Caller) L2TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20.contract.Call(opts, &out, "l2TokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L1ERC20 *L1ERC20Session) L2TokenBeacon() (common.Address, error) {
	return _L1ERC20.Contract.L2TokenBeacon(&_L1ERC20.CallOpts)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L1ERC20 *L1ERC20CallerSession) L2TokenBeacon() (common.Address, error) {
	return _L1ERC20.Contract.L2TokenBeacon(&_L1ERC20.CallOpts)
}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_L1ERC20 *L1ERC20Caller) L2TokenProxyBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1ERC20.contract.Call(opts, &out, "l2TokenProxyBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_L1ERC20 *L1ERC20Session) L2TokenProxyBytecodeHash() ([32]byte, error) {
	return _L1ERC20.Contract.L2TokenProxyBytecodeHash(&_L1ERC20.CallOpts)
}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_L1ERC20 *L1ERC20CallerSession) L2TokenProxyBytecodeHash() ([32]byte, error) {
	return _L1ERC20.Contract.L2TokenProxyBytecodeHash(&_L1ERC20.CallOpts)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1ERC20 *L1ERC20Transactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1ERC20 *L1ERC20Session) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20.Contract.ClaimFailedDeposit(&_L1ERC20.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1ERC20 *L1ERC20TransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20.Contract.ClaimFailedDeposit(&_L1ERC20.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_L1ERC20 *L1ERC20Transactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _L1ERC20.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_L1ERC20 *L1ERC20Session) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _L1ERC20.Contract.Deposit(&_L1ERC20.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_L1ERC20 *L1ERC20TransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _L1ERC20.Contract.Deposit(&_L1ERC20.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1ERC20 *L1ERC20Transactor) Deposit0(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1ERC20.contract.Transact(opts, "deposit0", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1ERC20 *L1ERC20Session) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1ERC20.Contract.Deposit0(&_L1ERC20.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1ERC20 *L1ERC20TransactorSession) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1ERC20.Contract.Deposit0(&_L1ERC20.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20 *L1ERC20Transactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20.contract.Transact(opts, "finalizeWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20 *L1ERC20Session) FinalizeWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20.Contract.FinalizeWithdrawal(&_L1ERC20.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20 *L1ERC20TransactorSession) FinalizeWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20.Contract.FinalizeWithdrawal(&_L1ERC20.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0473785.
//
// Solidity: function initialize(bytes[] _factoryDeps, address _l2TokenBeacon, address _governor, uint256 _deployBridgeImplementationFee, uint256 _deployBridgeProxyFee) payable returns()
func (_L1ERC20 *L1ERC20Transactor) Initialize(opts *bind.TransactOpts, _factoryDeps [][]byte, _l2TokenBeacon common.Address, _governor common.Address, _deployBridgeImplementationFee *big.Int, _deployBridgeProxyFee *big.Int) (*types.Transaction, error) {
	return _L1ERC20.contract.Transact(opts, "initialize", _factoryDeps, _l2TokenBeacon, _governor, _deployBridgeImplementationFee, _deployBridgeProxyFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0473785.
//
// Solidity: function initialize(bytes[] _factoryDeps, address _l2TokenBeacon, address _governor, uint256 _deployBridgeImplementationFee, uint256 _deployBridgeProxyFee) payable returns()
func (_L1ERC20 *L1ERC20Session) Initialize(_factoryDeps [][]byte, _l2TokenBeacon common.Address, _governor common.Address, _deployBridgeImplementationFee *big.Int, _deployBridgeProxyFee *big.Int) (*types.Transaction, error) {
	return _L1ERC20.Contract.Initialize(&_L1ERC20.TransactOpts, _factoryDeps, _l2TokenBeacon, _governor, _deployBridgeImplementationFee, _deployBridgeProxyFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0473785.
//
// Solidity: function initialize(bytes[] _factoryDeps, address _l2TokenBeacon, address _governor, uint256 _deployBridgeImplementationFee, uint256 _deployBridgeProxyFee) payable returns()
func (_L1ERC20 *L1ERC20TransactorSession) Initialize(_factoryDeps [][]byte, _l2TokenBeacon common.Address, _governor common.Address, _deployBridgeImplementationFee *big.Int, _deployBridgeProxyFee *big.Int) (*types.Transaction, error) {
	return _L1ERC20.Contract.Initialize(&_L1ERC20.TransactOpts, _factoryDeps, _l2TokenBeacon, _governor, _deployBridgeImplementationFee, _deployBridgeProxyFee)
}

// L1ERC20ClaimedFailedDepositIterator is returned from FilterClaimedFailedDeposit and is used to iterate over the raw logs and unpacked data for ClaimedFailedDeposit events raised by the L1ERC20 contract.
type L1ERC20ClaimedFailedDepositIterator struct {
	Event *L1ERC20ClaimedFailedDeposit // Event containing the contract specifics and raw log

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
func (it *L1ERC20ClaimedFailedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20ClaimedFailedDeposit)
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
		it.Event = new(L1ERC20ClaimedFailedDeposit)
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
func (it *L1ERC20ClaimedFailedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20ClaimedFailedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20ClaimedFailedDeposit represents a ClaimedFailedDeposit event raised by the L1ERC20 contract.
type L1ERC20ClaimedFailedDeposit struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDeposit is a free log retrieval operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20 *L1ERC20Filterer) FilterClaimedFailedDeposit(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*L1ERC20ClaimedFailedDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20.contract.FilterLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20ClaimedFailedDepositIterator{contract: _L1ERC20.contract, event: "ClaimedFailedDeposit", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDeposit is a free log subscription operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20 *L1ERC20Filterer) WatchClaimedFailedDeposit(opts *bind.WatchOpts, sink chan<- *L1ERC20ClaimedFailedDeposit, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20.contract.WatchLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20ClaimedFailedDeposit)
				if err := _L1ERC20.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
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
func (_L1ERC20 *L1ERC20Filterer) ParseClaimedFailedDeposit(log types.Log) (*L1ERC20ClaimedFailedDeposit, error) {
	event := new(L1ERC20ClaimedFailedDeposit)
	if err := _L1ERC20.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20DepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the L1ERC20 contract.
type L1ERC20DepositInitiatedIterator struct {
	Event *L1ERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1ERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20DepositInitiated)
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
		it.Event = new(L1ERC20DepositInitiated)
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
func (it *L1ERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20DepositInitiated represents a DepositInitiated event raised by the L1ERC20 contract.
type L1ERC20DepositInitiated struct {
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
func (_L1ERC20 *L1ERC20Filterer) FilterDepositInitiated(opts *bind.FilterOpts, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (*L1ERC20DepositInitiatedIterator, error) {

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

	logs, sub, err := _L1ERC20.contract.FilterLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20DepositInitiatedIterator{contract: _L1ERC20.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_L1ERC20 *L1ERC20Filterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1ERC20DepositInitiated, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1ERC20.contract.WatchLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20DepositInitiated)
				if err := _L1ERC20.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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
func (_L1ERC20 *L1ERC20Filterer) ParseDepositInitiated(log types.Log) (*L1ERC20DepositInitiated, error) {
	event := new(L1ERC20DepositInitiated)
	if err := _L1ERC20.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20WithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the L1ERC20 contract.
type L1ERC20WithdrawalFinalizedIterator struct {
	Event *L1ERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1ERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20WithdrawalFinalized)
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
		it.Event = new(L1ERC20WithdrawalFinalized)
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
func (it *L1ERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20WithdrawalFinalized represents a WithdrawalFinalized event raised by the L1ERC20 contract.
type L1ERC20WithdrawalFinalized struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20 *L1ERC20Filterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*L1ERC20WithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20.contract.FilterLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20WithdrawalFinalizedIterator{contract: _L1ERC20.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20 *L1ERC20Filterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1ERC20WithdrawalFinalized, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20.contract.WatchLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20WithdrawalFinalized)
				if err := _L1ERC20.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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
func (_L1ERC20 *L1ERC20Filterer) ParseWithdrawalFinalized(log types.Log) (*L1ERC20WithdrawalFinalized, error) {
	event := new(L1ERC20WithdrawalFinalized)
	if err := _L1ERC20.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
