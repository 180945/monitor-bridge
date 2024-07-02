package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tc_bridge/bridgeETH"
	"github.com/tc_bridge/bridgeTC"
	"github.com/tc_bridge/erc20"
	"github.com/tc_bridge/slack"
	"github.com/tc_bridge/swap"
	l2 "github.com/tc_bridge/zksyncBridge"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
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

type Swap struct {
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
}

type SwapFee struct {
	Amount0 *big.Int
	Amount1 *big.Int
}

var ethTokens = make(map[common.Address]*Token)
var tcTokens = make(map[common.Address]*Token)

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
var L1_MESSENGER_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000008008")

type Proof struct {
	L2BatchNumber     *big.Int
	L2MessageIndex    *big.Int
	L2TxNumberInBatch uint16
	Message           []byte
	Paths             [][32]byte
}

var slackInst *slack.Slack

func main() {

	client, err := ethclient.Dial("https://rpc.bvm-7132.l2aas.com")
	if err != nil {
		log.Fatal(err)
	}

	//clientL2, err := ethclient.Dial("https://mainnet.era.zksync.io")
	//if err != nil {
	//	log.Fatal(err)
	//}

	receiverAddr := common.HexToAddress("dac17f958d2ee523a2206206994597c13d831ec7") // todo: replace
	bitcoinAddr := common.HexToAddress("REPLACE_HERE")
	depositAmount := big.NewInt(10)

	startTime := time.Now().String()
	var wg sync.WaitGroup
	var keys = []string{}

	// fund process
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	wallet.PrivateKey(account)

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559

	sendAmount := "200"
	zksyncBridgeL1, err := l2.NewBridgeL1(common.HexToAddress("0x124287BfEbf1d2a9D49835b6697c81c296b7B206"), client)
	if err != nil {
		log.Fatal(err)
	}

	txCounts := []uint64{}
	for i := 0; i < len(keys); i++ {
		txCounts = append(txCounts, 0)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n start time")
		fmt.Println(startTime)
		fmt.Println("end time")
		fmt.Println(time.Now().String())
		arrSum := 0
		for i := 0; i < len(txCounts); i++ {
			arrSum = arrSum + int(txCounts[i])
		}
		fmt.Printf("total txs: %v \n", arrSum)
		// Run Cleanup
		os.Exit(1)
	}()

	for i, key := range keys {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(key string, index int) {
			defer wg.Done()

			privKey, err := crypto.HexToECDSA(key)
			if err != nil {
				log.Fatal(err)
			}

			auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(84163))
			if err != nil {
				log.Fatal(err)
			}

			// deposit
			amount := big.NewInt(0)
			amount.SetString(sendAmount, 10)
			auth.Value = amount
			auth.GasPrice = big.NewInt(1e9)
			auth.GasLimit = 500_000
			for {
				depositTx, err := zksyncBridgeL1.Deposit0(
					auth,
					receiverAddr,
					bitcoinAddr,
					depositAmount,
					big.NewInt(1e6),
					big.NewInt(800),
					auth.From,
				)
				if err != nil {
					fmt.Println(auth.From.String())
					fmt.Println(err.Error())
					fmt.Println("RETRY")
				} else {
					txCounts[index]++
					fmt.Println(depositTx.Hash().String())
				}

				//break
			}

		}(key, i)
	}
	wg.Wait()
}

