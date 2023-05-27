package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tc_bridge/bridgeETH"
	"github.com/tc_bridge/bridgeTC"
	"github.com/tc_bridge/erc20"
	cron "gopkg.in/robfig/cron.v2"
	"log"
	"math/big"
	"strings"
)

// tc section
type Mint struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
}
type Mint2 struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
}
type Burn struct {
	Token   common.Address
	Burner  common.Address
	Amount  *big.Int
	BtcAddr string
}

// eth section
type Deposit struct {
	Token     common.Address
	Sender    common.Address
	Amount    *big.Int
	Recipient common.Address
}

type WithdrawMultiToken struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
}

var ethTokens = make(map[common.Address]bool)
var tcTokens = make(map[common.Address]bool)

type Token struct {
	Address common.Address
	Decimal int
	Symbol  string
}

// todo: load from file/DB
var ETH_TOKEN_LIST = []Token{
	{
		Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		Decimal: 6,
		Symbol:  "USDC",
	},
	{
		Address: common.HexToAddress("0x6982508145454ce325ddbe47a25d4ec3d2311933"),
		Decimal: 18,
		Symbol:  "PEPE",
	},
	{
		Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Decimal: 18,
		Symbol:  "ETH",
	},
}

var TC_TOKEN_LIST = []Token{
	{
		Address: common.HexToAddress("0xA4462aC060A97BCA4032dfbCCAaE9b9419A12F18"),
		Decimal: 18,
		Symbol:  "PEPE",
	},
	{
		Address: common.HexToAddress("0x3ED8040D47133AB8A73Dc41d365578D6e7643E54"),
		Decimal: 18,
		Symbol:  "USDC",
	},
	{
		Address: common.HexToAddress("0x74B033e56434845E02c9bc4F0caC75438033b00D"),
		Decimal: 18,
		Symbol:  "ETH",
	},
	{
		Address: common.HexToAddress("0xfB83c18569fB43f1ABCbae09Baf7090bFFc8CBBD"),
		Decimal: 18,
		Symbol:  "BTC",
	},
}

var ETH_BRIDGE_ADDRESS = common.HexToAddress("0xa103f20367b18d004710141ff505a6b63ce6885c")
var ETH_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000000")

func main() {
	//todo: load last block from file/DB
	var lastETHBlock = 17344391
	var lastTCBlock = 9260
	var stepper = 500

	// load list tokens

	clientETH, err := ethclient.Dial("https://mainnet.infura.io/v3/9ef9bfbcb8a74ad48d473c2036b999a1")
	if err != nil {
		log.Fatal(err)
	}

	clientTC, err := ethclient.Dial("https://tc-node.trustless.computer")
	if err != nil {
		log.Fatal(err)
	}

	scan := func() {
		tempETHBlock, tempTCBlock, err := process(
			stepper,
			clientETH,
			lastETHBlock,
			clientTC,
			lastTCBlock,
		)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			lastTCBlock = tempTCBlock
			lastETHBlock = tempETHBlock
		}
	}
	scan()
	// init cron job
	c := cron.New()
	c.AddFunc("@every 0h10m", scan)
	c.Start()

	messages := make(chan string)
	<-messages

	//c.AddFunc("@daily", func() { fmt.Println("Every day") })
	defer c.Stop()
	defer clientTC.Close()
	defer clientETH.Close()
}

