package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tc_bridge/bridgeETH"
	"github.com/tc_bridge/bridgeTC"
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

func main() {

	// init cron job
	s := cron.New()

}

// scan eth

// scan tc
func scanTCBridge() ([]Deposit, []WithdrawMultiToken, error) {
	depositEvents := []Deposit{}
	withdrawEvents := []WithdrawMultiToken{}

	startBlockETH := 17229013
	clientETH, err := ethclient.Dial("https://mainnet.infura.io/v3/9ef9bfbcb8a74ad48d473c2036b999a1")
	if err != nil {
		log.Fatal(err)
	}

	ethBlockLatest, err := clientETH.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	logDepositSigHash := crypto.Keccak256Hash([]byte("Deposit(address,address,uint256,address)"))
	logWithdrawSig2Hash := crypto.Keccak256Hash([]byte("Withdraw(address[],address[],uint256[])"))
	logWithdrawSigHash := crypto.Keccak256Hash([]byte("Withdraw(address,address[],uint256[])"))

	// Bridge eth contract
	contractAddress := common.HexToAddress("0xa103f20367b18d004710141ff505a6b63ce6885c")
	gap := 500
	endBlock := startBlockETH
	// filter on
	for {
		endBlock += gap
		if endBlock > int(ethBlockLatest) {
			endBlock = int(ethBlockLatest)
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

		logs, err := clientETH.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		contractETHAbi, err := abi.JSON(strings.NewReader(bridgeETH.BridgeETHMetaData.ABI))
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logDepositSigHash:
				var depositEvent Deposit
				err = contractETHAbi.UnpackIntoInterface(&depositEvent, "Deposit", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				depositEvents = append(depositEvents, depositEvent)
			case logWithdrawSig2Hash:
				var withdraw0Event WithdrawMultiToken
				err = contractETHAbi.UnpackIntoInterface(&withdraw0Event, "Withdraw", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				withdrawEvents = append(withdrawEvents, withdraw0Event)
			default:
				panic("not handled yet")
			}
		}
		if endBlock == int(ethBlockLatest) {
			break
		}
		startBlockETH = endBlock + 1
	}
}

func scanETHBridge(gap int, startBlockTC int) ([]Mint, []Mint2, []Burn, error) {
	clientTC, err := ethclient.Dial("https://tc-node.trustless.computer")
	if err != nil {
		return nil, nil, nil, err
	}

	mintEvents := []Mint{}
	mint0Events := []Mint2{}
	burnEvents := []Burn{}
	tcBlockLatest, err := clientTC.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	logBurnSigHash := crypto.Keccak256Hash([]byte("Burn(address,address,uint256,string)"))
	logMintSig2Hash := crypto.Keccak256Hash([]byte("Mint(address[],address[],uint256[])"))
	logMintSigHash := crypto.Keccak256Hash([]byte("Mint(address,address[],uint256[])"))

	contractAddressTC := common.HexToAddress("0x63bfaC4D88aeD85E0A0880E501Ed4B9E1D64A47b")
	endBlockTC := startBlockTC
	// filter on
	for {
		endBlockTC += gap
		if endBlockTC > int(tcBlockLatest) {
			endBlockTC = int(tcBlockLatest)
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlockTC)),
			ToBlock:   big.NewInt(int64(endBlockTC)),
			Addresses: []common.Address{
				contractAddressTC,
			},
			Topics: [][]common.Hash{
				{},
			},
		}

		logs, err := clientTC.FilterLogs(context.Background(), query)
		if err != nil {
			return nil, nil, nil, err
		}

		contractTCAbi, err := abi.JSON(strings.NewReader(bridgeTC.BridgeTCMetaData.ABI))
		if err != nil {
			return nil, nil, nil, err
		}

		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logMintSigHash:
				var mintEvent Mint
				err = contractTCAbi.UnpackIntoInterface(&mintEvent, "Mint0", vLog.Data)
				if err != nil {
					return nil, nil, nil, err
				}
				mintEvents = append(mintEvents, mintEvent)
			case logMintSig2Hash:
				var mintEvent2 Mint2
				err = contractTCAbi.UnpackIntoInterface(&mintEvent2, "Mint", vLog.Data)
				if err != nil {
					return nil, nil, nil, err
				}
				mint0Events = append(mint0Events, mintEvent2)
			case logBurnSigHash:
				var burnEvent Burn
				err = contractTCAbi.UnpackIntoInterface(&burnEvent, "Burn", vLog.Data)
				if err != nil {
					return nil, nil, nil, err
				}
				burnEvents = append(burnEvents, burnEvent)
			}
		}
		if endBlockTC == int(tcBlockLatest) {
			break
		}
		startBlockTC = endBlockTC + 1
	}

	return mintEvents, mint0Events, burnEvents, nil
}
