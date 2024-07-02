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

// IMailboxL2CanonicalTransaction is an auto generated low-level Go binding around an user-defined struct.
type IMailboxL2CanonicalTransaction struct {
	TxType                 *big.Int
	From                   *big.Int
	To                     *big.Int
	GasLimit               *big.Int
	GasPerPubdataByteLimit *big.Int
	MaxFeePerGas           *big.Int
	MaxPriorityFeePerGas   *big.Int
	Paymaster              *big.Int
	Nonce                  *big.Int
	Value                  *big.Int
	Reserved               [4]*big.Int
	Data                   []byte
	Signature              []byte
	FactoryDeps            []*big.Int
	PaymasterInput         []byte
	ReservedDynamic        []byte
}

// L2Log is an auto generated low-level Go binding around an user-defined struct.
type L2Log struct {
	L2ShardId       uint8
	IsService       bool
	TxNumberInBatch uint16
	Sender          common.Address
	Key             [32]byte
	Value           [32]byte
}

// L2Message is an auto generated low-level Go binding around an user-defined struct.
type L2Message struct {
	TxNumberInBatch uint16
	Sender          common.Address
	Data            []byte
}

// L1NativeMetaData contains all meta data concerning the L1Native contract.
var L1NativeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"factoryDeps\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIMailbox.L2CanonicalTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeEthWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"}],\"name\":\"l2TransactionBaseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"proveL1ToL2TransactionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structL2Log\",\"name\":\"_log\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2LogInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structL2Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2MessageInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractL2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"requestL2Transaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// L1NativeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1NativeMetaData.ABI instead.
var L1NativeABI = L1NativeMetaData.ABI

// L1Native is an auto generated Go binding around an Ethereum contract.
type L1Native struct {
	L1NativeCaller     // Read-only binding to the contract
	L1NativeTransactor // Write-only binding to the contract
	L1NativeFilterer   // Log filterer for contract events
}

// L1NativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1NativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1NativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1NativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1NativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1NativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1NativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1NativeSession struct {
	Contract     *L1Native         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1NativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1NativeCallerSession struct {
	Contract *L1NativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// L1NativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1NativeTransactorSession struct {
	Contract     *L1NativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// L1NativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1NativeRaw struct {
	Contract *L1Native // Generic contract binding to access the raw methods on
}

// L1NativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1NativeCallerRaw struct {
	Contract *L1NativeCaller // Generic read-only contract binding to access the raw methods on
}

// L1NativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1NativeTransactorRaw struct {
	Contract *L1NativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1Native creates a new instance of L1Native, bound to a specific deployed contract.