func process(
	stepper int,
	eClient *ethclient.Client,
	lastETHBlock int,
	tcClient *ethclient.Client,
	lastTCBlock int,
) (int, int, error) {
	latestHeightETH, err := eClient.BlockNumber(context.Background())
	if err != nil {
		return 0, 0, err
	}
	totalDepositETH, totalWithdrawETH, err := scanETHBridge(stepper, lastETHBlock, int(latestHeightETH), eClient)
	if err != nil {
		return 0, 0, err
	}
	fmt.Println("deposit eth bridge")
	fmt.Println(totalDepositETH)
	fmt.Println("withdraw eth bridge")
	fmt.Println(totalWithdrawETH)

	latestHeightTC, err := tcClient.BlockNumber(context.Background())
	if err != nil {
		return 0, 0, err
	}
	totalDepositTC, totalWithdrawTC, err := scanTCBridge(stepper, lastTCBlock, int(latestHeightTC), tcClient)
	if err != nil {
		return 0, 0, err
	}
	fmt.Println("mint tc bridge")
	fmt.Println(totalDepositTC)
	fmt.Println("withdraw tc bridge")
	fmt.Println(totalWithdrawTC)

	// query balances
	for _, v := range ETH_TOKEN_LIST {
		if v.Address != ETH_ADDRESS {
			erc20Inst, _ := erc20.NewErc20(v.Address, eClient)
			baln, _ := erc20Inst.BalanceOf(nil, ETH_BRIDGE_ADDRESS)
			fmt.Printf("Balance %s %s \n", baln.String(), v.Symbol)
		} else {
			ethBal, _ := eClient.BalanceAt(context.Background(), ETH_BRIDGE_ADDRESS, nil)
			fmt.Printf("Balance %s %s \n", ethBal.String(), "ETH")
		}
	}

	for _, v := range TC_TOKEN_LIST {
		erc20Inst, _ := erc20.NewErc20(v.Address, tcClient)
		baln, _ := erc20Inst.TotalSupply(nil)
		fmt.Printf("Balance %s %s \n", baln.String(), v.Symbol)
	}
	fmt.Println("----------------------------------------------------------------")

	return int(latestHeightETH), int(latestHeightTC), nil
}

func scanETHBridge(gap int, startBlockETH int, ethBlockLatest int, eClient *ethclient.Client) (map[common.Address]*big.Int, map[common.Address]*big.Int, error) {
	depositEvents := []Deposit{}
	withdrawEvents := []WithdrawMultiToken{}

	logDepositSigHash := crypto.Keccak256Hash([]byte("Deposit(address,address,uint256,address)"))
	logWithdrawSig2Hash := crypto.Keccak256Hash([]byte("Withdraw(address[],address[],uint256[])"))
	logWithdrawSigHash := crypto.Keccak256Hash([]byte("Withdraw(address,address[],uint256[])"))

	// Bridge eth contract
	contractAddress := ETH_BRIDGE_ADDRESS
	endBlock := startBlockETH
	// filter on
	for {
		endBlock += gap
		if endBlock > ethBlockLatest {
			endBlock = ethBlockLatest
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlockETH)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: []common.Address{
				contractAddress,
			},
			Topics: [][]common.Hash{
				{logDepositSigHash, logWithdrawSigHash, logWithdrawSig2Hash},
			},
		}

		logs, err := eClient.FilterLogs(context.Background(), query)
		if err != nil {
			return nil, nil, err
		}

		contractETHAbi, err := abi.JSON(strings.NewReader(bridgeETH.BridgeETHMetaData.ABI))
		if err != nil {
			return nil, nil, err
		}

		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logDepositSigHash:
				var depositEvent Deposit
				err = contractETHAbi.UnpackIntoInterface(&depositEvent, "Deposit", vLog.Data)
				if err != nil {
					return nil, nil, err
				}
				depositEvents = append(depositEvents, depositEvent)
			case logWithdrawSig2Hash:
				var withdraw0Event WithdrawMultiToken
				err = contractETHAbi.UnpackIntoInterface(&withdraw0Event, "Withdraw", vLog.Data)
				if err != nil {
					return nil, nil, err
				}
				withdrawEvents = append(withdrawEvents, withdraw0Event)
			default:
				panic("not handled yet")
			}
		}
		if endBlock == ethBlockLatest {
			break
		}
		startBlockETH = endBlock + 1
	}

	// process data
	totalDepositETH := make(map[common.Address]*big.Int)
	totalWithdrawETH := make(map[common.Address]*big.Int)
	for _, v := range depositEvents {
		if totalDepositETH[v.Token] != nil {
			totalDepositETH[v.Token] = big.NewInt(0).Add(totalDepositETH[v.Token], v.Amount)
		} else {
			totalDepositETH[v.Token] = v.Amount
		}
		if v.Token != ETH_ADDRESS {
			ETH_TOKEN_LIST = addToken(ETH_TOKEN_LIST, v.Token, eClient)
		}
	}
	for _, v := range withdrawEvents {
		for i, v2 := range v.Tokens {
			if totalWithdrawETH[v2] != nil {
				totalWithdrawETH[v2] = big.NewInt(0).Add(totalWithdrawETH[v2], v.Amounts[i])
			} else {
				totalWithdrawETH[v2] = v.Amounts[i]
			}
		}
	}

	return totalDepositETH, totalWithdrawETH, nil
}

