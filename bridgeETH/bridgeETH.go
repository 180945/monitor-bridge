// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeETH

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

// BridgeETHMetaData contains all meta data concerning the BridgeETH contract.
var BridgeETHMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"externalAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"externalAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"safeMultisigContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeETHABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeETHMetaData.ABI instead.
var BridgeETHABI = BridgeETHMetaData.ABI

// BridgeETH is an auto generated Go binding around an Ethereum contract.
type BridgeETH struct {
	BridgeETHCaller     // Read-only binding to the contract
	BridgeETHTransactor // Write-only binding to the contract
	BridgeETHFilterer   // Log filterer for contract events
}

// BridgeETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeETHSession struct {
	Contract     *BridgeETH        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeETHCallerSession struct {
	Contract *BridgeETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgeETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeETHTransactorSession struct {
	Contract     *BridgeETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeETHRaw struct {
	Contract *BridgeETH // Generic contract binding to access the raw methods on
}

// BridgeETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeETHCallerRaw struct {
	Contract *BridgeETHCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeETHTransactorRaw struct {
	Contract *BridgeETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeETH creates a new instance of BridgeETH, bound to a specific deployed contract.
func NewBridgeETH(address common.Address, backend bind.ContractBackend) (*BridgeETH, error) {
	contract, err := bindBridgeETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeETH{BridgeETHCaller: BridgeETHCaller{contract: contract}, BridgeETHTransactor: BridgeETHTransactor{contract: contract}, BridgeETHFilterer: BridgeETHFilterer{contract: contract}}, nil
}

// NewBridgeETHCaller creates a new read-only instance of BridgeETH, bound to a specific deployed contract.
func NewBridgeETHCaller(address common.Address, caller bind.ContractCaller) (*BridgeETHCaller, error) {
	contract, err := bindBridgeETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeETHCaller{contract: contract}, nil
}

// NewBridgeETHTransactor creates a new write-only instance of BridgeETH, bound to a specific deployed contract.
func NewBridgeETHTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeETHTransactor, error) {
	contract, err := bindBridgeETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeETHTransactor{contract: contract}, nil
}

// NewBridgeETHFilterer creates a new log filterer instance of BridgeETH, bound to a specific deployed contract.
func NewBridgeETHFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeETHFilterer, error) {
	contract, err := bindBridgeETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeETHFilterer{contract: contract}, nil
}

// bindBridgeETH binds a generic wrapper to an already deployed contract.
func bindBridgeETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeETH *BridgeETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeETH.Contract.BridgeETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeETH *BridgeETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeETH.Contract.BridgeETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeETH *BridgeETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeETH.Contract.BridgeETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeETH *BridgeETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeETH *BridgeETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeETH *BridgeETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeETH.Contract.contract.Transact(opts, method, params...)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_BridgeETH *BridgeETHCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeETH.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_BridgeETH *BridgeETHSession) ETHTOKEN() (common.Address, error) {
	return _BridgeETH.Contract.ETHTOKEN(&_BridgeETH.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_BridgeETH *BridgeETHCallerSession) ETHTOKEN() (common.Address, error) {
	return _BridgeETH.Contract.ETHTOKEN(&_BridgeETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeETH *BridgeETHCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeETH.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeETH *BridgeETHSession) Owner() (common.Address, error) {
	return _BridgeETH.Contract.Owner(&_BridgeETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeETH *BridgeETHCallerSession) Owner() (common.Address, error) {
	return _BridgeETH.Contract.Owner(&_BridgeETH.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address externalAddr) payable returns()
func (_BridgeETH *BridgeETHTransactor) Deposit(opts *bind.TransactOpts, externalAddr common.Address) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "deposit", externalAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address externalAddr) payable returns()
func (_BridgeETH *BridgeETHSession) Deposit(externalAddr common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.Deposit(&_BridgeETH.TransactOpts, externalAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address externalAddr) payable returns()
func (_BridgeETH *BridgeETHTransactorSession) Deposit(externalAddr common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.Deposit(&_BridgeETH.TransactOpts, externalAddr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address token, uint256 amount, address externalAddr) returns()
func (_BridgeETH *BridgeETHTransactor) Deposit0(opts *bind.TransactOpts, token common.Address, amount *big.Int, externalAddr common.Address) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "deposit0", token, amount, externalAddr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address token, uint256 amount, address externalAddr) returns()
func (_BridgeETH *BridgeETHSession) Deposit0(token common.Address, amount *big.Int, externalAddr common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.Deposit0(&_BridgeETH.TransactOpts, token, amount, externalAddr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address token, uint256 amount, address externalAddr) returns()
func (_BridgeETH *BridgeETHTransactorSession) Deposit0(token common.Address, amount *big.Int, externalAddr common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.Deposit0(&_BridgeETH.TransactOpts, token, amount, externalAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_BridgeETH *BridgeETHTransactor) Initialize(opts *bind.TransactOpts, safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "initialize", safeMultisigContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_BridgeETH *BridgeETHSession) Initialize(safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.Initialize(&_BridgeETH.TransactOpts, safeMultisigContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_BridgeETH *BridgeETHTransactorSession) Initialize(safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.Initialize(&_BridgeETH.TransactOpts, safeMultisigContractAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeETH *BridgeETHTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeETH *BridgeETHSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeETH.Contract.RenounceOwnership(&_BridgeETH.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeETH *BridgeETHTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeETH.Contract.RenounceOwnership(&_BridgeETH.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeETH *BridgeETHTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeETH *BridgeETHSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.TransferOwnership(&_BridgeETH.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeETH *BridgeETHTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeETH.Contract.TransferOwnership(&_BridgeETH.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xedbd7668.
//
// Solidity: function withdraw(address token, address[] recipients, uint256[] amounts) returns()
func (_BridgeETH *BridgeETHTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "withdraw", token, recipients, amounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xedbd7668.
//
// Solidity: function withdraw(address token, address[] recipients, uint256[] amounts) returns()
func (_BridgeETH *BridgeETHSession) Withdraw(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeETH.Contract.Withdraw(&_BridgeETH.TransactOpts, token, recipients, amounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xedbd7668.
//
// Solidity: function withdraw(address token, address[] recipients, uint256[] amounts) returns()
func (_BridgeETH *BridgeETHTransactorSession) Withdraw(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeETH.Contract.Withdraw(&_BridgeETH.TransactOpts, token, recipients, amounts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf7ece0cf.
//
// Solidity: function withdraw(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_BridgeETH *BridgeETHTransactor) Withdraw0(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeETH.contract.Transact(opts, "withdraw0", tokens, recipients, amounts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf7ece0cf.
//
// Solidity: function withdraw(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_BridgeETH *BridgeETHSession) Withdraw0(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeETH.Contract.Withdraw0(&_BridgeETH.TransactOpts, tokens, recipients, amounts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf7ece0cf.
//
// Solidity: function withdraw(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_BridgeETH *BridgeETHTransactorSession) Withdraw0(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeETH.Contract.Withdraw0(&_BridgeETH.TransactOpts, tokens, recipients, amounts)
}

// BridgeETHDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BridgeETH contract.
type BridgeETHDepositIterator struct {
	Event *BridgeETHDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeETHDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeETHDeposit)
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
		it.Event = new(BridgeETHDeposit)
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
func (it *BridgeETHDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeETHDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeETHDeposit represents a Deposit event raised by the BridgeETH contract.
type BridgeETHDeposit struct {
	Token     common.Address
	Sender    common.Address
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x364bb76a44233df8584c690de6da7810626a5e77192f3ebc942c35bcb1add24f.
//
// Solidity: event Deposit(address token, address sender, uint256 amount, address recipient)
func (_BridgeETH *BridgeETHFilterer) FilterDeposit(opts *bind.FilterOpts) (*BridgeETHDepositIterator, error) {

	logs, sub, err := _BridgeETH.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BridgeETHDepositIterator{contract: _BridgeETH.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x364bb76a44233df8584c690de6da7810626a5e77192f3ebc942c35bcb1add24f.
//
// Solidity: event Deposit(address token, address sender, uint256 amount, address recipient)
func (_BridgeETH *BridgeETHFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeETHDeposit) (event.Subscription, error) {

	logs, sub, err := _BridgeETH.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeETHDeposit)
				if err := _BridgeETH.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x364bb76a44233df8584c690de6da7810626a5e77192f3ebc942c35bcb1add24f.
//
// Solidity: event Deposit(address token, address sender, uint256 amount, address recipient)
func (_BridgeETH *BridgeETHFilterer) ParseDeposit(log types.Log) (*BridgeETHDeposit, error) {
	event := new(BridgeETHDeposit)
	if err := _BridgeETH.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeETHInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeETH contract.
type BridgeETHInitializedIterator struct {
	Event *BridgeETHInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeETHInitialized)
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
		it.Event = new(BridgeETHInitialized)
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
func (it *BridgeETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeETHInitialized represents a Initialized event raised by the BridgeETH contract.
type BridgeETHInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeETH *BridgeETHFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeETHInitializedIterator, error) {

	logs, sub, err := _BridgeETH.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeETHInitializedIterator{contract: _BridgeETH.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeETH *BridgeETHFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeETHInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeETH.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeETHInitialized)
				if err := _BridgeETH.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeETH *BridgeETHFilterer) ParseInitialized(log types.Log) (*BridgeETHInitialized, error) {
	event := new(BridgeETHInitialized)
	if err := _BridgeETH.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeETHOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeETH contract.
type BridgeETHOwnershipTransferredIterator struct {
	Event *BridgeETHOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeETHOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeETHOwnershipTransferred)
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
		it.Event = new(BridgeETHOwnershipTransferred)
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
func (it *BridgeETHOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeETHOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeETHOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeETH contract.
type BridgeETHOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeETH *BridgeETHFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeETHOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeETH.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeETHOwnershipTransferredIterator{contract: _BridgeETH.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeETH *BridgeETHFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeETHOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeETH.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeETHOwnershipTransferred)
				if err := _BridgeETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeETH *BridgeETHFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeETHOwnershipTransferred, error) {
	event := new(BridgeETHOwnershipTransferred)
	if err := _BridgeETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeETHWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BridgeETH contract.
type BridgeETHWithdrawIterator struct {
	Event *BridgeETHWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeETHWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeETHWithdraw)
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
		it.Event = new(BridgeETHWithdraw)
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
func (it *BridgeETHWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeETHWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeETHWithdraw represents a Withdraw event raised by the BridgeETH contract.
type BridgeETHWithdraw struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xdf8b54622ff1f8c4514f6bd21ae42dd7c0f3111b74c909c7fd9ea7a8b9087b40.
//
// Solidity: event Withdraw(address[] tokens, address[] recipients, uint256[] amounts)
func (_BridgeETH *BridgeETHFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BridgeETHWithdrawIterator, error) {

	logs, sub, err := _BridgeETH.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BridgeETHWithdrawIterator{contract: _BridgeETH.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xdf8b54622ff1f8c4514f6bd21ae42dd7c0f3111b74c909c7fd9ea7a8b9087b40.
//
// Solidity: event Withdraw(address[] tokens, address[] recipients, uint256[] amounts)
func (_BridgeETH *BridgeETHFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeETHWithdraw) (event.Subscription, error) {

	logs, sub, err := _BridgeETH.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeETHWithdraw)
				if err := _BridgeETH.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xdf8b54622ff1f8c4514f6bd21ae42dd7c0f3111b74c909c7fd9ea7a8b9087b40.
//
// Solidity: event Withdraw(address[] tokens, address[] recipients, uint256[] amounts)
func (_BridgeETH *BridgeETHFilterer) ParseWithdraw(log types.Log) (*BridgeETHWithdraw, error) {
	event := new(BridgeETHWithdraw)
	if err := _BridgeETH.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeETHWithdraw0Iterator is returned from FilterWithdraw0 and is used to iterate over the raw logs and unpacked data for Withdraw0 events raised by the BridgeETH contract.
type BridgeETHWithdraw0Iterator struct {
	Event *BridgeETHWithdraw0 // Event containing the contract specifics and raw log

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
func (it *BridgeETHWithdraw0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeETHWithdraw0)
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
		it.Event = new(BridgeETHWithdraw0)
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
func (it *BridgeETHWithdraw0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeETHWithdraw0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeETHWithdraw0 represents a Withdraw0 event raised by the BridgeETH contract.
type BridgeETHWithdraw0 struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw0 is a free log retrieval operation binding the contract event 0xa7bd5a2454bc906a642fda333278545850c8711673010966530dcabede413367.
//
// Solidity: event Withdraw(address token, address[] recipients, uint256[] amounts)
func (_BridgeETH *BridgeETHFilterer) FilterWithdraw0(opts *bind.FilterOpts) (*BridgeETHWithdraw0Iterator, error) {

	logs, sub, err := _BridgeETH.contract.FilterLogs(opts, "Withdraw0")
	if err != nil {
		return nil, err
	}
	return &BridgeETHWithdraw0Iterator{contract: _BridgeETH.contract, event: "Withdraw0", logs: logs, sub: sub}, nil
}

// WatchWithdraw0 is a free log subscription operation binding the contract event 0xa7bd5a2454bc906a642fda333278545850c8711673010966530dcabede413367.
//
// Solidity: event Withdraw(address token, address[] recipients, uint256[] amounts)
func (_BridgeETH *BridgeETHFilterer) WatchWithdraw0(opts *bind.WatchOpts, sink chan<- *BridgeETHWithdraw0) (event.Subscription, error) {

	logs, sub, err := _BridgeETH.contract.WatchLogs(opts, "Withdraw0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeETHWithdraw0)
				if err := _BridgeETH.contract.UnpackLog(event, "Withdraw0", log); err != nil {
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

// ParseWithdraw0 is a log parse operation binding the contract event 0xa7bd5a2454bc906a642fda333278545850c8711673010966530dcabede413367.
//
// Solidity: event Withdraw(address token, address[] recipients, uint256[] amounts)
func (_BridgeETH *BridgeETHFilterer) ParseWithdraw0(log types.Log) (*BridgeETHWithdraw0, error) {
	event := new(BridgeETHWithdraw0)
	if err := _BridgeETH.contract.UnpackLog(event, "Withdraw0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
