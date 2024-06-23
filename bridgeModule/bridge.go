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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ETH_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"externalAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeToken\",\"inputs\":[{\"name\":\"externalAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"burnableToken\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"safeMultisigContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractWrappedToken[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractWrappedToken\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOperator\",\"inputs\":[{\"name\":\"operator_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateToken\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"isBurns\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractWrappedToken\"},{\"name\":\"burner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extddr\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractWrappedToken[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractWrappedToken\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506118af806100206000396000f3fe6080604052600436106100c65760003560e01c806358bc83371161007f5780638da5cb5b116100595780638da5cb5b1461021b578063a3bf277e14610239578063d4546d2314610259578063f2fde38b1461026c57600080fd5b806358bc8337146101d1578063715018a6146101e657806377a24f36146101fb57600080fd5b806307ef210c146100d25780630e93b35c1461011757806329605e77146101395780634432e6b7146101595780635530f4a514610179578063570ca7351461019957600080fd5b366100cd57005b600080fd5b3480156100de57600080fd5b506101026100ed36600461128c565b60656020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561012357600080fd5b506101376101323660046112f9565b61028c565b005b34801561014557600080fd5b5061013761015436600461128c565b610334565b34801561016557600080fd5b506101376101743660046113a2565b6103b0565b34801561018557600080fd5b5061013761019436600461140e565b6104f7565b3480156101a557600080fd5b506066546101b9906001600160a01b031681565b6040516001600160a01b03909116815260200161010e565b3480156101dd57600080fd5b506101b9600081565b3480156101f257600080fd5b50610137610773565b34801561020757600080fd5b506101376102163660046114a8565b610787565b34801561022757600080fd5b506033546001600160a01b03166101b9565b34801561024557600080fd5b50610137610254366004611501565b610996565b610137610267366004611584565b610b63565b34801561027857600080fd5b5061013761028736600461128c565b610bfb565b468181036102e15760405162461bcd60e51b815260206004820152601d60248201527f4272696467653a20696e76616c6964206465737420636861696e20696400000060448201526064015b60405180910390fd5b6102eb8686610c74565b7fc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61863387878787604051610324969594939291906115d0565b60405180910390a1505050505050565b6066546001600160a01b0316336001600160a01b03161461038e5760405162461bcd60e51b8152602060048201526014602482015273109c9a5919d94e881d5b985d5d1a1bdc9a5cd95960621b60448201526064016102d8565b606680546001600160a01b0319166001600160a01b0392909216919091179055565b6066546001600160a01b0316336001600160a01b03161461040a5760405162461bcd60e51b8152602060048201526014602482015273109c9a5919d94e881d5b985d5d1a1bdc9a5cd95960621b60448201526064016102d8565b8281146104595760405162461bcd60e51b815260206004820152601c60248201527f4272696467653a206d69736d617463682064617461206c656e6774680000000060448201526064016102d8565b60005b838110156104f0578282828181106104765761047661162b565b905060200201602081019061048b919061164f565b606560008787858181106104a1576104a161162b565b90506020020160208101906104b6919061128c565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055806104e88161166c565b91505061045c565b5050505050565b6104ff610d21565b848314801561050d57508281145b6105595760405162461bcd60e51b815260206004820152601a60248201527f4272696467653a20696e76616c696420696e707574206461746100000000000060448201526064016102d8565b60005b838110156107395760008787838181106105785761057861162b565b905060200201602081019061058d919061128c565b6001600160a01b0316141580156105d957506105d9308888848181106105b5576105b561162b565b90506020020160208101906105ca919061128c565b6001600160a01b031690610d7b565b156106b8578686828181106105f0576105f061162b565b9050602002016020810190610605919061128c565b6001600160a01b03166340c10f198686848181106106255761062561162b565b905060200201602081019061063a919061128c565b85858581811061064c5761064c61162b565b6040516001600160e01b031960e087901b1681526001600160a01b0390941660048501526020029190910135602483015250604401600060405180830381600087803b15801561069b57600080fd5b505af11580156106af573d6000803e3d6000fd5b50505050610727565b6107278787838181106106cd576106cd61162b565b90506020020160208101906106e2919061128c565b8686848181106106f4576106f461162b565b9050602002016020810190610709919061128c565b85858581811061071b5761071b61162b565b90506020020135610e55565b806107318161166c565b91505061055c565b507fe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e8686868686866040516103249695949392919061170e565b61077b610d21565b6107856000610f25565b565b600054610100900460ff16158080156107a75750600054600160ff909116105b806107c15750303b1580156107c1575060005460ff166001145b6108245760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102d8565b6000805460ff191660011790558015610847576000805461ff0019166101001790555b6001600160a01b0385161580159061086757506001600160a01b03841615155b6108b35760405162461bcd60e51b815260206004820152601760248201527f4272696467653a20696e76616c6964206164647265737300000000000000000060448201526064016102d8565b6108bc85610f25565b606680546001600160a01b0319166001600160a01b03861617905560005b82811015610949576001606560008686858181106108fa576108fa61162b565b905060200201602081019061090f919061128c565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055806109418161166c565b9150506108da565b5080156104f0576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b61099e610d21565b8281146109ed5760405162461bcd60e51b815260206004820152601a60248201527f4272696467653a20696e76616c696420696e707574206461746100000000000060448201526064016102d8565b60006001600160a01b03861615801590610a155750610a156001600160a01b03871630610d7b565b905060005b84811015610b2b578115610adc57866001600160a01b03166340c10f19878784818110610a4957610a4961162b565b9050602002016020810190610a5e919061128c565b868685818110610a7057610a7061162b565b6040516001600160e01b031960e087901b1681526001600160a01b0390941660048501526020029190910135602483015250604401600060405180830381600087803b158015610abf57600080fd5b505af1158015610ad3573d6000803e3d6000fd5b50505050610b19565b610b1987878784818110610af257610af261162b565b9050602002016020810190610b07919061128c565b86868581811061071b5761071b61162b565b80610b238161166c565b915050610a1a565b507fa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd8686868686604051610324959493929190611788565b46818103610bb35760405162461bcd60e51b815260206004820152601d60248201527f4272696467653a20696e76616c6964206465737420636861696e20696400000060448201526064016102d8565b7fc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c6160003334878787604051610bed969594939291906115d0565b60405180910390a150505050565b610c03610d21565b6001600160a01b038116610c685760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102d8565b610c7181610f25565b50565b6001600160a01b03821660009081526065602052604090205460ff16610cad57610ca96001600160a01b038316333084610f77565b5050565b6001600160a01b0382166379cc6790336040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101849052604401600060405180830381600087803b158015610d0557600080fd5b505af1158015610d19573d6000803e3d6000fd5b505050505050565b6033546001600160a01b031633146107855760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d8565b60408051600481526024810182526020810180516001600160e01b0316638da5cb5b60e01b1790529051600091829182916001600160a01b03871691610dc191906117f0565b600060405180830381855afa9150503d8060008114610dfc576040519150601f19603f3d011682016040523d82523d6000602084013e610e01565b606091505b5091509150811580610e1257508051155b15610e2257600092505050610e4f565b836001600160a01b031681806020019051810190610e40919061180c565b6001600160a01b031614925050505b92915050565b6001600160a01b038316610f0c576000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610eb0576040519150601f19603f3d011682016040523d82523d6000602084013e610eb5565b606091505b5050905080610f065760405162461bcd60e51b815260206004820152601b60248201527f4272696467653a207472616e7366657220657468206661696c6564000000000060448201526064016102d8565b50505050565b610f206001600160a01b0384168383610fe2565b505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040516001600160a01b0380851660248301528316604482015260648101829052610f069085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611012565b6040516001600160a01b038316602482015260448101829052610f2090849063a9059cbb60e01b90606401610fab565b6000611067826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110e79092919063ffffffff16565b90508051600014806110885750808060200190518101906110889190611829565b610f205760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102d8565b60606110f684846000856110fe565b949350505050565b60608247101561115f5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102d8565b600080866001600160a01b0316858760405161117b91906117f0565b60006040518083038185875af1925050503d80600081146111b8576040519150601f19603f3d011682016040523d82523d6000602084013e6111bd565b606091505b50915091506111ce878383876111d9565b979650505050505050565b60608315611248578251600003611241576001600160a01b0385163b6112415760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102d8565b50816110f6565b6110f6838381511561125d5781518083602001fd5b8060405162461bcd60e51b81526004016102d89190611846565b6001600160a01b0381168114610c7157600080fd5b60006020828403121561129e57600080fd5b81356112a981611277565b9392505050565b60008083601f8401126112c257600080fd5b50813567ffffffffffffffff8111156112da57600080fd5b6020830191508360208285010111156112f257600080fd5b9250929050565b60008060008060006080868803121561131157600080fd5b853561131c81611277565b945060208601359350604086013567ffffffffffffffff81111561133f57600080fd5b61134b888289016112b0565b96999598509660600135949350505050565b60008083601f84011261136f57600080fd5b50813567ffffffffffffffff81111561138757600080fd5b6020830191508360208260051b85010111156112f257600080fd5b600080600080604085870312156113b857600080fd5b843567ffffffffffffffff808211156113d057600080fd5b6113dc8883890161135d565b909650945060208701359150808211156113f557600080fd5b506114028782880161135d565b95989497509550505050565b6000806000806000806060878903121561142757600080fd5b863567ffffffffffffffff8082111561143f57600080fd5b61144b8a838b0161135d565b9098509650602089013591508082111561146457600080fd5b6114708a838b0161135d565b9096509450604089013591508082111561148957600080fd5b5061149689828a0161135d565b979a9699509497509295939492505050565b600080600080606085870312156114be57600080fd5b84356114c981611277565b935060208501356114d981611277565b9250604085013567ffffffffffffffff8111156114f557600080fd5b6114028782880161135d565b60008060008060006060868803121561151957600080fd5b853561152481611277565b9450602086013567ffffffffffffffff8082111561154157600080fd5b61154d89838a0161135d565b9096509450604088013591508082111561156657600080fd5b506115738882890161135d565b969995985093965092949392505050565b60008060006040848603121561159957600080fd5b833567ffffffffffffffff8111156115b057600080fd5b6115bc868287016112b0565b909790965060209590950135949350505050565b6001600160a01b038781168252861660208201526040810185905260a06060820181905281018390526000838560c0840137600060c0858401015260c0601f19601f8601168301019050826080830152979650505050505050565b634e487b7160e01b600052603260045260246000fd5b8015158114610c7157600080fd5b60006020828403121561166157600080fd5b81356112a981611641565b60006001820161168c57634e487b7160e01b600052601160045260246000fd5b5060010190565b8183526000602080850194508260005b858110156116d15781356116b681611277565b6001600160a01b0316875295820195908201906001016116a3565b509495945050505050565b81835260006001600160fb1b038311156116f557600080fd5b8260051b80836020870137939093016020019392505050565b6060808252810186905260008760808301825b8981101561175157823561173481611277565b6001600160a01b0316825260209283019290910190600101611721565b50838103602085015261176581888a611693565b915050828103604084015261177b8185876116dc565b9998505050505050505050565b6001600160a01b03861681526060602082018190526000906117ad9083018688611693565b82810360408401526117c08185876116dc565b98975050505050505050565b60005b838110156117e75781810151838201526020016117cf565b50506000910152565b600082516118028184602087016117cc565b9190910192915050565b60006020828403121561181e57600080fd5b81516112a981611277565b60006020828403121561183b57600080fd5b81516112a981611641565b60208152600082518060208401526118658160408501602087016117cc565b601f01601f1916919091016040019291505056fea2646970667358221220aa9df2a3025a5dec16835dde6412bcb301e6630b2662aff905471ba5f68ccfa964736f6c63430008110033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Bridge *BridgeCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Bridge *BridgeSession) ETHTOKEN() (common.Address, error) {
	return _Bridge.Contract.ETHTOKEN(&_Bridge.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Bridge *BridgeCallerSession) ETHTOKEN() (common.Address, error) {
	return _Bridge.Contract.ETHTOKEN(&_Bridge.CallOpts)
}

