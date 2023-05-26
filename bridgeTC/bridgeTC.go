// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeTC

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

// BridgeTCMetaData contains all meta data concerning the BridgeTC contract.
var BridgeTCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddr\",\"type\":\"string\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"externalAddr\",\"type\":\"string\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"safeMultisigContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeTCABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTCMetaData.ABI instead.
var BridgeTCABI = BridgeTCMetaData.ABI

// BridgeTC is an auto generated Go binding around an Ethereum contract.
type BridgeTC struct {
	BridgeTCCaller     // Read-only binding to the contract
	BridgeTCTransactor // Write-only binding to the contract
	BridgeTCFilterer   // Log filterer for contract events
}

// BridgeTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTCSession struct {
	Contract     *BridgeTC         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTCCallerSession struct {
	Contract *BridgeTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BridgeTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTCTransactorSession struct {
	Contract     *BridgeTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BridgeTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTCRaw struct {
	Contract *BridgeTC // Generic contract binding to access the raw methods on
}

// BridgeTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTCCallerRaw struct {
	Contract *BridgeTCCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTCTransactorRaw struct {
	Contract *BridgeTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTC creates a new instance of BridgeTC, bound to a specific deployed contract.
func NewBridgeTC(address common.Address, backend bind.ContractBackend) (*BridgeTC, error) {
	contract, err := bindBridgeTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTC{BridgeTCCaller: BridgeTCCaller{contract: contract}, BridgeTCTransactor: BridgeTCTransactor{contract: contract}, BridgeTCFilterer: BridgeTCFilterer{contract: contract}}, nil
}

// NewBridgeTCCaller creates a new read-only instance of BridgeTC, bound to a specific deployed contract.
func NewBridgeTCCaller(address common.Address, caller bind.ContractCaller) (*BridgeTCCaller, error) {
	contract, err := bindBridgeTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTCCaller{contract: contract}, nil
}

// NewBridgeTCTransactor creates a new write-only instance of BridgeTC, bound to a specific deployed contract.
func NewBridgeTCTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTCTransactor, error) {
	contract, err := bindBridgeTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTCTransactor{contract: contract}, nil
}

// NewBridgeTCFilterer creates a new log filterer instance of BridgeTC, bound to a specific deployed contract.
func NewBridgeTCFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTCFilterer, error) {
	contract, err := bindBridgeTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTCFilterer{contract: contract}, nil
}

// bindBridgeTC binds a generic wrapper to an already deployed contract.
func bindBridgeTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTC *BridgeTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTC.Contract.BridgeTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTC *BridgeTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTC.Contract.BridgeTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTC *BridgeTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTC.Contract.BridgeTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTC *BridgeTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTC *BridgeTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTC *BridgeTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTC.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTC *BridgeTCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTC *BridgeTCSession) Owner() (common.Address, error) {
	return _BridgeTC.Contract.Owner(&_BridgeTC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTC *BridgeTCCallerSession) Owner() (common.Address, error) {
	return _BridgeTC.Contract.Owner(&_BridgeTC.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x15f570dc.
//
// Solidity: function burn(address token, uint256 amount, string externalAddr) returns()
func (_BridgeTC *BridgeTCTransactor) Burn(opts *bind.TransactOpts, token common.Address, amount *big.Int, externalAddr string) (*types.Transaction, error) {
	return _BridgeTC.contract.Transact(opts, "burn", token, amount, externalAddr)
}

// Burn is a paid mutator transaction binding the contract method 0x15f570dc.
//
// Solidity: function burn(address token, uint256 amount, string externalAddr) returns()
func (_BridgeTC *BridgeTCSession) Burn(token common.Address, amount *big.Int, externalAddr string) (*types.Transaction, error) {
	return _BridgeTC.Contract.Burn(&_BridgeTC.TransactOpts, token, amount, externalAddr)
}

// Burn is a paid mutator transaction binding the contract method 0x15f570dc.
//
// Solidity: function burn(address token, uint256 amount, string externalAddr) returns()
func (_BridgeTC *BridgeTCTransactorSession) Burn(token common.Address, amount *big.Int, externalAddr string) (*types.Transaction, error) {
	return _BridgeTC.Contract.Burn(&_BridgeTC.TransactOpts, token, amount, externalAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_BridgeTC *BridgeTCTransactor) Initialize(opts *bind.TransactOpts, safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _BridgeTC.contract.Transact(opts, "initialize", safeMultisigContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_BridgeTC *BridgeTCSession) Initialize(safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _BridgeTC.Contract.Initialize(&_BridgeTC.TransactOpts, safeMultisigContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_BridgeTC *BridgeTCTransactorSession) Initialize(safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _BridgeTC.Contract.Initialize(&_BridgeTC.TransactOpts, safeMultisigContractAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_BridgeTC *BridgeTCTransactor) Mint(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeTC.contract.Transact(opts, "mint", tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_BridgeTC *BridgeTCSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeTC.Contract.Mint(&_BridgeTC.TransactOpts, tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_BridgeTC *BridgeTCTransactorSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeTC.Contract.Mint(&_BridgeTC.TransactOpts, tokens, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_BridgeTC *BridgeTCTransactor) Mint0(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeTC.contract.Transact(opts, "mint0", token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_BridgeTC *BridgeTCSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeTC.Contract.Mint0(&_BridgeTC.TransactOpts, token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_BridgeTC *BridgeTCTransactorSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeTC.Contract.Mint0(&_BridgeTC.TransactOpts, token, recipients, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeTC *BridgeTCTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTC.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeTC *BridgeTCSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeTC.Contract.RenounceOwnership(&_BridgeTC.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeTC *BridgeTCTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeTC.Contract.RenounceOwnership(&_BridgeTC.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTC *BridgeTCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTC *BridgeTCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTC.Contract.TransferOwnership(&_BridgeTC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTC *BridgeTCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTC.Contract.TransferOwnership(&_BridgeTC.TransactOpts, newOwner)
}

// BridgeTCBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BridgeTC contract.
type BridgeTCBurnIterator struct {
	Event *BridgeTCBurn // Event containing the contract specifics and raw log

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
func (it *BridgeTCBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTCBurn)
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
		it.Event = new(BridgeTCBurn)
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
func (it *BridgeTCBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTCBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTCBurn represents a Burn event raised by the BridgeTC contract.
type BridgeTCBurn struct {
	Token   common.Address
	Burner  common.Address
	Amount  *big.Int
	BtcAddr string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xeab88a0b2198f2928ade5fb787115a9a6ffbbf3705277143953a7c26769157ff.
//
// Solidity: event Burn(address token, address burner, uint256 amount, string btcAddr)
func (_BridgeTC *BridgeTCFilterer) FilterBurn(opts *bind.FilterOpts) (*BridgeTCBurnIterator, error) {

	logs, sub, err := _BridgeTC.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &BridgeTCBurnIterator{contract: _BridgeTC.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xeab88a0b2198f2928ade5fb787115a9a6ffbbf3705277143953a7c26769157ff.
//
// Solidity: event Burn(address token, address burner, uint256 amount, string btcAddr)
func (_BridgeTC *BridgeTCFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BridgeTCBurn) (event.Subscription, error) {

	logs, sub, err := _BridgeTC.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTCBurn)
				if err := _BridgeTC.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xeab88a0b2198f2928ade5fb787115a9a6ffbbf3705277143953a7c26769157ff.
//
// Solidity: event Burn(address token, address burner, uint256 amount, string btcAddr)
func (_BridgeTC *BridgeTCFilterer) ParseBurn(log types.Log) (*BridgeTCBurn, error) {
	event := new(BridgeTCBurn)
	if err := _BridgeTC.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTCInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeTC contract.
type BridgeTCInitializedIterator struct {
	Event *BridgeTCInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeTCInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTCInitialized)
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
		it.Event = new(BridgeTCInitialized)
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
func (it *BridgeTCInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTCInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTCInitialized represents a Initialized event raised by the BridgeTC contract.
type BridgeTCInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTC *BridgeTCFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeTCInitializedIterator, error) {

	logs, sub, err := _BridgeTC.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeTCInitializedIterator{contract: _BridgeTC.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTC *BridgeTCFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeTCInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeTC.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTCInitialized)
				if err := _BridgeTC.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeTC *BridgeTCFilterer) ParseInitialized(log types.Log) (*BridgeTCInitialized, error) {
	event := new(BridgeTCInitialized)
	if err := _BridgeTC.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTCMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BridgeTC contract.
type BridgeTCMintIterator struct {
	Event *BridgeTCMint // Event containing the contract specifics and raw log

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
func (it *BridgeTCMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTCMint)
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
		it.Event = new(BridgeTCMint)
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
func (it *BridgeTCMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTCMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTCMint represents a Mint event raised by the BridgeTC contract.
type BridgeTCMint struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_BridgeTC *BridgeTCFilterer) FilterMint(opts *bind.FilterOpts) (*BridgeTCMintIterator, error) {

	logs, sub, err := _BridgeTC.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &BridgeTCMintIterator{contract: _BridgeTC.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_BridgeTC *BridgeTCFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BridgeTCMint) (event.Subscription, error) {

	logs, sub, err := _BridgeTC.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTCMint)
				if err := _BridgeTC.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_BridgeTC *BridgeTCFilterer) ParseMint(log types.Log) (*BridgeTCMint, error) {
	event := new(BridgeTCMint)
	if err := _BridgeTC.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTCMint0Iterator is returned from FilterMint0 and is used to iterate over the raw logs and unpacked data for Mint0 events raised by the BridgeTC contract.
type BridgeTCMint0Iterator struct {
	Event *BridgeTCMint0 // Event containing the contract specifics and raw log

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
func (it *BridgeTCMint0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTCMint0)
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
		it.Event = new(BridgeTCMint0)
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
func (it *BridgeTCMint0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTCMint0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTCMint0 represents a Mint0 event raised by the BridgeTC contract.
type BridgeTCMint0 struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint0 is a free log retrieval operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_BridgeTC *BridgeTCFilterer) FilterMint0(opts *bind.FilterOpts) (*BridgeTCMint0Iterator, error) {

	logs, sub, err := _BridgeTC.contract.FilterLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return &BridgeTCMint0Iterator{contract: _BridgeTC.contract, event: "Mint0", logs: logs, sub: sub}, nil
}

// WatchMint0 is a free log subscription operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_BridgeTC *BridgeTCFilterer) WatchMint0(opts *bind.WatchOpts, sink chan<- *BridgeTCMint0) (event.Subscription, error) {

	logs, sub, err := _BridgeTC.contract.WatchLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTCMint0)
				if err := _BridgeTC.contract.UnpackLog(event, "Mint0", log); err != nil {
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

// ParseMint0 is a log parse operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_BridgeTC *BridgeTCFilterer) ParseMint0(log types.Log) (*BridgeTCMint0, error) {
	event := new(BridgeTCMint0)
	if err := _BridgeTC.contract.UnpackLog(event, "Mint0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeTC contract.
type BridgeTCOwnershipTransferredIterator struct {
	Event *BridgeTCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeTCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTCOwnershipTransferred)
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
		it.Event = new(BridgeTCOwnershipTransferred)
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
func (it *BridgeTCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTCOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeTC contract.
type BridgeTCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTC *BridgeTCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeTCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeTC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTCOwnershipTransferredIterator{contract: _BridgeTC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTC *BridgeTCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeTCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeTC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTCOwnershipTransferred)
				if err := _BridgeTC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeTC *BridgeTCFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeTCOwnershipTransferred, error) {
	event := new(BridgeTCOwnershipTransferred)
	if err := _BridgeTC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
