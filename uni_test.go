package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tc_bridge/swap"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UniTestSuite struct {
	suite.Suite
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *UniTestSuite) SetupTest() {

}

func TestSolTestSuite(t *testing.T) {
	suite.Run(t, new(UniTestSuite))
}

func (suite *UniTestSuite) TestScanPoolCreated() {
	clientETH, err := ethclient.Dial("https://rpc.mevblocker.io")
	if err != nil {
		log.Fatal(err)
	}

	latestBlockHeight := 18711302
	result, err := ScanPairsCreated(1000, 18709230, latestBlockHeight, common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"), clientETH)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("test result %+v\n", result)
}

func (suite *UniTestSuite) TestParseLiquidityCallData() {
	// mint calldata direct
	mintCallData, _ := hex.DecodeString("8831645600000000000000000000000077e06c9eccf2e797fd462a92b6d7642ef85b0a44000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000002710000000000000000000000000000000000000000000000000000000000002c628000000000000000000000000000000000000000000000000000000000002fc10000000000000000000000000000000000000000000000000000000000ec6735800000000000000000000000000000000000000000000000000872d1c9fa5bc59000000000000000000000000000000000000000000000000000000000ea663290000000000000000000000000000000000000000000000000086021e335edcea000000000000000000000000758a913515dc9db89687dc1727ee25701067800700000000000000000000000000000000000000000000000000000000656d8243")
	increaseLResult, mintResult, err := ParseLiquidityAdded(mintCallData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\ntest result increaseLiquidity %+v\n", increaseLResult)
	fmt.Printf("test result mint %+v\n", mintResult)

	// increase call direct
	increaseLCallData, _ := hex.DecodeString("219f5d17000000000000000000000000000000000000000000000000000000000009661c00000000000000000000000000000000000000000000004e34fbd51e3c8d58d100000000000000000000000000000000000000000000000000000001f0e3351600000000000000000000000000000000000000000000004da7acda3dbcbcd39700000000000000000000000000000000000000000000000000000001ec47fe4800000000000000000000000000000000000000000000000000000000656d8687")
	increaseLResult, mintResult, err = ParseLiquidityAdded(increaseLCallData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\ntest result increaseLiquidity %+v\n", increaseLResult)
	fmt.Printf("test result mint %+v\n", mintResult)

	// multicall data direct
	multicallData, _ := hex.DecodeString("ac9650d800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000016488316456000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000f19308f923582a6f7c465e5ce7a9dc1bec6665b10000000000000000000000000000000000000000000000000000000000002710000000000000000000000000000000000000000000000000000000000002de60000000000000000000000000000000000000000000000000000000000002f05800000000000000000000000000000000000000000000000014dc73db29fd2cb00000000000000000000000000000000000000000008dcb96e90e393ea931eec30000000000000000000000000000000000000000000000001476f90766a6841500000000000000000000000000000000000000000089a62aa693d98e0fb4c72e000000000000000000000000b166b0d15ba559776cb00c1b40269d76d06f7fe400000000000000000000000000000000000000000000000000000000656d833f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000412210e8a00000000000000000000000000000000000000000000000000000000")
	increaseLResult, mintResult, err = ParseLiquidityAdded(multicallData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\ntest result increaseLiquidity %+v\n", increaseLResult)
	fmt.Printf("test result mint %+v\n", mintResult)
}

func (suite *UniTestSuite) TestSwapTokens() {
	clientETH, err := ethclient.Dial("https://goerli.infura.io/v3/3544588d65864af7aaab0e945ec54a01")
	if err != nil {
		log.Fatal(err)
	}

	// key
	privateKey := os.Getenv("PRIVATE_KEY")

	// quote v2 inst
	quotev2, _ := swap.NewQuote(common.HexToAddress(QUOTE_V2), clientETH)

	// swap ETH -> UNI
	paths1 := []common.Address{common.HexToAddress(WETH), common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984")}
	fees1 := []int64{500}
	swapAmount1 := big.NewInt(1e14)
	minimumOut, _ := GetTokenPrice(quotev2, paths1[0], paths1[1], swapAmount1, big.NewInt(fees1[0]))

	tx1, err := SwapTokens(clientETH, paths1, fees1, swapAmount1, minimumOut, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx1: %v\n", tx1.Hash().String())

	// swap UNI -> ETH
	paths2 := []common.Address{common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"), common.HexToAddress(WETH)}
	fees2 := []int64{500}
	swapAmount2 := big.NewInt(1e12)
	minimumOut2, _ := GetTokenPrice(quotev2, paths2[0], paths2[1], swapAmount2, big.NewInt(fees2[0]))

	tx2, err := SwapTokens(clientETH, paths2, fees2, swapAmount2, minimumOut2, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx2: %v\n", tx2.Hash().String())
}