func checkTxReadyToFinalized(clientL2 *ethclient.Client, contactL1 *l2.L1, txHash common.Hash) (bool, *Proof, error) {
	receipt, err := clientL2.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return false, nil, err
	}

	if receipt == nil {
		return false, nil, errors.New("invalid withdraw tx")
	}

	var logRes2 *types.Log
	for _, logValue := range receipt.Logs {
		if logValue.Address == L1_MESSENGER_ADDRESS && logValue.Topics[0] == crypto.Keccak256Hash([]byte("L1MessageSent(address,bytes32,bytes)")) {
			logRes2 = logValue
			break
		}
	}
	fmt.Println(logRes2)

	r := make(map[string]interface{})
	err = clientL2.Client().Call(&r, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return false, nil, err
	}

	index := new(big.Int)
	logL1ToL2Map := r["l2ToL1Logs"].([]interface{})
	for _, v := range logL1ToL2Map {
		vmap := v.(map[string]interface{})
		sender := common.HexToAddress(vmap["sender"].(string))
		if sender == L1_MESSENGER_ADDRESS {
			index.SetString(vmap["logIndex"].(string)[2:], 16)
			break
		}
	}

	var proofRes map[string]interface{}
	err = clientL2.Client().Call(&proofRes, "zks_getL2ToL1LogProof", txHash, index)
	if err != nil {
		return false, nil, err
	}

	l1BatchNumber := new(big.Int)
	l1BatchNumber.SetString(r["l1BatchNumber"].(string)[2:], 16)
	l1BatchTxIndex := new(big.Int)
	l1BatchTxIndex.SetString(r["l1BatchTxIndex"].(string)[2:], 16)
	proofId := big.NewInt(int64(proofRes["id"].(float64)))
	l2ChainId, err := clientL2.ChainID(context.Background())
	if err != nil {
		return false, nil, err
	}

	isFinalized, err := contactL1.IsWithdrawalFinalized(nil, l2ChainId, l1BatchNumber, proofId)
	if err != nil {
		return false, nil, err
	}

	if isFinalized {
		return true, nil, nil
	}

	paths := [][32]byte{}
	temp := [32]byte{}
	for _, v := range proofRes["proof"].([]interface{}) {
		vBytes, err := hex.DecodeString(v.(string)[2:])
		if err != nil {
			return false, nil, err
		}
		copy(temp[:], vBytes[0:32])
		paths = append(paths, temp)
	}

	//decodeBytes
	messageData := logRes2.Data
	if len(messageData) >= 64 {
		messageData = logRes2.Data[64 : logRes2.Data[63]+64]
	}

	// proof
	proof := &Proof{
		L2BatchNumber:     l1BatchNumber,
		L2MessageIndex:    proofId,
		L2TxNumberInBatch: uint16(l1BatchTxIndex.Uint64()),
		Message:           messageData,
		Paths:             paths,
	}

	return false, proof, nil
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
	// handle last block height = 0
	if lastETHBlock == 0 {
		lastETHBlock = int(latestHeightETH) - 3600 // 3600 ~ 12hours
	}
	totalDepositETH, totalWithdrawETH, err := scanETHBridge(stepper, lastETHBlock, int(latestHeightETH), eClient)
	if err != nil {
		return 0, 0, err
	}

	latestHeightTC, err := tcClient.BlockNumber(context.Background())
	if err != nil {
		return 0, 0, err
	}
	// handle last block height = 0
	if lastTCBlock == 0 {
		lastTCBlock = int(latestHeightTC) - 73 // 73 ~ 12hours
	}

	totalDepositTC, totalWithdrawTC, err := scanTCBridge(stepper, lastTCBlock, int(latestHeightTC), tcClient)
	if err != nil {
		return 0, 0, err
	}

	var result string
	// notify to slack 12h mint burn
	result += formatOutput(totalDepositETH, totalWithdrawETH, "eth", ethTokens)
	result += formatOutput(totalDepositTC, totalWithdrawTC, "tc", tcTokens)

	result += "\n"
	// query balances
	result += formatOuputBalances("eth", ethTokens, eClient, ETH_BRIDGE_ADDRESS)
	result += formatOuputBalances("tc", tcTokens, tcClient, common.Address{})

	channelID := os.Getenv("SLACK_CHANNEL_ID") // todo
	if _, _, err := slackInst.SendMessageToSlackWithChannel(channelID, "Bridge monitor", "Process", result); err != nil {
		fmt.Println(channelID)
		fmt.Println("slackInst.SendMessageToSlackWithChannel err", err)
		return 0, 0, err
	}
	// fmt.Println(result)

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
		if startBlockETH > endBlock {
			break
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
		if v.Token != ETH_ADDRESS && ethTokens[v.Token] == nil {
			ethTokens[v.Token] = addToken(v.Token, eClient)
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

func addToken(token common.Address, client *ethclient.Client) *Token {
	erc20Inst, _ := erc20.NewErc20(token, client)
	tokenDec, _ := erc20Inst.Decimals(nil)
	tokenSymbol, _ := erc20Inst.Symbol(nil)

	return &Token{
		Address: token,
		Decimal: int(tokenDec.Uint64()),
		Symbol:  tokenSymbol,
	}
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
		if startBlockTC > endBlockTC {
			break
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
		if tcTokens[v.Token] == nil {
			tcTokens[v.Token] = addToken(v.Token, tcClient)
		}
	}

	for _, v := range mint0Events {
		for i, v2 := range v.Amounts {
			if totalDepositETH[v.Tokens[i]] != nil {
				totalDepositETH[v.Tokens[i]] = big.NewInt(0).Add(totalDepositETH[v.Tokens[i]], v2)
			} else {
				totalDepositETH[v.Tokens[i]] = v2
			}
			if tcTokens[v.Tokens[i]] == nil {
				tcTokens[v.Tokens[i]] = addToken(v.Tokens[i], tcClient)
			}
		}
	}

	return totalDepositETH, totalWithdrawETH, nil
}

func formatOutput(deposits map[common.Address]*big.Int, withdraws map[common.Address]*big.Int, prefix string, tokenList map[common.Address]*Token) string {
	var notifyMintBurn string
	if prefix == "eth" {
		notifyMintBurn = "\nETHEREUM BRIDGE 12h\n"
	} else if prefix == "tc" {
		notifyMintBurn = "\nTC BRIDGE 12h\n"
	} else {
		panic("invalid prefix")
	}
	notifyMintBurn += "Deposit Volume: \n"
	if len(deposits) > 0 {
		for a, v := range deposits {
			temp1, _ := big.NewFloat(0).SetString(v.String())
			temp := big.NewFloat(0).Quo(temp1, big.NewFloat(math.Pow(10, float64(tokenList[a].Decimal))))
			notifyMintBurn += fmt.Sprintf("%v %v \n", temp.Text('f', 4), tokenList[a].Symbol)
		}
	} else {
		notifyMintBurn += "$0\n"
	}
	notifyMintBurn += "Withdraw Volume: \n"
	if len(withdraws) > 0 {
		for a, v := range withdraws {
			temp1, _ := big.NewFloat(0).SetString(v.String())
			temp := big.NewFloat(0).Quo(temp1, big.NewFloat(math.Pow(10, float64(tokenList[a].Decimal))))
			notifyMintBurn += fmt.Sprintf("%v %v \n", temp.Text('f', 4), tokenList[a].Symbol)
		}
	} else {
		notifyMintBurn += "$0\n"
	}

	return notifyMintBurn
}

func formatOuputBalances(prefix string, tokenList map[common.Address]*Token, clientInst *ethclient.Client, contractAddr common.Address) string {
	var notifyBalances string
	if prefix == "eth" {
		notifyBalances = "\nETHEREUM BRIDGE BALANCES \n"
	} else if prefix == "tc" {
		notifyBalances = "\nTC BRIDGE BALANCES \n"
	} else {
		panic("invalid prefix")
	}

	for _, v := range tokenList {
		var bal *big.Int
		if prefix == "eth" {
			if v.Address != ETH_ADDRESS {
				erc20Inst, _ := erc20.NewErc20(v.Address, clientInst)
				bal, _ = erc20Inst.BalanceOf(nil, contractAddr)

			} else {
				bal, _ = clientInst.BalanceAt(context.Background(), contractAddr, nil)
			}
		} else {
			erc20Inst, _ := erc20.NewErc20(v.Address, clientInst)
			bal, _ = erc20Inst.TotalSupply(nil)
		}
		temp1, _ := big.NewFloat(0).SetString(bal.String())
		temp := big.NewFloat(0).Quo(temp1, big.NewFloat(math.Pow(10, float64(v.Decimal))))
		notifyBalances += fmt.Sprintf("%s %s \n", temp.Text('f', 4), v.Symbol)
	}

	return notifyBalances
}

type Data struct {
	TxHash    string `json:"tx_hash"`
	Address   string `json:"address"`
	FromToken string `json:"from_token"`
	ToToken   string `json:"to_token"`
	Amount    string `json:"amount"`
}

const BTC = "BTC"
const ETH = "ETH"

func collectSwapKeyData(gap int, startBlockETH int, ethBlockLatest int, contractAddress common.Address, eClient *ethclient.Client) {
	logSwapHash := crypto.Keccak256Hash([]byte("Swap(address,address,int256,int256,uint160,uint128,int24)"))
	endBlock := ethBlockLatest
	swapEvents := []Swap{}
	totalVolume := big.NewInt(0)
	jsonDatas := []Data{}
	for {
		endBlock += gap
		if endBlock > ethBlockLatest {
			endBlock = ethBlockLatest
		}
		if startBlockETH > endBlock {
			break
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlockETH)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: []common.Address{
				contractAddress,
			},
			Topics: [][]common.Hash{
				{logSwapHash},
			},
		}

		logs, err := eClient.FilterLogs(context.Background(), query)
		if err != nil {
			panic(err)
		}

		contractETHAbi, err := abi.JSON(strings.NewReader(swap.SwapMetaData.ABI))
		if err != nil {
			panic(err)
		}
		mapData := make(map[string]bool)
		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logSwapHash:
				var swapEvent Swap
				err = contractETHAbi.UnpackIntoInterface(&swapEvent, "Swap", vLog.Data)
				if err != nil {
					panic(err)
				}
				swapEvents = append(swapEvents, swapEvent)

				fromToken := BTC
				toToken := ETH
				amount := swapEvent.Amount0
				if swapEvent.Amount0.Cmp(big.NewInt(0)) < 0 {
					fromToken = ETH
					toToken = BTC
					amount = swapEvent.Amount1
				}
				if !mapData[vLog.TxHash.String()] {
					mapData[vLog.TxHash.String()] = true
				} else {
					continue
				}
				// add to data
				jsonDatas = append(jsonDatas, Data{
					TxHash:    vLog.TxHash.String(),
					Address:   common.HexToAddress(vLog.Topics[2].String()).String(),
					FromToken: fromToken,
					ToToken:   toToken,
					Amount:    amount.String(),
				})
			}
		}
		if endBlock == ethBlockLatest {
			break
		}
		startBlockETH = endBlock + 1
	}

	for _, event := range swapEvents {
		if event.Amount1.Cmp(big.NewInt(0)) > 0 {
			totalVolume = big.NewInt(0).Add(totalVolume, event.Amount1)
		} else {
			totalVolume = big.NewInt(0).Sub(totalVolume, event.Amount1)
		}
	}

	fmt.Printf("%+v \n", swapEvents)
	fmt.Println("Volume 24h")
	fmt.Printf("Total tx %v \n", len(swapEvents))
	fmt.Printf("Total volume %v ETH \n", totalVolume.String())

	file, _ := json.MarshalIndent(jsonDatas, "", " ")
	_ = ioutil.WriteFile("swaps.json", file, 0644)
}

