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

// L2TransactionRequestTwoBridgesInner is an auto generated low-level Go binding around an user-defined struct.
type L2TransactionRequestTwoBridgesInner struct {
	MagicValue  [32]byte
	L2Contract  common.Address
	L2Calldata  []byte
	FactoryDeps [][]byte
	TxDataHash  [32]byte
}

// L1MetaData contains all meta data concerning the L1 contract.
var L1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1WethAddress\",\"type\":\"address\"},{\"internalType\":\"contractIBridgehub\",\"name\":\"_bridgehub\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_eraChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_eraDiamondProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositBaseTokenInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubDepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDepositSharedBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LegacyDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalizedSharedBridge\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_HUB\",\"outputs\":[{\"internalType\":\"contractIBridgehub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERA_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERA_DIAMOND_PROXY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_WETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"bridgehubConfirmL2Transaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"bridgehubDeposit\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"magicValue\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"l2Calldata\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structL2TransactionRequestTwoBridgesInner\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bridgehubDepositBaseToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"chainBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDepositLegacyErc20Bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"}],\"name\":\"depositHappened\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"depositDataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"depositLegacyErc20Bridge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawalLegacyErc20Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_l2BridgeAddress\",\"type\":\"address\"}],\"name\":\"initializeChainGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2ToL1MessageNumber\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"l2BridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l2Bridge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"legacyBridge\",\"outputs\":[{\"internalType\":\"contractIL1ERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"receiveEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPerToken\",\"type\":\"uint256\"}],\"name\":\"safeTransferFundsFromLegacy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eraLegacyBridgeLastDepositBatch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eraLegacyBridgeLastDepositTxNumber\",\"type\":\"uint256\"}],\"name\":\"setEraLegacyBridgeLastDepositTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eraPostDiamondUpgradeFirstBatch\",\"type\":\"uint256\"}],\"name\":\"setEraPostDiamondUpgradeFirstBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eraPostLegacyBridgeUpgradeFirstBatch\",\"type\":\"uint256\"}],\"name\":\"setEraPostLegacyBridgeUpgradeFirstBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_legacyBridge\",\"type\":\"address\"}],\"name\":\"setL1Erc20Bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetChainId\",\"type\":\"uint256\"}],\"name\":\"transferFundsFromLegacy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L1ABI is the input ABI used to generate the binding from.
// Deprecated: Use L1MetaData.ABI instead.
var L1ABI = L1MetaData.ABI

// L1 is an auto generated Go binding around an Ethereum contract.
type L1 struct {
	L1Caller     // Read-only binding to the contract
	L1Transactor // Write-only binding to the contract
	L1Filterer   // Log filterer for contract events
}

// L1Caller is an auto generated read-only Go binding around an Ethereum contract.
type L1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1Session struct {
	Contract     *L1               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1CallerSession struct {
	Contract *L1Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1TransactorSession struct {
	Contract     *L1Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1Raw is an auto generated low-level Go binding around an Ethereum contract.
type L1Raw struct {
	Contract *L1 // Generic contract binding to access the raw methods on
}

// L1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1CallerRaw struct {
	Contract *L1Caller // Generic read-only contract binding to access the raw methods on
}

// L1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1TransactorRaw struct {
	Contract *L1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL1 creates a new instance of L1, bound to a specific deployed contract.
func NewL1(address common.Address, backend bind.ContractBackend) (*L1, error) {
	contract, err := bindL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1{L1Caller: L1Caller{contract: contract}, L1Transactor: L1Transactor{contract: contract}, L1Filterer: L1Filterer{contract: contract}}, nil
}

// NewL1Caller creates a new read-only instance of L1, bound to a specific deployed contract.
func NewL1Caller(address common.Address, caller bind.ContractCaller) (*L1Caller, error) {
	contract, err := bindL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1Caller{contract: contract}, nil
}

// NewL1Transactor creates a new write-only instance of L1, bound to a specific deployed contract.
func NewL1Transactor(address common.Address, transactor bind.ContractTransactor) (*L1Transactor, error) {
	contract, err := bindL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1Transactor{contract: contract}, nil
}

// NewL1Filterer creates a new log filterer instance of L1, bound to a specific deployed contract.
func NewL1Filterer(address common.Address, filterer bind.ContractFilterer) (*L1Filterer, error) {
	contract, err := bindL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1Filterer{contract: contract}, nil
}

// bindL1 binds a generic wrapper to an already deployed contract.
func bindL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1 *L1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1.Contract.L1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1 *L1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.Contract.L1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1 *L1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1.Contract.L1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1 *L1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1 *L1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1 *L1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_L1 *L1Caller) BRIDGEHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "BRIDGE_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_L1 *L1Session) BRIDGEHUB() (common.Address, error) {
	return _L1.Contract.BRIDGEHUB(&_L1.CallOpts)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_L1 *L1CallerSession) BRIDGEHUB() (common.Address, error) {
	return _L1.Contract.BRIDGEHUB(&_L1.CallOpts)
}

// ERACHAINID is a free data retrieval call binding the contract method 0xef011dff.
//
// Solidity: function ERA_CHAIN_ID() view returns(uint256)
func (_L1 *L1Caller) ERACHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "ERA_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERACHAINID is a free data retrieval call binding the contract method 0xef011dff.
//
// Solidity: function ERA_CHAIN_ID() view returns(uint256)
func (_L1 *L1Session) ERACHAINID() (*big.Int, error) {
	return _L1.Contract.ERACHAINID(&_L1.CallOpts)
}

