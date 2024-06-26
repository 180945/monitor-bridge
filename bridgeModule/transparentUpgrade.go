// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// TransparentUpgradeMetaData contains all meta data concerning the TransparentUpgrade contract.
var TransparentUpgradeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405260405162000ebf38038062000ebf833981016040819052620000269162000497565b828162000036828260006200004d565b50620000449050826200008a565b505050620005ca565b6200005883620000e5565b600082511180620000665750805b1562000085576200008383836200012760201b620001691760201c565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620000b562000156565b604080516001600160a01b03928316815291841660208301520160405180910390a1620000e2816200018f565b50565b620000f08162000244565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606200014f838360405180606001604052806027815260200162000e9860279139620002f8565b9392505050565b60006200018060008051602062000e7883398151915260001b6200037760201b620001951760201c565b546001600160a01b0316919050565b6001600160a01b038116620001fa5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b806200022360008051602062000e7883398151915260001b6200037760201b620001951760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6200025a816200037a60201b620001981760201c565b620002be5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401620001f1565b80620002237f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6200037760201b620001951760201c565b6060600080856001600160a01b03168560405162000317919062000577565b600060405180830381855af49150503d806000811462000354576040519150601f19603f3d011682016040523d82523d6000602084013e62000359565b606091505b5090925090506200036d8683838762000389565b9695505050505050565b90565b6001600160a01b03163b151590565b60608315620003fd578251600003620003f5576001600160a01b0385163b620003f55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401620001f1565b508162000409565b62000409838362000411565b949350505050565b815115620004225781518083602001fd5b8060405162461bcd60e51b8152600401620001f1919062000595565b80516001600160a01b03811681146200045657600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156200048e57818101518382015260200162000474565b50506000910152565b600080600060608486031215620004ad57600080fd5b620004b8846200043e565b9250620004c8602085016200043e565b60408501519092506001600160401b0380821115620004e657600080fd5b818601915086601f830112620004fb57600080fd5b8151818111156200051057620005106200045b565b604051601f8201601f19908116603f011681019083821181831017156200053b576200053b6200045b565b816040528281528960208487010111156200055557600080fd5b6200056883602083016020880162000471565b80955050505050509250925092565b600082516200058b81846020870162000471565b9190910192915050565b6020815260008251806020840152620005b681604085016020870162000471565b601f01601f19169190910160400192915050565b61089e80620005da6000396000f3fe60806040523661001357610011610017565b005b6100115b61001f6101a7565b6001600160a01b0316330361015f5760606001600160e01b0319600035166364d3180d60e11b810161005a576100536101da565b9150610157565b63587086bd60e11b6001600160e01b031982160161007a57610053610231565b63070d7c6960e41b6001600160e01b031982160161009a57610053610277565b621eb96f60e61b6001600160e01b03198216016100b9576100536102a8565b63a39f25e560e01b6001600160e01b03198216016100d9576100536102e8565b60405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4015b60405180910390fd5b815160208301f35b6101676102fc565b565b606061018e83836040518060600160405280602781526020016108426027913961030c565b9392505050565b90565b6001600160a01b03163b151590565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b546001600160a01b0316919050565b60606101e4610384565b60006101f33660048184610695565b81019061020091906106db565b905061021d8160405180602001604052806000815250600061038f565b505060408051602081019091526000815290565b60606000806102433660048184610695565b810190610250919061070c565b915091506102608282600161038f565b604051806020016040528060008152509250505090565b6060610281610384565b60006102903660048184610695565b81019061029d91906106db565b905061021d816103bb565b60606102b2610384565b60006102bc6101a7565b604080516001600160a01b03831660208201529192500160405160208183030381529060405291505090565b60606102f2610384565b60006102bc610412565b610167610307610412565b610421565b6060600080856001600160a01b03168560405161032991906107f2565b600060405180830381855af49150503d8060008114610364576040519150601f19603f3d011682016040523d82523d6000602084013e610369565b606091505b509150915061037a86838387610445565b9695505050505050565b341561016757600080fd5b610398836104c6565b6000825111806103a55750805b156103b6576103b48383610169565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6103e46101a7565b604080516001600160a01b03928316815291841660208301520160405180910390a161040f81610506565b50565b600061041c6105af565b905090565b3660008037600080366000845af43d6000803e808015610440573d6000f35b3d6000fd5b606083156104b45782516000036104ad576001600160a01b0385163b6104ad5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161014e565b50816104be565b6104be83836105d7565b949350505050565b6104cf81610601565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6001600160a01b03811661056b5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b606482015260840161014e565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80546001600160a01b0319166001600160a01b039290921691909117905550565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6101cb565b8151156105e75781518083602001fd5b8060405162461bcd60e51b815260040161014e919061080e565b6001600160a01b0381163b61066e5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161014e565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61058e565b600080858511156106a557600080fd5b838611156106b257600080fd5b5050820193919092039150565b80356001600160a01b03811681146106d657600080fd5b919050565b6000602082840312156106ed57600080fd5b61018e826106bf565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561071f57600080fd5b610728836106bf565b9150602083013567ffffffffffffffff8082111561074557600080fd5b818501915085601f83011261075957600080fd5b81358181111561076b5761076b6106f6565b604051601f8201601f19908116603f01168101908382118183101715610793576107936106f6565b816040528281528860208487010111156107ac57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60005b838110156107e95781810151838201526020016107d1565b50506000910152565b600082516108048184602087016107ce565b9190910192915050565b602081526000825180602084015261082d8160408501602087016107ce565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220458f1d0a30b83c8d3cd4a443a44d6bfaf46822ea12726359400f54a9a262edad64736f6c63430008110033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// TransparentUpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TransparentUpgradeMetaData.ABI instead.
var TransparentUpgradeABI = TransparentUpgradeMetaData.ABI

// TransparentUpgradeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransparentUpgradeMetaData.Bin instead.
var TransparentUpgradeBin = TransparentUpgradeMetaData.Bin

// DeployTransparentUpgrade deploys a new Ethereum contract, binding an instance of TransparentUpgrade to it.
func DeployTransparentUpgrade(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, admin_ common.Address, _data []byte) (common.Address, *types.Transaction, *TransparentUpgrade, error) {
	parsed, err := TransparentUpgradeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransparentUpgradeBin), backend, _logic, admin_, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransparentUpgrade{TransparentUpgradeCaller: TransparentUpgradeCaller{contract: contract}, TransparentUpgradeTransactor: TransparentUpgradeTransactor{contract: contract}, TransparentUpgradeFilterer: TransparentUpgradeFilterer{contract: contract}}, nil
}

// TransparentUpgrade is an auto generated Go binding around an Ethereum contract.
type TransparentUpgrade struct {
	TransparentUpgradeCaller     // Read-only binding to the contract
	TransparentUpgradeTransactor // Write-only binding to the contract
	TransparentUpgradeFilterer   // Log filterer for contract events
}

// TransparentUpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransparentUpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransparentUpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransparentUpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransparentUpgradeSession struct {
	Contract     *TransparentUpgrade // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransparentUpgradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransparentUpgradeCallerSession struct {
	Contract *TransparentUpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TransparentUpgradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransparentUpgradeTransactorSession struct {
	Contract     *TransparentUpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TransparentUpgradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransparentUpgradeRaw struct {
	Contract *TransparentUpgrade // Generic contract binding to access the raw methods on
}

// TransparentUpgradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransparentUpgradeCallerRaw struct {
	Contract *TransparentUpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// TransparentUpgradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransparentUpgradeTransactorRaw struct {
	Contract *TransparentUpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransparentUpgrade creates a new instance of TransparentUpgrade, bound to a specific deployed contract.
func NewTransparentUpgrade(address common.Address, backend bind.ContractBackend) (*TransparentUpgrade, error) {
	contract, err := bindTransparentUpgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgrade{TransparentUpgradeCaller: TransparentUpgradeCaller{contract: contract}, TransparentUpgradeTransactor: TransparentUpgradeTransactor{contract: contract}, TransparentUpgradeFilterer: TransparentUpgradeFilterer{contract: contract}}, nil
}

// NewTransparentUpgradeCaller creates a new read-only instance of TransparentUpgrade, bound to a specific deployed contract.
func NewTransparentUpgradeCaller(address common.Address, caller bind.ContractCaller) (*TransparentUpgradeCaller, error) {
	contract, err := bindTransparentUpgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeCaller{contract: contract}, nil
}

// NewTransparentUpgradeTransactor creates a new write-only instance of TransparentUpgrade, bound to a specific deployed contract.
func NewTransparentUpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TransparentUpgradeTransactor, error) {
	contract, err := bindTransparentUpgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeTransactor{contract: contract}, nil
}