func collectFeeSwapData(gap int, startBlockETH int, ethBlockLatest int, contractAddress common.Address, eClient *ethclient.Client) {
	logSwapHash := crypto.Keccak256Hash([]byte("CollectProtocol(address,address,uint128,uint128)"))
	fmt.Println(logSwapHash.String())
	endBlock := ethBlockLatest
	feeBTC := big.NewInt(0)
	feeETH := big.NewInt(0)
	for {
		endBlock += gap
		if endBlock > ethBlockLatest {
			endBlock = ethBlockLatest
		}
		if startBlockETH > endBlock {
			break
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlockETH)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: []common.Address{
				contractAddress,
			},
			Topics: [][]common.Hash{
				{logSwapHash},
			},
		}

		logs, err := eClient.FilterLogs(context.Background(), query)
		if err != nil {
			panic(err)
		}

		contractETHAbi, err := abi.JSON(strings.NewReader(swap.SwapMetaData.ABI))
		if err != nil {
			panic(err)
		}
		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logSwapHash:
				var swapEvent SwapFee
				err = contractETHAbi.UnpackIntoInterface(&swapEvent, "CollectProtocol", vLog.Data)
				if err != nil {
					panic(err)
				}

				recipient := common.BytesToAddress(vLog.Topics[2].Bytes())
				fmt.Println(recipient.String())

				feeBTC = big.NewInt(0).Add(feeBTC, swapEvent.Amount0)
				feeETH = big.NewInt(0).Add(feeETH, swapEvent.Amount1)
			}
		}
		if endBlock == ethBlockLatest {
			break
		}
		startBlockETH = endBlock + 1
	}

	fmt.Printf("Fee ETH %v \n", feeETH.String())
	fmt.Printf("Fee BTC %v \n", feeBTC.String())
}