// BurnableToken is a free data retrieval call binding the contract method 0x07ef210c.
//
// Solidity: function burnableToken(address ) view returns(bool)
func (_Bridge *BridgeCaller) BurnableToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "burnableToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnableToken is a free data retrieval call binding the contract method 0x07ef210c.
//
// Solidity: function burnableToken(address ) view returns(bool)
func (_Bridge *BridgeSession) BurnableToken(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.BurnableToken(&_Bridge.CallOpts, arg0)
}

// BurnableToken is a free data retrieval call binding the contract method 0x07ef210c.
//
// Solidity: function burnableToken(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) BurnableToken(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.BurnableToken(&_Bridge.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Bridge *BridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Bridge *BridgeSession) Operator() (common.Address, error) {
	return _Bridge.Contract.Operator(&_Bridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Bridge *BridgeCallerSession) Operator() (common.Address, error) {
	return _Bridge.Contract.Operator(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x0e93b35c.
//
// Solidity: function bridgeToken(address token, uint256 amount, string externalAddr, uint256 destChainId) returns()
func (_Bridge *BridgeTransactor) BridgeToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "bridgeToken", token, amount, externalAddr, destChainId)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x0e93b35c.
//
// Solidity: function bridgeToken(address token, uint256 amount, string externalAddr, uint256 destChainId) returns()
func (_Bridge *BridgeSession) BridgeToken(token common.Address, amount *big.Int, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeToken(&_Bridge.TransactOpts, token, amount, externalAddr, destChainId)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x0e93b35c.
//
// Solidity: function bridgeToken(address token, uint256 amount, string externalAddr, uint256 destChainId) returns()
func (_Bridge *BridgeTransactorSession) BridgeToken(token common.Address, amount *big.Int, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeToken(&_Bridge.TransactOpts, token, amount, externalAddr, destChainId)
}

// BridgeToken0 is a paid mutator transaction binding the contract method 0xd4546d23.
//
// Solidity: function bridgeToken(string externalAddr, uint256 destChainId) payable returns()
func (_Bridge *BridgeTransactor) BridgeToken0(opts *bind.TransactOpts, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "bridgeToken0", externalAddr, destChainId)
}

// BridgeToken0 is a paid mutator transaction binding the contract method 0xd4546d23.
//
// Solidity: function bridgeToken(string externalAddr, uint256 destChainId) payable returns()
func (_Bridge *BridgeSession) BridgeToken0(externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeToken0(&_Bridge.TransactOpts, externalAddr, destChainId)
}

// BridgeToken0 is a paid mutator transaction binding the contract method 0xd4546d23.
//
// Solidity: function bridgeToken(string externalAddr, uint256 destChainId) payable returns()
func (_Bridge *BridgeTransactorSession) BridgeToken0(externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeToken0(&_Bridge.TransactOpts, externalAddr, destChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address safeMultisigContractAddress, address operator_, address[] tokens) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, safeMultisigContractAddress common.Address, operator_ common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", safeMultisigContractAddress, operator_, tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address safeMultisigContractAddress, address operator_, address[] tokens) returns()
func (_Bridge *BridgeSession) Initialize(safeMultisigContractAddress common.Address, operator_ common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, safeMultisigContractAddress, operator_, tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address safeMultisigContractAddress, address operator_, address[] tokens) returns()
func (_Bridge *BridgeTransactorSession) Initialize(safeMultisigContractAddress common.Address, operator_ common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, safeMultisigContractAddress, operator_, tokens)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Bridge *BridgeTransactor) Mint(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "mint", tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Bridge *BridgeSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Mint(&_Bridge.TransactOpts, tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Bridge *BridgeTransactorSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Mint(&_Bridge.TransactOpts, tokens, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Bridge *BridgeTransactor) Mint0(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "mint0", token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Bridge *BridgeSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Mint0(&_Bridge.TransactOpts, token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Bridge *BridgeTransactorSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Mint0(&_Bridge.TransactOpts, token, recipients, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address operator_) returns()
func (_Bridge *BridgeTransactor) TransferOperator(opts *bind.TransactOpts, operator_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOperator", operator_)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address operator_) returns()
func (_Bridge *BridgeSession) TransferOperator(operator_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOperator(&_Bridge.TransactOpts, operator_)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address operator_) returns()
func (_Bridge *BridgeTransactorSession) TransferOperator(operator_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOperator(&_Bridge.TransactOpts, operator_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x4432e6b7.
//
// Solidity: function updateToken(address[] tokens, bool[] isBurns) returns()
func (_Bridge *BridgeTransactor) UpdateToken(opts *bind.TransactOpts, tokens []common.Address, isBurns []bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateToken", tokens, isBurns)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x4432e6b7.
//
// Solidity: function updateToken(address[] tokens, bool[] isBurns) returns()
func (_Bridge *BridgeSession) UpdateToken(tokens []common.Address, isBurns []bool) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateToken(&_Bridge.TransactOpts, tokens, isBurns)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x4432e6b7.
//
// Solidity: function updateToken(address[] tokens, bool[] isBurns) returns()
func (_Bridge *BridgeTransactorSession) UpdateToken(tokens []common.Address, isBurns []bool) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateToken(&_Bridge.TransactOpts, tokens, isBurns)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeBridgeTokenIterator is returned from FilterBridgeToken and is used to iterate over the raw logs and unpacked data for BridgeToken events raised by the Bridge contract.
type BridgeBridgeTokenIterator struct {
	Event *BridgeBridgeToken // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeToken)
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
		it.Event = new(BridgeBridgeToken)
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
func (it *BridgeBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeToken represents a BridgeToken event raised by the Bridge contract.
type BridgeBridgeToken struct {
	Token       common.Address
	Burner      common.Address
	Amount      *big.Int
	Extddr      string
	DestChainId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeToken is a free log retrieval operation binding the contract event 0xc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61.
//
// Solidity: event BridgeToken(address token, address burner, uint256 amount, string extddr, uint256 destChainId)
func (_Bridge *BridgeFilterer) FilterBridgeToken(opts *bind.FilterOpts) (*BridgeBridgeTokenIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeTokenIterator{contract: _Bridge.contract, event: "BridgeToken", logs: logs, sub: sub}, nil
}

// WatchBridgeToken is a free log subscription operation binding the contract event 0xc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61.
//
// Solidity: event BridgeToken(address token, address burner, uint256 amount, string extddr, uint256 destChainId)
func (_Bridge *BridgeFilterer) WatchBridgeToken(opts *bind.WatchOpts, sink chan<- *BridgeBridgeToken) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeToken)
				if err := _Bridge.contract.UnpackLog(event, "BridgeToken", log); err != nil {
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

// ParseBridgeToken is a log parse operation binding the contract event 0xc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61.
//
// Solidity: event BridgeToken(address token, address burner, uint256 amount, string extddr, uint256 destChainId)
func (_Bridge *BridgeFilterer) ParseBridgeToken(log types.Log) (*BridgeBridgeToken, error) {
	event := new(BridgeBridgeToken)
	if err := _Bridge.contract.UnpackLog(event, "BridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Bridge contract.
type BridgeMintIterator struct {
	Event *BridgeMint // Event containing the contract specifics and raw log

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
func (it *BridgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMint)
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
		it.Event = new(BridgeMint)
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
func (it *BridgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMint represents a Mint event raised by the Bridge contract.
type BridgeMint struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Bridge *BridgeFilterer) FilterMint(opts *bind.FilterOpts) (*BridgeMintIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &BridgeMintIterator{contract: _Bridge.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Bridge *BridgeFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BridgeMint) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMint)
				if err := _Bridge.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseMint(log types.Log) (*BridgeMint, error) {
	event := new(BridgeMint)
	if err := _Bridge.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMint0Iterator is returned from FilterMint0 and is used to iterate over the raw logs and unpacked data for Mint0 events raised by the Bridge contract.
type BridgeMint0Iterator struct {
	Event *BridgeMint0 // Event containing the contract specifics and raw log

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
func (it *BridgeMint0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMint0)
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
		it.Event = new(BridgeMint0)
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
func (it *BridgeMint0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMint0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMint0 represents a Mint0 event raised by the Bridge contract.
type BridgeMint0 struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint0 is a free log retrieval operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Bridge *BridgeFilterer) FilterMint0(opts *bind.FilterOpts) (*BridgeMint0Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return &BridgeMint0Iterator{contract: _Bridge.contract, event: "Mint0", logs: logs, sub: sub}, nil
}

// WatchMint0 is a free log subscription operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Bridge *BridgeFilterer) WatchMint0(opts *bind.WatchOpts, sink chan<- *BridgeMint0) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMint0)
				if err := _Bridge.contract.UnpackLog(event, "Mint0", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseMint0(log types.Log) (*BridgeMint0, error) {
	event := new(BridgeMint0)
	if err := _Bridge.contract.UnpackLog(event, "Mint0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