// NewTransparentUpgradeFilterer creates a new log filterer instance of TransparentUpgrade, bound to a specific deployed contract.
func NewTransparentUpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TransparentUpgradeFilterer, error) {
	contract, err := bindTransparentUpgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeFilterer{contract: contract}, nil
}

// bindTransparentUpgrade binds a generic wrapper to an already deployed contract.
func bindTransparentUpgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransparentUpgradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgrade *TransparentUpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgrade.Contract.TransparentUpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgrade *TransparentUpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.TransparentUpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgrade *TransparentUpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.TransparentUpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgrade *TransparentUpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgrade *TransparentUpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgrade *TransparentUpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgrade *TransparentUpgradeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgrade.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgrade *TransparentUpgradeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.Fallback(&_TransparentUpgrade.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgrade *TransparentUpgradeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.Fallback(&_TransparentUpgrade.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgrade *TransparentUpgradeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgrade.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgrade *TransparentUpgradeSession) Receive() (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.Receive(&_TransparentUpgrade.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgrade *TransparentUpgradeTransactorSession) Receive() (*types.Transaction, error) {
	return _TransparentUpgrade.Contract.Receive(&_TransparentUpgrade.TransactOpts)
}

// TransparentUpgradeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TransparentUpgrade contract.
type TransparentUpgradeAdminChangedIterator struct {
	Event *TransparentUpgradeAdminChanged // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeAdminChanged)
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
		it.Event = new(TransparentUpgradeAdminChanged)
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
func (it *TransparentUpgradeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeAdminChanged represents a AdminChanged event raised by the TransparentUpgrade contract.
type TransparentUpgradeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgrade *TransparentUpgradeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TransparentUpgradeAdminChangedIterator, error) {

	logs, sub, err := _TransparentUpgrade.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeAdminChangedIterator{contract: _TransparentUpgrade.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgrade *TransparentUpgradeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgrade.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeAdminChanged)
				if err := _TransparentUpgrade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgrade *TransparentUpgradeFilterer) ParseAdminChanged(log types.Log) (*TransparentUpgradeAdminChanged, error) {
	event := new(TransparentUpgradeAdminChanged)
	if err := _TransparentUpgrade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TransparentUpgrade contract.
type TransparentUpgradeBeaconUpgradedIterator struct {
	Event *TransparentUpgradeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeBeaconUpgraded)
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
		it.Event = new(TransparentUpgradeBeaconUpgraded)
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
func (it *TransparentUpgradeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeBeaconUpgraded represents a BeaconUpgraded event raised by the TransparentUpgrade contract.
type TransparentUpgradeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentUpgrade *TransparentUpgradeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TransparentUpgradeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TransparentUpgrade.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeBeaconUpgradedIterator{contract: _TransparentUpgrade.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentUpgrade *TransparentUpgradeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TransparentUpgrade.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeBeaconUpgraded)
				if err := _TransparentUpgrade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentUpgrade *TransparentUpgradeFilterer) ParseBeaconUpgraded(log types.Log) (*TransparentUpgradeBeaconUpgraded, error) {
	event := new(TransparentUpgradeBeaconUpgraded)
	if err := _TransparentUpgrade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TransparentUpgrade contract.
type TransparentUpgradeUpgradedIterator struct {
	Event *TransparentUpgradeUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeUpgraded)
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
		it.Event = new(TransparentUpgradeUpgraded)
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
func (it *TransparentUpgradeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeUpgraded represents a Upgraded event raised by the TransparentUpgrade contract.
type TransparentUpgradeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgrade *TransparentUpgradeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentUpgradeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgrade.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeUpgradedIterator{contract: _TransparentUpgrade.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgrade *TransparentUpgradeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgrade.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeUpgraded)
				if err := _TransparentUpgrade.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgrade *TransparentUpgradeFilterer) ParseUpgraded(log types.Log) (*TransparentUpgradeUpgraded, error) {
	event := new(TransparentUpgradeUpgraded)
	if err := _TransparentUpgrade.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