////////////////////////////////////////////////////////////////

/// @dev Support uniswap bot

type PoolCreatedData struct {
	TickSpacing *big.Int       `json:"tickSpacing"`
	Pool        common.Address `json:"pool"`
}

type PoolCreated struct {
	Token0 common.Address `json:"token0"`
	Token1 common.Address `json:"token1"`
	Fee    *big.Int       `json:"fee"`
	PoolCreatedData
}

// Scan new pairs
func ScanPairsCreated(gap int, startBlockETH int, ethBlockLatest int, contractAddress common.Address, eClient *ethclient.Client) ([]PoolCreated, error) {
	poolCreatedHash := crypto.Keccak256Hash([]byte("PoolCreated(address,address,uint24,int24,address)"))
	endBlock := ethBlockLatest
	pools := []PoolCreated{}

	for {
		endBlock += gap
		if endBlock > ethBlockLatest {
			endBlock = ethBlockLatest
		}
		if startBlockETH > endBlock {
			break
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlockETH)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: []common.Address{
				contractAddress,
			},
			Topics: [][]common.Hash{
				{poolCreatedHash},
			},
		}

		logs, err := eClient.FilterLogs(context.Background(), query)
		if err != nil {
			return pools, err
		}

		contractETHAbi, err := abi.JSON(strings.NewReader(swap.FactoryMetaData.ABI))
		if err != nil {
			return pools, err
		}
		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case poolCreatedHash:
				var poolCreatedData PoolCreatedData
				err = contractETHAbi.UnpackIntoInterface(&poolCreatedData, "PoolCreated", vLog.Data)
				if err != nil {
					return pools, err
				}

				if len(vLog.Topics) < 4 {
					continue
				}

				token0 := common.BytesToAddress(vLog.Topics[1].Bytes())
				token1 := common.BytesToAddress(vLog.Topics[2].Bytes())
				fee := big.NewInt(0).SetBytes(vLog.Topics[3].Bytes())

				pools = append(pools, PoolCreated{
					Token0:          token0,
					Token1:          token1,
					Fee:             fee,
					PoolCreatedData: poolCreatedData,
				})

			}
		}
		if endBlock == ethBlockLatest {
			break
		}
		startBlockETH = endBlock + 1
	}

	return pools, nil
}