func addToken(list []Token, token common.Address, client *ethclient.Client) []Token {
	if !isContains(list, token) {
		erc20Inst, _ := erc20.NewErc20(token, client)
		tokenDec, _ := erc20Inst.Decimals(nil)
		tokenSymbol, _ := erc20Inst.Symbol(nil)
		list = append(list, Token{
			Address: token,
			Decimal: int(tokenDec.Uint64()),
			Symbol:  tokenSymbol,
		})
	}

	return list
}

func isContains(list []Token, newToken common.Address) bool {
	for _, v := range list {
		if v.Address == newToken {
			return true
		}
	}

	return false
}

func scanTCBridge(gap int, startBlockTC int, tcBlockLatest int, tcClient *ethclient.Client) (map[common.Address]*big.Int, map[common.Address]*big.Int, error) {
	mintEvents := []Mint{}
	mint0Events := []Mint2{}
	burnEvents := []Burn{}
	logBurnSigHash := crypto.Keccak256Hash([]byte("Burn(address,address,uint256,string)"))
	logMintSig2Hash := crypto.Keccak256Hash([]byte("Mint(address[],address[],uint256[])"))
	logMintSigHash := crypto.Keccak256Hash([]byte("Mint(address,address[],uint256[])"))

	contractAddressTC := common.HexToAddress("0x63bfaC4D88aeD85E0A0880E501Ed4B9E1D64A47b")
	endBlockTC := startBlockTC
	// filter on
	for {
		endBlockTC += gap
		if endBlockTC > tcBlockLatest {
			endBlockTC = tcBlockLatest
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlockTC)),
			ToBlock:   big.NewInt(int64(endBlockTC)),
			Addresses: []common.Address{
				contractAddressTC,
			},
			Topics: [][]common.Hash{
				{logBurnSigHash, logMintSigHash, logMintSig2Hash},
			},
		}

		logs, err := tcClient.FilterLogs(context.Background(), query)
		if err != nil {
			return nil, nil, err
		}

		contractTCAbi, err := abi.JSON(strings.NewReader(bridgeTC.BridgeTCMetaData.ABI))
		if err != nil {
			return nil, nil, err
		}

		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logMintSigHash:
				var mintEvent Mint
				err = contractTCAbi.UnpackIntoInterface(&mintEvent, "Mint0", vLog.Data)
				if err != nil {
					return nil, nil, err
				}
				mintEvents = append(mintEvents, mintEvent)
			case logMintSig2Hash:
				var mintEvent2 Mint2
				err = contractTCAbi.UnpackIntoInterface(&mintEvent2, "Mint", vLog.Data)
				if err != nil {
					return nil, nil, err
				}
				mint0Events = append(mint0Events, mintEvent2)
			case logBurnSigHash:
				var burnEvent Burn
				err = contractTCAbi.UnpackIntoInterface(&burnEvent, "Burn", vLog.Data)
				if err != nil {
					return nil, nil, err
				}
				burnEvents = append(burnEvents, burnEvent)
			}
		}
		if endBlockTC == tcBlockLatest {
			break
		}
		startBlockTC = endBlockTC + 1
	}

	// process data
	totalDepositETH := make(map[common.Address]*big.Int)
	totalWithdrawETH := make(map[common.Address]*big.Int)
	for _, v := range burnEvents {
		if totalWithdrawETH[v.Token] != nil {
			totalWithdrawETH[v.Token] = big.NewInt(0).Add(totalWithdrawETH[v.Token], v.Amount)
		} else {
			totalWithdrawETH[v.Token] = v.Amount
		}
	}

	for _, v := range mintEvents {
		for _, v2 := range v.Amounts {
			if totalDepositETH[v.Token] != nil {
				totalDepositETH[v.Token] = big.NewInt(0).Add(totalDepositETH[v.Token], v2)
			} else {
				totalDepositETH[v.Token] = v2
			}
		}
		TC_TOKEN_LIST = addToken(TC_TOKEN_LIST, v.Token, tcClient)
	}

	for _, v := range mint0Events {
		for i, v2 := range v.Amounts {
			if totalDepositETH[v.Tokens[i]] != nil {
				totalDepositETH[v.Tokens[i]] = big.NewInt(0).Add(totalDepositETH[v.Tokens[i]], v2)
			} else {
				totalDepositETH[v.Tokens[i]] = v2
			}
			TC_TOKEN_LIST = addToken(TC_TOKEN_LIST, v.Tokens[i], tcClient)
		}
	}

	return totalDepositETH, totalWithdrawETH, nil
}