func NewL1Native(address common.Address, backend bind.ContractBackend) (*L1Native, error) {
	contract, err := bindL1Native(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1Native{L1NativeCaller: L1NativeCaller{contract: contract}, L1NativeTransactor: L1NativeTransactor{contract: contract}, L1NativeFilterer: L1NativeFilterer{contract: contract}}, nil
}

// NewL1NativeCaller creates a new read-only instance of L1Native, bound to a specific deployed contract.
func NewL1NativeCaller(address common.Address, caller bind.ContractCaller) (*L1NativeCaller, error) {
	contract, err := bindL1Native(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1NativeCaller{contract: contract}, nil
}

// NewL1NativeTransactor creates a new write-only instance of L1Native, bound to a specific deployed contract.
func NewL1NativeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1NativeTransactor, error) {
	contract, err := bindL1Native(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1NativeTransactor{contract: contract}, nil
}

// NewL1NativeFilterer creates a new log filterer instance of L1Native, bound to a specific deployed contract.
func NewL1NativeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1NativeFilterer, error) {
	contract, err := bindL1Native(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1NativeFilterer{contract: contract}, nil
}

// bindL1Native binds a generic wrapper to an already deployed contract.
func bindL1Native(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1NativeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Native *L1NativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Native.Contract.L1NativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Native *L1NativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Native.Contract.L1NativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Native *L1NativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Native.Contract.L1NativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Native *L1NativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Native.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Native *L1NativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Native.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Native *L1NativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Native.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_L1Native *L1NativeCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1Native.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_L1Native *L1NativeSession) GetName() (string, error) {
	return _L1Native.Contract.GetName(&_L1Native.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_L1Native *L1NativeCallerSession) GetName() (string, error) {
	return _L1Native.Contract.GetName(&_L1Native.CallOpts)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_L1Native *L1NativeCaller) L2TransactionBaseCost(opts *bind.CallOpts, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1Native.contract.Call(opts, &out, "l2TransactionBaseCost", _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_L1Native *L1NativeSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _L1Native.Contract.L2TransactionBaseCost(&_L1Native.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_L1Native *L1NativeCallerSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _L1Native.Contract.L2TransactionBaseCost(&_L1Native.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_L1Native *L1NativeCaller) ProveL1ToL2TransactionStatus(opts *bind.CallOpts, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	var out []interface{}
	err := _L1Native.contract.Call(opts, &out, "proveL1ToL2TransactionStatus", _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_L1Native *L1NativeSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _L1Native.Contract.ProveL1ToL2TransactionStatus(&_L1Native.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_L1Native *L1NativeCallerSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _L1Native.Contract.ProveL1ToL2TransactionStatus(&_L1Native.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_L1Native *L1NativeCaller) ProveL2LogInclusion(opts *bind.CallOpts, _batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _L1Native.contract.Call(opts, &out, "proveL2LogInclusion", _batchNumber, _index, _log, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_L1Native *L1NativeSession) ProveL2LogInclusion(_batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _L1Native.Contract.ProveL2LogInclusion(&_L1Native.CallOpts, _batchNumber, _index, _log, _proof)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_L1Native *L1NativeCallerSession) ProveL2LogInclusion(_batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _L1Native.Contract.ProveL2LogInclusion(&_L1Native.CallOpts, _batchNumber, _index, _log, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_L1Native *L1NativeCaller) ProveL2MessageInclusion(opts *bind.CallOpts, _batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _L1Native.contract.Call(opts, &out, "proveL2MessageInclusion", _batchNumber, _index, _message, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_L1Native *L1NativeSession) ProveL2MessageInclusion(_batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _L1Native.Contract.ProveL2MessageInclusion(&_L1Native.CallOpts, _batchNumber, _index, _message, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_L1Native *L1NativeCallerSession) ProveL2MessageInclusion(_batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _L1Native.Contract.ProveL2MessageInclusion(&_L1Native.CallOpts, _batchNumber, _index, _message, _proof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1Native *L1NativeTransactor) FinalizeEthWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1Native.contract.Transact(opts, "finalizeEthWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1Native *L1NativeSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1Native.Contract.FinalizeEthWithdrawal(&_L1Native.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1Native *L1NativeTransactorSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1Native.Contract.FinalizeEthWithdrawal(&_L1Native.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_L1Native *L1NativeTransactor) RequestL2Transaction(opts *bind.TransactOpts, _contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1Native.contract.Transact(opts, "requestL2Transaction", _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_L1Native *L1NativeSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1Native.Contract.RequestL2Transaction(&_L1Native.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_L1Native *L1NativeTransactorSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1Native.Contract.RequestL2Transaction(&_L1Native.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// L1NativeEthWithdrawalFinalizedIterator is returned from FilterEthWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for EthWithdrawalFinalized events raised by the L1Native contract.
type L1NativeEthWithdrawalFinalizedIterator struct {
	Event *L1NativeEthWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1NativeEthWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1NativeEthWithdrawalFinalized)
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
		it.Event = new(L1NativeEthWithdrawalFinalized)
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
func (it *L1NativeEthWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1NativeEthWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1NativeEthWithdrawalFinalized represents a EthWithdrawalFinalized event raised by the L1Native contract.
type L1NativeEthWithdrawalFinalized struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawalFinalized is a free log retrieval operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_L1Native *L1NativeFilterer) FilterEthWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address) (*L1NativeEthWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1Native.contract.FilterLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return &L1NativeEthWithdrawalFinalizedIterator{contract: _L1Native.contract, event: "EthWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawalFinalized is a free log subscription operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_L1Native *L1NativeFilterer) WatchEthWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1NativeEthWithdrawalFinalized, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1Native.contract.WatchLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1NativeEthWithdrawalFinalized)
				if err := _L1Native.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
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

// ParseEthWithdrawalFinalized is a log parse operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_L1Native *L1NativeFilterer) ParseEthWithdrawalFinalized(log types.Log) (*L1NativeEthWithdrawalFinalized, error) {
	event := new(L1NativeEthWithdrawalFinalized)
	if err := _L1Native.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1NativeNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the L1Native contract.
type L1NativeNewPriorityRequestIterator struct {
	Event *L1NativeNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *L1NativeNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1NativeNewPriorityRequest)
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
		it.Event = new(L1NativeNewPriorityRequest)
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
func (it *L1NativeNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1NativeNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1NativeNewPriorityRequest represents a NewPriorityRequest event raised by the L1Native contract.
type L1NativeNewPriorityRequest struct {
	TxId                *big.Int
	TxHash              [32]byte
	ExpirationTimestamp uint64
	Transaction         IMailboxL2CanonicalTransaction
	FactoryDeps         [][]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_L1Native *L1NativeFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*L1NativeNewPriorityRequestIterator, error) {

	logs, sub, err := _L1Native.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &L1NativeNewPriorityRequestIterator{contract: _L1Native.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_L1Native *L1NativeFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *L1NativeNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _L1Native.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1NativeNewPriorityRequest)
				if err := _L1Native.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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

// ParseNewPriorityRequest is a log parse operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_L1Native *L1NativeFilterer) ParseNewPriorityRequest(log types.Log) (*L1NativeNewPriorityRequest, error) {
	event := new(L1NativeNewPriorityRequest)
	if err := _L1Native.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