// ERACHAINID is a free data retrieval call binding the contract method 0xef011dff.
//
// Solidity: function ERA_CHAIN_ID() view returns(uint256)
func (_L1 *L1CallerSession) ERACHAINID() (*big.Int, error) {
	return _L1.Contract.ERACHAINID(&_L1.CallOpts)
}

// ERADIAMONDPROXY is a free data retrieval call binding the contract method 0x246a61de.
//
// Solidity: function ERA_DIAMOND_PROXY() view returns(address)
func (_L1 *L1Caller) ERADIAMONDPROXY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "ERA_DIAMOND_PROXY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERADIAMONDPROXY is a free data retrieval call binding the contract method 0x246a61de.
//
// Solidity: function ERA_DIAMOND_PROXY() view returns(address)
func (_L1 *L1Session) ERADIAMONDPROXY() (common.Address, error) {
	return _L1.Contract.ERADIAMONDPROXY(&_L1.CallOpts)
}

// ERADIAMONDPROXY is a free data retrieval call binding the contract method 0x246a61de.
//
// Solidity: function ERA_DIAMOND_PROXY() view returns(address)
func (_L1 *L1CallerSession) ERADIAMONDPROXY() (common.Address, error) {
	return _L1.Contract.ERADIAMONDPROXY(&_L1.CallOpts)
}

// L1WETHTOKEN is a free data retrieval call binding the contract method 0x41c841c3.
//
// Solidity: function L1_WETH_TOKEN() view returns(address)
func (_L1 *L1Caller) L1WETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "L1_WETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1WETHTOKEN is a free data retrieval call binding the contract method 0x41c841c3.
//
// Solidity: function L1_WETH_TOKEN() view returns(address)
func (_L1 *L1Session) L1WETHTOKEN() (common.Address, error) {
	return _L1.Contract.L1WETHTOKEN(&_L1.CallOpts)
}

// L1WETHTOKEN is a free data retrieval call binding the contract method 0x41c841c3.
//
// Solidity: function L1_WETH_TOKEN() view returns(address)
func (_L1 *L1CallerSession) L1WETHTOKEN() (common.Address, error) {
	return _L1.Contract.L1WETHTOKEN(&_L1.CallOpts)
}