// Parse add liquidity data
func ParseLiquidityAdded(data []byte) ([]swap.INonfungiblePositionManagerIncreaseLiquidityParams, []swap.INonfungiblePositionManagerMintParams, error) {
	// switch case below
	increaseLiquidity := []swap.INonfungiblePositionManagerIncreaseLiquidityParams{}
	mintParams := []swap.INonfungiblePositionManagerMintParams{}

	parsedABI, err := abi.JSON(strings.NewReader(swap.NonfungiblePositionManagerMetaData.ABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	methodById, err := parsedABI.MethodById(data[0:4])
	if err != nil {
		return increaseLiquidity, mintParams, err
	}

	var calldatas [][]byte
	if methodById.Name == "multicall" {
		var args = make(map[string]interface{})
		err = methodById.Inputs.UnpackIntoMap(args, data[4:])
		if err != nil {
			fmt.Printf("Failed to decode multicall ABI: %v", err)
			return increaseLiquidity, mintParams, err
		}
		calldatas = args["data"].([][]byte)
	} else {
		calldatas = append(calldatas, data)
	}

	for _, v := range calldatas {
		// decode with mint
		sliceMethodIDOut := v[4:]
		methodById, err = parsedABI.MethodById(v[0:4])
		if err != nil {
			continue
		}

		if methodById.Name == "mint" {
			var tempMint swap.INonfungiblePositionManagerMintParams
			var args = make(map[string]interface{})
			err = methodById.Inputs.UnpackIntoMap(args, sliceMethodIDOut)
			if err != nil {
				fmt.Printf("\nFailed to decode increaseLiquidity ABI: %v", err)
				continue
			} else {
				mintDataJson, _ := json.Marshal(args["params"])
				if err := json.Unmarshal(mintDataJson, &tempMint); err != nil {
					fmt.Println(err)
					continue
				}

				mintParams = append(mintParams, tempMint)
			}
		} else if methodById.Name == "increaseLiquidity" {
			var tempIncreaseLiquidity swap.INonfungiblePositionManagerIncreaseLiquidityParams
			var args = make(map[string]interface{})
			err = methodById.Inputs.UnpackIntoMap(args, sliceMethodIDOut)
			if err != nil {
				fmt.Printf("\nFailed to decode increaseLiquidity ABI: %v", err)
				continue
			} else {
				mintDataJson, _ := json.Marshal(args["params"])
				if err := json.Unmarshal(mintDataJson, &tempIncreaseLiquidity); err != nil {
					fmt.Println(err)
					continue
				}
				increaseLiquidity = append(increaseLiquidity, tempIncreaseLiquidity)
			}
		}
	}

	return increaseLiquidity, mintParams, nil
}

const ROUTER_MAINNET = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
const WETH = "0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6"
const QUOTE_V2 = "0x61fFE014bA17989E743c5F6cB21bF9697530B21e"

// swap
func SwapTokens(clientInst *ethclient.Client, paths []common.Address, fees []int64, amount *big.Int, minimumAmountOut *big.Int, privateKeyStr string) (*types.Transaction, error) {
	// validate inputs
	if len(paths) < 2 || len(paths) != len(fees)+1 {
		return nil, fmt.Errorf("invalid inputs")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	chainId, err := clientInst.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	// approve token
	tokenInput, err := erc20.NewErc20(paths[0], clientInst)
	if err != nil {
		return nil, err
	}

	routerAddr := common.HexToAddress(ROUTER_MAINNET)
	weth := common.HexToAddress(WETH)

	if paths[0] != weth {
		// check allowance
		allowanceBal, err := tokenInput.Allowance(nil, auth.From, routerAddr)
		if err != nil {
			return nil, err
		}

		if allowanceBal.Cmp(amount) < 0 {
			_, err = tokenInput.Approve(auth, routerAddr, amount)
			if err != nil {
				return nil, err
			}
		}
	}

	// swap
	routerInst, err := swap.NewRouter(routerAddr, clientInst)
	if err != nil {
		return nil, err
	}

	swapPath, err := BuildPath(paths, fees)
	if err != nil {
		return nil, err
	}

	var calldata [][]byte
	// build calldata
	routerABI, err := abi.JSON(strings.NewReader(swap.RouterMetaData.ABI))
	if err != nil {
		return nil, err
	}

	exactInputData := swap.ISwapRouterExactInputParams{
		Path:             swapPath,
		Recipient:        auth.From,
		Deadline:         big.NewInt(time.Now().Unix() + 1000),
		AmountIn:         amount,
		AmountOutMinimum: minimumAmountOut,
	}

	var unwrapEth []byte
	// handle eth
	if paths[0] == weth {
		auth.Value = amount
	} else if paths[len(paths)-1] == weth {
		unwrapEth, err = routerABI.Pack("unwrapWETH9", minimumAmountOut, auth.From)
		if err != nil {
			return nil, err
		}
		exactInputData.Recipient = common.Address{}
	}

	callSwapData, err := routerABI.Pack("exactInput", exactInputData)
	if err != nil {
		return nil, err
	}

	calldata = append(calldata, callSwapData)
	if len(unwrapEth) > 0 {
		calldata = append(calldata, unwrapEth)
	}

	return routerInst.Multicall(auth, calldata)
}

// Get price of a pair
func GetTokenPrice(quoteV2Inst *swap.Quote, fromToken common.Address, toToken common.Address, amount *big.Int, fee *big.Int) (*big.Int, error) {
	quoteSingle := swap.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           fromToken,
		TokenOut:          toToken,
		AmountIn:          amount,
		Fee:               fee,
		SqrtPriceLimitX96: big.NewInt(0),
	}

	output, err := quoteV2Inst.QuoteExactInputSingle(nil, quoteSingle)

	if err != nil {
		return nil, err
	}

	return output.AmountOut, nil
}

func BuildPath(paths []common.Address, fees []int64) ([]byte, error) {
	var temp []byte
	for i := 0; i < len(fees); i++ {
		temp = append(temp, paths[i].Bytes()...)
		temp1 := fmt.Sprintf("%06x", fees[i])
		fee, err := hex.DecodeString(temp1)
		if err != nil {
			return temp, err
		}
		temp = append(temp, fee...)
	}
	temp = append(temp, paths[len(paths)-1].Bytes()...)

	return temp, nil
}