// ChainBalance is a free data retrieval call binding the contract method 0x9cd45184.
//
// Solidity: function chainBalance(uint256 chainId, address l1Token) view returns(uint256 balance)
func (_L1 *L1Caller) ChainBalance(opts *bind.CallOpts, chainId *big.Int, l1Token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "chainBalance", chainId, l1Token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainBalance is a free data retrieval call binding the contract method 0x9cd45184.
//
// Solidity: function chainBalance(uint256 chainId, address l1Token) view returns(uint256 balance)
func (_L1 *L1Session) ChainBalance(chainId *big.Int, l1Token common.Address) (*big.Int, error) {
	return _L1.Contract.ChainBalance(&_L1.CallOpts, chainId, l1Token)
}

// ChainBalance is a free data retrieval call binding the contract method 0x9cd45184.
//
// Solidity: function chainBalance(uint256 chainId, address l1Token) view returns(uint256 balance)
func (_L1 *L1CallerSession) ChainBalance(chainId *big.Int, l1Token common.Address) (*big.Int, error) {
	return _L1.Contract.ChainBalance(&_L1.CallOpts, chainId, l1Token)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 chainId, bytes32 l2DepositTxHash) view returns(bytes32 depositDataHash)
func (_L1 *L1Caller) DepositHappened(opts *bind.CallOpts, chainId *big.Int, l2DepositTxHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "depositHappened", chainId, l2DepositTxHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 chainId, bytes32 l2DepositTxHash) view returns(bytes32 depositDataHash)
func (_L1 *L1Session) DepositHappened(chainId *big.Int, l2DepositTxHash [32]byte) ([32]byte, error) {
	return _L1.Contract.DepositHappened(&_L1.CallOpts, chainId, l2DepositTxHash)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 chainId, bytes32 l2DepositTxHash) view returns(bytes32 depositDataHash)
func (_L1 *L1CallerSession) DepositHappened(chainId *big.Int, l2DepositTxHash [32]byte) ([32]byte, error) {
	return _L1.Contract.DepositHappened(&_L1.CallOpts, chainId, l2DepositTxHash)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 chainId, uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_L1 *L1Caller) IsWithdrawalFinalized(opts *bind.CallOpts, chainId *big.Int, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "isWithdrawalFinalized", chainId, l2BatchNumber, l2ToL1MessageNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 chainId, uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_L1 *L1Session) IsWithdrawalFinalized(chainId *big.Int, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _L1.Contract.IsWithdrawalFinalized(&_L1.CallOpts, chainId, l2BatchNumber, l2ToL1MessageNumber)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 chainId, uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_L1 *L1CallerSession) IsWithdrawalFinalized(chainId *big.Int, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _L1.Contract.IsWithdrawalFinalized(&_L1.CallOpts, chainId, l2BatchNumber, l2ToL1MessageNumber)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 chainId) view returns(address l2Bridge)
func (_L1 *L1Caller) L2BridgeAddress(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "l2BridgeAddress", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 chainId) view returns(address l2Bridge)
func (_L1 *L1Session) L2BridgeAddress(chainId *big.Int) (common.Address, error) {
	return _L1.Contract.L2BridgeAddress(&_L1.CallOpts, chainId)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 chainId) view returns(address l2Bridge)
func (_L1 *L1CallerSession) L2BridgeAddress(chainId *big.Int) (common.Address, error) {
	return _L1.Contract.L2BridgeAddress(&_L1.CallOpts, chainId)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_L1 *L1Caller) LegacyBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "legacyBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_L1 *L1Session) LegacyBridge() (common.Address, error) {
	return _L1.Contract.LegacyBridge(&_L1.CallOpts)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_L1 *L1CallerSession) LegacyBridge() (common.Address, error) {
	return _L1.Contract.LegacyBridge(&_L1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1 *L1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1 *L1Session) Owner() (common.Address, error) {
	return _L1.Contract.Owner(&_L1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1 *L1CallerSession) Owner() (common.Address, error) {
	return _L1.Contract.Owner(&_L1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1 *L1Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1 *L1Session) Paused() (bool, error) {
	return _L1.Contract.Paused(&_L1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1 *L1CallerSession) Paused() (bool, error) {
	return _L1.Contract.Paused(&_L1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_L1 *L1Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_L1 *L1Session) PendingOwner() (common.Address, error) {
	return _L1.Contract.PendingOwner(&_L1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_L1 *L1CallerSession) PendingOwner() (common.Address, error) {
	return _L1.Contract.PendingOwner(&_L1.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_L1 *L1Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_L1 *L1Session) AcceptOwnership() (*types.Transaction, error) {
	return _L1.Contract.AcceptOwnership(&_L1.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_L1 *L1TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _L1.Contract.AcceptOwnership(&_L1.TransactOpts)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_L1 *L1Transactor) BridgehubConfirmL2Transaction(opts *bind.TransactOpts, _chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "bridgehubConfirmL2Transaction", _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_L1 *L1Session) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _L1.Contract.BridgehubConfirmL2Transaction(&_L1.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_L1 *L1TransactorSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _L1.Contract.BridgehubConfirmL2Transaction(&_L1.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, uint256 _l2Value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_L1 *L1Transactor) BridgehubDeposit(opts *bind.TransactOpts, _chainId *big.Int, _prevMsgSender common.Address, _l2Value *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "bridgehubDeposit", _chainId, _prevMsgSender, _l2Value, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, uint256 _l2Value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_L1 *L1Session) BridgehubDeposit(_chainId *big.Int, _prevMsgSender common.Address, _l2Value *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1.Contract.BridgehubDeposit(&_L1.TransactOpts, _chainId, _prevMsgSender, _l2Value, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, uint256 _l2Value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_L1 *L1TransactorSession) BridgehubDeposit(_chainId *big.Int, _prevMsgSender common.Address, _l2Value *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1.Contract.BridgehubDeposit(&_L1.TransactOpts, _chainId, _prevMsgSender, _l2Value, _data)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_L1 *L1Transactor) BridgehubDepositBaseToken(opts *bind.TransactOpts, _chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "bridgehubDepositBaseToken", _chainId, _prevMsgSender, _l1Token, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_L1 *L1Session) BridgehubDepositBaseToken(_chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.BridgehubDepositBaseToken(&_L1.TransactOpts, _chainId, _prevMsgSender, _l1Token, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_L1 *L1TransactorSession) BridgehubDepositBaseToken(_chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.BridgehubDepositBaseToken(&_L1.TransactOpts, _chainId, _prevMsgSender, _l1Token, _amount)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1 *L1Transactor) ClaimFailedDeposit(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "claimFailedDeposit", _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1 *L1Session) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.ClaimFailedDeposit(&_L1.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1 *L1TransactorSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.ClaimFailedDeposit(&_L1.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1 *L1Transactor) ClaimFailedDepositLegacyErc20Bridge(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "claimFailedDepositLegacyErc20Bridge", _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1 *L1Session) ClaimFailedDepositLegacyErc20Bridge(_depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.ClaimFailedDepositLegacyErc20Bridge(&_L1.TransactOpts, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_L1 *L1TransactorSession) ClaimFailedDepositLegacyErc20Bridge(_depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.ClaimFailedDepositLegacyErc20Bridge(&_L1.TransactOpts, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _prevMsgSender, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1 *L1Transactor) DepositLegacyErc20Bridge(opts *bind.TransactOpts, _prevMsgSender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "depositLegacyErc20Bridge", _prevMsgSender, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _prevMsgSender, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1 *L1Session) DepositLegacyErc20Bridge(_prevMsgSender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1.Contract.DepositLegacyErc20Bridge(&_L1.TransactOpts, _prevMsgSender, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _prevMsgSender, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1 *L1TransactorSession) DepositLegacyErc20Bridge(_prevMsgSender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1.Contract.DepositLegacyErc20Bridge(&_L1.TransactOpts, _prevMsgSender, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1 *L1Transactor) FinalizeWithdrawal(opts *bind.TransactOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "finalizeWithdrawal", _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1 *L1Session) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.FinalizeWithdrawal(&_L1.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_L1 *L1TransactorSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.FinalizeWithdrawal(&_L1.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawalLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x7ab08472.
//
// Solidity: function finalizeWithdrawalLegacyErc20Bridge(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns(address l1Receiver, address l1Token, uint256 amount)
func (_L1 *L1Transactor) FinalizeWithdrawalLegacyErc20Bridge(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "finalizeWithdrawalLegacyErc20Bridge", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawalLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x7ab08472.
//
// Solidity: function finalizeWithdrawalLegacyErc20Bridge(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns(address l1Receiver, address l1Token, uint256 amount)
func (_L1 *L1Session) FinalizeWithdrawalLegacyErc20Bridge(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.FinalizeWithdrawalLegacyErc20Bridge(&_L1.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawalLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x7ab08472.
//
// Solidity: function finalizeWithdrawalLegacyErc20Bridge(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns(address l1Receiver, address l1Token, uint256 amount)
func (_L1 *L1TransactorSession) FinalizeWithdrawalLegacyErc20Bridge(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1.Contract.FinalizeWithdrawalLegacyErc20Bridge(&_L1.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L1 *L1Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L1 *L1Session) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _L1.Contract.Initialize(&_L1.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L1 *L1TransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _L1.Contract.Initialize(&_L1.TransactOpts, _owner)
}

// InitializeChainGovernance is a paid mutator transaction binding the contract method 0xf280efbe.
//
// Solidity: function initializeChainGovernance(uint256 _chainId, address _l2BridgeAddress) returns()
func (_L1 *L1Transactor) InitializeChainGovernance(opts *bind.TransactOpts, _chainId *big.Int, _l2BridgeAddress common.Address) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "initializeChainGovernance", _chainId, _l2BridgeAddress)
}

// InitializeChainGovernance is a paid mutator transaction binding the contract method 0xf280efbe.
//
// Solidity: function initializeChainGovernance(uint256 _chainId, address _l2BridgeAddress) returns()
func (_L1 *L1Session) InitializeChainGovernance(_chainId *big.Int, _l2BridgeAddress common.Address) (*types.Transaction, error) {
	return _L1.Contract.InitializeChainGovernance(&_L1.TransactOpts, _chainId, _l2BridgeAddress)
}

// InitializeChainGovernance is a paid mutator transaction binding the contract method 0xf280efbe.
//
// Solidity: function initializeChainGovernance(uint256 _chainId, address _l2BridgeAddress) returns()
func (_L1 *L1TransactorSession) InitializeChainGovernance(_chainId *big.Int, _l2BridgeAddress common.Address) (*types.Transaction, error) {
	return _L1.Contract.InitializeChainGovernance(&_L1.TransactOpts, _chainId, _l2BridgeAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1 *L1Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1 *L1Session) Pause() (*types.Transaction, error) {
	return _L1.Contract.Pause(&_L1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1 *L1TransactorSession) Pause() (*types.Transaction, error) {
	return _L1.Contract.Pause(&_L1.TransactOpts)
}

// ReceiveEth is a paid mutator transaction binding the contract method 0xc2aaf9c4.
//
// Solidity: function receiveEth(uint256 _chainId) payable returns()
func (_L1 *L1Transactor) ReceiveEth(opts *bind.TransactOpts, _chainId *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "receiveEth", _chainId)
}

// ReceiveEth is a paid mutator transaction binding the contract method 0xc2aaf9c4.
//
// Solidity: function receiveEth(uint256 _chainId) payable returns()
func (_L1 *L1Session) ReceiveEth(_chainId *big.Int) (*types.Transaction, error) {
	return _L1.Contract.ReceiveEth(&_L1.TransactOpts, _chainId)
}

// ReceiveEth is a paid mutator transaction binding the contract method 0xc2aaf9c4.
//
// Solidity: function receiveEth(uint256 _chainId) payable returns()
func (_L1 *L1TransactorSession) ReceiveEth(_chainId *big.Int) (*types.Transaction, error) {
	return _L1.Contract.ReceiveEth(&_L1.TransactOpts, _chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1 *L1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1 *L1Session) RenounceOwnership() (*types.Transaction, error) {
	return _L1.Contract.RenounceOwnership(&_L1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1 *L1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1.Contract.RenounceOwnership(&_L1.TransactOpts)
}

// SafeTransferFundsFromLegacy is a paid mutator transaction binding the contract method 0x7a1d8d3a.
//
// Solidity: function safeTransferFundsFromLegacy(address _token, address _target, uint256 _targetChainId, uint256 _gasPerToken) returns()
func (_L1 *L1Transactor) SafeTransferFundsFromLegacy(opts *bind.TransactOpts, _token common.Address, _target common.Address, _targetChainId *big.Int, _gasPerToken *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "safeTransferFundsFromLegacy", _token, _target, _targetChainId, _gasPerToken)
}

// SafeTransferFundsFromLegacy is a paid mutator transaction binding the contract method 0x7a1d8d3a.
//
// Solidity: function safeTransferFundsFromLegacy(address _token, address _target, uint256 _targetChainId, uint256 _gasPerToken) returns()
func (_L1 *L1Session) SafeTransferFundsFromLegacy(_token common.Address, _target common.Address, _targetChainId *big.Int, _gasPerToken *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SafeTransferFundsFromLegacy(&_L1.TransactOpts, _token, _target, _targetChainId, _gasPerToken)
}

// SafeTransferFundsFromLegacy is a paid mutator transaction binding the contract method 0x7a1d8d3a.
//
// Solidity: function safeTransferFundsFromLegacy(address _token, address _target, uint256 _targetChainId, uint256 _gasPerToken) returns()
func (_L1 *L1TransactorSession) SafeTransferFundsFromLegacy(_token common.Address, _target common.Address, _targetChainId *big.Int, _gasPerToken *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SafeTransferFundsFromLegacy(&_L1.TransactOpts, _token, _target, _targetChainId, _gasPerToken)
}

// SetEraLegacyBridgeLastDepositTime is a paid mutator transaction binding the contract method 0xdd85df2d.
//
// Solidity: function setEraLegacyBridgeLastDepositTime(uint256 _eraLegacyBridgeLastDepositBatch, uint256 _eraLegacyBridgeLastDepositTxNumber) returns()
func (_L1 *L1Transactor) SetEraLegacyBridgeLastDepositTime(opts *bind.TransactOpts, _eraLegacyBridgeLastDepositBatch *big.Int, _eraLegacyBridgeLastDepositTxNumber *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "setEraLegacyBridgeLastDepositTime", _eraLegacyBridgeLastDepositBatch, _eraLegacyBridgeLastDepositTxNumber)
}

// SetEraLegacyBridgeLastDepositTime is a paid mutator transaction binding the contract method 0xdd85df2d.
//
// Solidity: function setEraLegacyBridgeLastDepositTime(uint256 _eraLegacyBridgeLastDepositBatch, uint256 _eraLegacyBridgeLastDepositTxNumber) returns()
func (_L1 *L1Session) SetEraLegacyBridgeLastDepositTime(_eraLegacyBridgeLastDepositBatch *big.Int, _eraLegacyBridgeLastDepositTxNumber *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SetEraLegacyBridgeLastDepositTime(&_L1.TransactOpts, _eraLegacyBridgeLastDepositBatch, _eraLegacyBridgeLastDepositTxNumber)
}

// SetEraLegacyBridgeLastDepositTime is a paid mutator transaction binding the contract method 0xdd85df2d.
//
// Solidity: function setEraLegacyBridgeLastDepositTime(uint256 _eraLegacyBridgeLastDepositBatch, uint256 _eraLegacyBridgeLastDepositTxNumber) returns()
func (_L1 *L1TransactorSession) SetEraLegacyBridgeLastDepositTime(_eraLegacyBridgeLastDepositBatch *big.Int, _eraLegacyBridgeLastDepositTxNumber *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SetEraLegacyBridgeLastDepositTime(&_L1.TransactOpts, _eraLegacyBridgeLastDepositBatch, _eraLegacyBridgeLastDepositTxNumber)
}

// SetEraPostDiamondUpgradeFirstBatch is a paid mutator transaction binding the contract method 0xcc3fbc63.
//
// Solidity: function setEraPostDiamondUpgradeFirstBatch(uint256 _eraPostDiamondUpgradeFirstBatch) returns()
func (_L1 *L1Transactor) SetEraPostDiamondUpgradeFirstBatch(opts *bind.TransactOpts, _eraPostDiamondUpgradeFirstBatch *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "setEraPostDiamondUpgradeFirstBatch", _eraPostDiamondUpgradeFirstBatch)
}

// SetEraPostDiamondUpgradeFirstBatch is a paid mutator transaction binding the contract method 0xcc3fbc63.
//
// Solidity: function setEraPostDiamondUpgradeFirstBatch(uint256 _eraPostDiamondUpgradeFirstBatch) returns()
func (_L1 *L1Session) SetEraPostDiamondUpgradeFirstBatch(_eraPostDiamondUpgradeFirstBatch *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SetEraPostDiamondUpgradeFirstBatch(&_L1.TransactOpts, _eraPostDiamondUpgradeFirstBatch)
}

// SetEraPostDiamondUpgradeFirstBatch is a paid mutator transaction binding the contract method 0xcc3fbc63.
//
// Solidity: function setEraPostDiamondUpgradeFirstBatch(uint256 _eraPostDiamondUpgradeFirstBatch) returns()
func (_L1 *L1TransactorSession) SetEraPostDiamondUpgradeFirstBatch(_eraPostDiamondUpgradeFirstBatch *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SetEraPostDiamondUpgradeFirstBatch(&_L1.TransactOpts, _eraPostDiamondUpgradeFirstBatch)
}

// SetEraPostLegacyBridgeUpgradeFirstBatch is a paid mutator transaction binding the contract method 0xbe65940a.
//
// Solidity: function setEraPostLegacyBridgeUpgradeFirstBatch(uint256 _eraPostLegacyBridgeUpgradeFirstBatch) returns()
func (_L1 *L1Transactor) SetEraPostLegacyBridgeUpgradeFirstBatch(opts *bind.TransactOpts, _eraPostLegacyBridgeUpgradeFirstBatch *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "setEraPostLegacyBridgeUpgradeFirstBatch", _eraPostLegacyBridgeUpgradeFirstBatch)
}

// SetEraPostLegacyBridgeUpgradeFirstBatch is a paid mutator transaction binding the contract method 0xbe65940a.
//
// Solidity: function setEraPostLegacyBridgeUpgradeFirstBatch(uint256 _eraPostLegacyBridgeUpgradeFirstBatch) returns()
func (_L1 *L1Session) SetEraPostLegacyBridgeUpgradeFirstBatch(_eraPostLegacyBridgeUpgradeFirstBatch *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SetEraPostLegacyBridgeUpgradeFirstBatch(&_L1.TransactOpts, _eraPostLegacyBridgeUpgradeFirstBatch)
}

// SetEraPostLegacyBridgeUpgradeFirstBatch is a paid mutator transaction binding the contract method 0xbe65940a.
//
// Solidity: function setEraPostLegacyBridgeUpgradeFirstBatch(uint256 _eraPostLegacyBridgeUpgradeFirstBatch) returns()
func (_L1 *L1TransactorSession) SetEraPostLegacyBridgeUpgradeFirstBatch(_eraPostLegacyBridgeUpgradeFirstBatch *big.Int) (*types.Transaction, error) {
	return _L1.Contract.SetEraPostLegacyBridgeUpgradeFirstBatch(&_L1.TransactOpts, _eraPostLegacyBridgeUpgradeFirstBatch)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_L1 *L1Transactor) SetL1Erc20Bridge(opts *bind.TransactOpts, _legacyBridge common.Address) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "setL1Erc20Bridge", _legacyBridge)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_L1 *L1Session) SetL1Erc20Bridge(_legacyBridge common.Address) (*types.Transaction, error) {
	return _L1.Contract.SetL1Erc20Bridge(&_L1.TransactOpts, _legacyBridge)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_L1 *L1TransactorSession) SetL1Erc20Bridge(_legacyBridge common.Address) (*types.Transaction, error) {
	return _L1.Contract.SetL1Erc20Bridge(&_L1.TransactOpts, _legacyBridge)
}

// TransferFundsFromLegacy is a paid mutator transaction binding the contract method 0xc846f6df.
//
// Solidity: function transferFundsFromLegacy(address _token, address _target, uint256 _targetChainId) returns()
func (_L1 *L1Transactor) TransferFundsFromLegacy(opts *bind.TransactOpts, _token common.Address, _target common.Address, _targetChainId *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "transferFundsFromLegacy", _token, _target, _targetChainId)
}

// TransferFundsFromLegacy is a paid mutator transaction binding the contract method 0xc846f6df.
//
// Solidity: function transferFundsFromLegacy(address _token, address _target, uint256 _targetChainId) returns()
func (_L1 *L1Session) TransferFundsFromLegacy(_token common.Address, _target common.Address, _targetChainId *big.Int) (*types.Transaction, error) {
	return _L1.Contract.TransferFundsFromLegacy(&_L1.TransactOpts, _token, _target, _targetChainId)
}

// TransferFundsFromLegacy is a paid mutator transaction binding the contract method 0xc846f6df.
//
// Solidity: function transferFundsFromLegacy(address _token, address _target, uint256 _targetChainId) returns()
func (_L1 *L1TransactorSession) TransferFundsFromLegacy(_token common.Address, _target common.Address, _targetChainId *big.Int) (*types.Transaction, error) {
	return _L1.Contract.TransferFundsFromLegacy(&_L1.TransactOpts, _token, _target, _targetChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1 *L1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1 *L1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1.Contract.TransferOwnership(&_L1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1 *L1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1.Contract.TransferOwnership(&_L1.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1 *L1Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1 *L1Session) Unpause() (*types.Transaction, error) {
	return _L1.Contract.Unpause(&_L1.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1 *L1TransactorSession) Unpause() (*types.Transaction, error) {
	return _L1.Contract.Unpause(&_L1.TransactOpts)
}

// L1BridgehubDepositBaseTokenInitiatedIterator is returned from FilterBridgehubDepositBaseTokenInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositBaseTokenInitiated events raised by the L1 contract.
type L1BridgehubDepositBaseTokenInitiatedIterator struct {
	Event *L1BridgehubDepositBaseTokenInitiated // Event containing the contract specifics and raw log

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
func (it *L1BridgehubDepositBaseTokenInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1BridgehubDepositBaseTokenInitiated)
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
		it.Event = new(L1BridgehubDepositBaseTokenInitiated)
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
func (it *L1BridgehubDepositBaseTokenInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1BridgehubDepositBaseTokenInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1BridgehubDepositBaseTokenInitiated represents a BridgehubDepositBaseTokenInitiated event raised by the L1 contract.
type L1BridgehubDepositBaseTokenInitiated struct {
	ChainId *big.Int
	From    common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositBaseTokenInitiated is a free log retrieval operation binding the contract event 0x249bc8a55d0c4a0034b9aaa6be739bec2d4466e5d859bec9566a8553c405c838.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, address l1Token, uint256 amount)
func (_L1 *L1Filterer) FilterBridgehubDepositBaseTokenInitiated(opts *bind.FilterOpts, chainId []*big.Int, from []common.Address) (*L1BridgehubDepositBaseTokenInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1BridgehubDepositBaseTokenInitiatedIterator{contract: _L1.contract, event: "BridgehubDepositBaseTokenInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositBaseTokenInitiated is a free log subscription operation binding the contract event 0x249bc8a55d0c4a0034b9aaa6be739bec2d4466e5d859bec9566a8553c405c838.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, address l1Token, uint256 amount)
func (_L1 *L1Filterer) WatchBridgehubDepositBaseTokenInitiated(opts *bind.WatchOpts, sink chan<- *L1BridgehubDepositBaseTokenInitiated, chainId []*big.Int, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1BridgehubDepositBaseTokenInitiated)
				if err := _L1.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
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

// ParseBridgehubDepositBaseTokenInitiated is a log parse operation binding the contract event 0x249bc8a55d0c4a0034b9aaa6be739bec2d4466e5d859bec9566a8553c405c838.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, address l1Token, uint256 amount)
func (_L1 *L1Filterer) ParseBridgehubDepositBaseTokenInitiated(log types.Log) (*L1BridgehubDepositBaseTokenInitiated, error) {
	event := new(L1BridgehubDepositBaseTokenInitiated)
	if err := _L1.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1BridgehubDepositFinalizedIterator is returned from FilterBridgehubDepositFinalized and is used to iterate over the raw logs and unpacked data for BridgehubDepositFinalized events raised by the L1 contract.
type L1BridgehubDepositFinalizedIterator struct {
	Event *L1BridgehubDepositFinalized // Event containing the contract specifics and raw log

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
func (it *L1BridgehubDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1BridgehubDepositFinalized)
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
		it.Event = new(L1BridgehubDepositFinalized)
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
func (it *L1BridgehubDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1BridgehubDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1BridgehubDepositFinalized represents a BridgehubDepositFinalized event raised by the L1 contract.
type L1BridgehubDepositFinalized struct {
	ChainId         *big.Int
	TxDataHash      [32]byte
	L2DepositTxHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositFinalized is a free log retrieval operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_L1 *L1Filterer) FilterBridgehubDepositFinalized(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (*L1BridgehubDepositFinalizedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return &L1BridgehubDepositFinalizedIterator{contract: _L1.contract, event: "BridgehubDepositFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositFinalized is a free log subscription operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_L1 *L1Filterer) WatchBridgehubDepositFinalized(opts *bind.WatchOpts, sink chan<- *L1BridgehubDepositFinalized, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1BridgehubDepositFinalized)
				if err := _L1.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
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

// ParseBridgehubDepositFinalized is a log parse operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_L1 *L1Filterer) ParseBridgehubDepositFinalized(log types.Log) (*L1BridgehubDepositFinalized, error) {
	event := new(L1BridgehubDepositFinalized)
	if err := _L1.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1BridgehubDepositInitiatedIterator is returned from FilterBridgehubDepositInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositInitiated events raised by the L1 contract.
type L1BridgehubDepositInitiatedIterator struct {
	Event *L1BridgehubDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1BridgehubDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1BridgehubDepositInitiated)
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
		it.Event = new(L1BridgehubDepositInitiated)
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
func (it *L1BridgehubDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1BridgehubDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1BridgehubDepositInitiated represents a BridgehubDepositInitiated event raised by the L1 contract.
type L1BridgehubDepositInitiated struct {
	ChainId    *big.Int
	TxDataHash [32]byte
	From       common.Address
	To         common.Address
	L1Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositInitiated is a free log retrieval operation binding the contract event 0x8768405a01370685449c74c293804d6c9cc216d170acdbdba50b33ed4144447f.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_L1 *L1Filterer) FilterBridgehubDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (*L1BridgehubDepositInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1BridgehubDepositInitiatedIterator{contract: _L1.contract, event: "BridgehubDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositInitiated is a free log subscription operation binding the contract event 0x8768405a01370685449c74c293804d6c9cc216d170acdbdba50b33ed4144447f.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_L1 *L1Filterer) WatchBridgehubDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1BridgehubDepositInitiated, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1BridgehubDepositInitiated)
				if err := _L1.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
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

// ParseBridgehubDepositInitiated is a log parse operation binding the contract event 0x8768405a01370685449c74c293804d6c9cc216d170acdbdba50b33ed4144447f.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_L1 *L1Filterer) ParseBridgehubDepositInitiated(log types.Log) (*L1BridgehubDepositInitiated, error) {
	event := new(L1BridgehubDepositInitiated)
	if err := _L1.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ClaimedFailedDepositSharedBridgeIterator is returned from FilterClaimedFailedDepositSharedBridge and is used to iterate over the raw logs and unpacked data for ClaimedFailedDepositSharedBridge events raised by the L1 contract.
type L1ClaimedFailedDepositSharedBridgeIterator struct {
	Event *L1ClaimedFailedDepositSharedBridge // Event containing the contract specifics and raw log

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
func (it *L1ClaimedFailedDepositSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ClaimedFailedDepositSharedBridge)
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
		it.Event = new(L1ClaimedFailedDepositSharedBridge)
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
func (it *L1ClaimedFailedDepositSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ClaimedFailedDepositSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ClaimedFailedDepositSharedBridge represents a ClaimedFailedDepositSharedBridge event raised by the L1 contract.
type L1ClaimedFailedDepositSharedBridge struct {
	ChainId *big.Int
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDepositSharedBridge is a free log retrieval operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_L1 *L1Filterer) FilterClaimedFailedDepositSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, to []common.Address, l1Token []common.Address) (*L1ClaimedFailedDepositSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "ClaimedFailedDepositSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ClaimedFailedDepositSharedBridgeIterator{contract: _L1.contract, event: "ClaimedFailedDepositSharedBridge", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDepositSharedBridge is a free log subscription operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_L1 *L1Filterer) WatchClaimedFailedDepositSharedBridge(opts *bind.WatchOpts, sink chan<- *L1ClaimedFailedDepositSharedBridge, chainId []*big.Int, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "ClaimedFailedDepositSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ClaimedFailedDepositSharedBridge)
				if err := _L1.contract.UnpackLog(event, "ClaimedFailedDepositSharedBridge", log); err != nil {
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

// ParseClaimedFailedDepositSharedBridge is a log parse operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_L1 *L1Filterer) ParseClaimedFailedDepositSharedBridge(log types.Log) (*L1ClaimedFailedDepositSharedBridge, error) {
	event := new(L1ClaimedFailedDepositSharedBridge)
	if err := _L1.contract.UnpackLog(event, "ClaimedFailedDepositSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1 contract.
type L1InitializedIterator struct {
	Event *L1Initialized // Event containing the contract specifics and raw log

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
func (it *L1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1Initialized)
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
		it.Event = new(L1Initialized)
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
func (it *L1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1Initialized represents a Initialized event raised by the L1 contract.
type L1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1 *L1Filterer) FilterInitialized(opts *bind.FilterOpts) (*L1InitializedIterator, error) {

	logs, sub, err := _L1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1InitializedIterator{contract: _L1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1 *L1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1Initialized) (event.Subscription, error) {

	logs, sub, err := _L1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1Initialized)
				if err := _L1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1 *L1Filterer) ParseInitialized(log types.Log) (*L1Initialized, error) {
	event := new(L1Initialized)
	if err := _L1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1LegacyDepositInitiatedIterator is returned from FilterLegacyDepositInitiated and is used to iterate over the raw logs and unpacked data for LegacyDepositInitiated events raised by the L1 contract.
type L1LegacyDepositInitiatedIterator struct {
	Event *L1LegacyDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1LegacyDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1LegacyDepositInitiated)
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
		it.Event = new(L1LegacyDepositInitiated)
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
func (it *L1LegacyDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1LegacyDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1LegacyDepositInitiated represents a LegacyDepositInitiated event raised by the L1 contract.
type L1LegacyDepositInitiated struct {
	ChainId         *big.Int
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLegacyDepositInitiated is a free log retrieval operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_L1 *L1Filterer) FilterLegacyDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (*L1LegacyDepositInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "LegacyDepositInitiated", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1LegacyDepositInitiatedIterator{contract: _L1.contract, event: "LegacyDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchLegacyDepositInitiated is a free log subscription operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_L1 *L1Filterer) WatchLegacyDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1LegacyDepositInitiated, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "LegacyDepositInitiated", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1LegacyDepositInitiated)
				if err := _L1.contract.UnpackLog(event, "LegacyDepositInitiated", log); err != nil {
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

// ParseLegacyDepositInitiated is a log parse operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_L1 *L1Filterer) ParseLegacyDepositInitiated(log types.Log) (*L1LegacyDepositInitiated, error) {
	event := new(L1LegacyDepositInitiated)
	if err := _L1.contract.UnpackLog(event, "LegacyDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the L1 contract.
type L1OwnershipTransferStartedIterator struct {
	Event *L1OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *L1OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OwnershipTransferStarted)
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
		it.Event = new(L1OwnershipTransferStarted)
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
func (it *L1OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the L1 contract.
type L1OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1OwnershipTransferStartedIterator{contract: _L1.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *L1OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OwnershipTransferStarted)
				if err := _L1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) ParseOwnershipTransferStarted(log types.Log) (*L1OwnershipTransferStarted, error) {
	event := new(L1OwnershipTransferStarted)
	if err := _L1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1 contract.
type L1OwnershipTransferredIterator struct {
	Event *L1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OwnershipTransferred)
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
		it.Event = new(L1OwnershipTransferred)
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
func (it *L1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OwnershipTransferred represents a OwnershipTransferred event raised by the L1 contract.
type L1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1OwnershipTransferredIterator{contract: _L1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OwnershipTransferred)
				if err := _L1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) ParseOwnershipTransferred(log types.Log) (*L1OwnershipTransferred, error) {
	event := new(L1OwnershipTransferred)
	if err := _L1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1 contract.
type L1PausedIterator struct {
	Event *L1Paused // Event containing the contract specifics and raw log

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
func (it *L1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1Paused)
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
		it.Event = new(L1Paused)
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
func (it *L1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1Paused represents a Paused event raised by the L1 contract.
type L1Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1 *L1Filterer) FilterPaused(opts *bind.FilterOpts) (*L1PausedIterator, error) {

	logs, sub, err := _L1.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1PausedIterator{contract: _L1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1 *L1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1Paused) (event.Subscription, error) {

	logs, sub, err := _L1.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1Paused)
				if err := _L1.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1 *L1Filterer) ParsePaused(log types.Log) (*L1Paused, error) {
	event := new(L1Paused)
	if err := _L1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1 contract.
type L1UnpausedIterator struct {
	Event *L1Unpaused // Event containing the contract specifics and raw log

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
func (it *L1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1Unpaused)
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
		it.Event = new(L1Unpaused)
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
func (it *L1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1Unpaused represents a Unpaused event raised by the L1 contract.
type L1Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1 *L1Filterer) FilterUnpaused(opts *bind.FilterOpts) (*L1UnpausedIterator, error) {

	logs, sub, err := _L1.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1UnpausedIterator{contract: _L1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1 *L1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1Unpaused) (event.Subscription, error) {

	logs, sub, err := _L1.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1Unpaused)
				if err := _L1.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1 *L1Filterer) ParseUnpaused(log types.Log) (*L1Unpaused, error) {
	event := new(L1Unpaused)
	if err := _L1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1WithdrawalFinalizedSharedBridgeIterator is returned from FilterWithdrawalFinalizedSharedBridge and is used to iterate over the raw logs and unpacked data for WithdrawalFinalizedSharedBridge events raised by the L1 contract.
type L1WithdrawalFinalizedSharedBridgeIterator struct {
	Event *L1WithdrawalFinalizedSharedBridge // Event containing the contract specifics and raw log

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
func (it *L1WithdrawalFinalizedSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1WithdrawalFinalizedSharedBridge)
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
		it.Event = new(L1WithdrawalFinalizedSharedBridge)
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
func (it *L1WithdrawalFinalizedSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1WithdrawalFinalizedSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1WithdrawalFinalizedSharedBridge represents a WithdrawalFinalizedSharedBridge event raised by the L1 contract.
type L1WithdrawalFinalizedSharedBridge struct {
	ChainId *big.Int
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalizedSharedBridge is a free log retrieval operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_L1 *L1Filterer) FilterWithdrawalFinalizedSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, to []common.Address, l1Token []common.Address) (*L1WithdrawalFinalizedSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "WithdrawalFinalizedSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1WithdrawalFinalizedSharedBridgeIterator{contract: _L1.contract, event: "WithdrawalFinalizedSharedBridge", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalizedSharedBridge is a free log subscription operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_L1 *L1Filterer) WatchWithdrawalFinalizedSharedBridge(opts *bind.WatchOpts, sink chan<- *L1WithdrawalFinalizedSharedBridge, chainId []*big.Int, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "WithdrawalFinalizedSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1WithdrawalFinalizedSharedBridge)
				if err := _L1.contract.UnpackLog(event, "WithdrawalFinalizedSharedBridge", log); err != nil {
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

// ParseWithdrawalFinalizedSharedBridge is a log parse operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_L1 *L1Filterer) ParseWithdrawalFinalizedSharedBridge(log types.Log) (*L1WithdrawalFinalizedSharedBridge, error) {
	event := new(L1WithdrawalFinalizedSharedBridge)
	if err := _L1.contract.UnpackLog(event, "WithdrawalFinalizedSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
