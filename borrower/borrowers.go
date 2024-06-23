// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package borrowers

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

// BorrowersMetaData contains all meta data concerning the Borrowers contract.
var BorrowersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GrantBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractICompliance\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NewCompliance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldHYController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newHYController\",\"type\":\"address\"}],\"name\":\"NewHYController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPassivePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPassivePool\",\"type\":\"address\"}],\"name\":\"NewPassivePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TokensFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TokensUnfrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"}],\"name\":\"__HYToken_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractControllerInterface\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"_setHYController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPassivePool\",\"name\":\"newPassivePool\",\"type\":\"address\"}],\"name\":\"_setPassivePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"badDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claims\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"forcedTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"freezePartialTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"grantBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"grantBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hyController\",\"outputs\":[{\"internalType\":\"contractControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingToken_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"contractHYController\",\"name\":\"hyController_\",\"type\":\"address\"},{\"internalType\":\"contractIPassivePool\",\"name\":\"passivePool_\",\"type\":\"address\"},{\"internalType\":\"contractRewardsInterface\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolOwner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passivePool\",\"outputs\":[{\"internalType\":\"contractIPassivePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_freeze\",\"type\":\"bool\"}],\"name\":\"setAddressFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"compliance_\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAccruing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAccruingInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"supplyFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCanClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unfreezePartialTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unwindClaimedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unwindAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badDebt_\",\"type\":\"uint256\"}],\"name\":\"unwindCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BorrowersABI is the input ABI used to generate the binding from.
// Deprecated: Use BorrowersMetaData.ABI instead.
var BorrowersABI = BorrowersMetaData.ABI

// Borrowers is an auto generated Go binding around an Ethereum contract.
type Borrowers struct {
	BorrowersCaller     // Read-only binding to the contract
	BorrowersTransactor // Write-only binding to the contract
	BorrowersFilterer   // Log filterer for contract events
}

// BorrowersCaller is an auto generated read-only Go binding around an Ethereum contract.
type BorrowersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BorrowersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BorrowersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BorrowersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BorrowersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BorrowersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BorrowersSession struct {
	Contract     *Borrowers        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BorrowersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BorrowersCallerSession struct {
	Contract *BorrowersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BorrowersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BorrowersTransactorSession struct {
	Contract     *BorrowersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BorrowersRaw is an auto generated low-level Go binding around an Ethereum contract.
type BorrowersRaw struct {
	Contract *Borrowers // Generic contract binding to access the raw methods on
}

// BorrowersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BorrowersCallerRaw struct {
	Contract *BorrowersCaller // Generic read-only contract binding to access the raw methods on
}

// BorrowersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BorrowersTransactorRaw struct {
	Contract *BorrowersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBorrowers creates a new instance of Borrowers, bound to a specific deployed contract.
func NewBorrowers(address common.Address, backend bind.ContractBackend) (*Borrowers, error) {
	contract, err := bindBorrowers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Borrowers{BorrowersCaller: BorrowersCaller{contract: contract}, BorrowersTransactor: BorrowersTransactor{contract: contract}, BorrowersFilterer: BorrowersFilterer{contract: contract}}, nil
}

// NewBorrowersCaller creates a new read-only instance of Borrowers, bound to a specific deployed contract.
func NewBorrowersCaller(address common.Address, caller bind.ContractCaller) (*BorrowersCaller, error) {
	contract, err := bindBorrowers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BorrowersCaller{contract: contract}, nil
}

// NewBorrowersTransactor creates a new write-only instance of Borrowers, bound to a specific deployed contract.
func NewBorrowersTransactor(address common.Address, transactor bind.ContractTransactor) (*BorrowersTransactor, error) {
	contract, err := bindBorrowers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BorrowersTransactor{contract: contract}, nil
}

// NewBorrowersFilterer creates a new log filterer instance of Borrowers, bound to a specific deployed contract.
func NewBorrowersFilterer(address common.Address, filterer bind.ContractFilterer) (*BorrowersFilterer, error) {
	contract, err := bindBorrowers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BorrowersFilterer{contract: contract}, nil
}

// bindBorrowers binds a generic wrapper to an already deployed contract.
func bindBorrowers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BorrowersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Borrowers *BorrowersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Borrowers.Contract.BorrowersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Borrowers *BorrowersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.Contract.BorrowersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Borrowers *BorrowersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Borrowers.Contract.BorrowersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Borrowers *BorrowersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Borrowers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Borrowers *BorrowersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Borrowers *BorrowersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Borrowers.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Borrowers *BorrowersCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Borrowers *BorrowersSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Borrowers.Contract.DEFAULTADMINROLE(&_Borrowers.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Borrowers *BorrowersCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Borrowers.Contract.DEFAULTADMINROLE(&_Borrowers.CallOpts)
}

// MINTER is a free data retrieval call binding the contract method 0xfe6d8124.
//
// Solidity: function MINTER() view returns(bytes32)
func (_Borrowers *BorrowersCaller) MINTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "MINTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTER is a free data retrieval call binding the contract method 0xfe6d8124.
//
// Solidity: function MINTER() view returns(bytes32)
func (_Borrowers *BorrowersSession) MINTER() ([32]byte, error) {
	return _Borrowers.Contract.MINTER(&_Borrowers.CallOpts)
}

// MINTER is a free data retrieval call binding the contract method 0xfe6d8124.
//
// Solidity: function MINTER() view returns(bytes32)
func (_Borrowers *BorrowersCallerSession) MINTER() ([32]byte, error) {
	return _Borrowers.Contract.MINTER(&_Borrowers.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_Borrowers *BorrowersCaller) OPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_Borrowers *BorrowersSession) OPERATOR() ([32]byte, error) {
	return _Borrowers.Contract.OPERATOR(&_Borrowers.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_Borrowers *BorrowersCallerSession) OPERATOR() ([32]byte, error) {
	return _Borrowers.Contract.OPERATOR(&_Borrowers.CallOpts)
}

// AccrualTimeStamp is a free data retrieval call binding the contract method 0xe79b9667.
//
// Solidity: function accrualTimeStamp() view returns(uint256)
func (_Borrowers *BorrowersCaller) AccrualTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "accrualTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualTimeStamp is a free data retrieval call binding the contract method 0xe79b9667.
//
// Solidity: function accrualTimeStamp() view returns(uint256)
func (_Borrowers *BorrowersSession) AccrualTimeStamp() (*big.Int, error) {
	return _Borrowers.Contract.AccrualTimeStamp(&_Borrowers.CallOpts)
}

// AccrualTimeStamp is a free data retrieval call binding the contract method 0xe79b9667.
//
// Solidity: function accrualTimeStamp() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) AccrualTimeStamp() (*big.Int, error) {
	return _Borrowers.Contract.AccrualTimeStamp(&_Borrowers.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Borrowers *BorrowersCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Borrowers *BorrowersSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Borrowers.Contract.Allowance(&_Borrowers.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Borrowers *BorrowersCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Borrowers.Contract.Allowance(&_Borrowers.CallOpts, owner, spender)
}

// BadDebt is a free data retrieval call binding the contract method 0xbbcac557.
//
// Solidity: function badDebt() view returns(uint256)
func (_Borrowers *BorrowersCaller) BadDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "badDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BadDebt is a free data retrieval call binding the contract method 0xbbcac557.
//
// Solidity: function badDebt() view returns(uint256)
func (_Borrowers *BorrowersSession) BadDebt() (*big.Int, error) {
	return _Borrowers.Contract.BadDebt(&_Borrowers.CallOpts)
}

// BadDebt is a free data retrieval call binding the contract method 0xbbcac557.
//
// Solidity: function badDebt() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) BadDebt() (*big.Int, error) {
	return _Borrowers.Contract.BadDebt(&_Borrowers.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Borrowers *BorrowersCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Borrowers *BorrowersSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Borrowers.Contract.BalanceOf(&_Borrowers.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Borrowers *BorrowersCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Borrowers.Contract.BalanceOf(&_Borrowers.CallOpts, account)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) view returns(uint256)
func (_Borrowers *BorrowersCaller) BalanceOfUnderlying(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "balanceOfUnderlying", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) view returns(uint256)
func (_Borrowers *BorrowersSession) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	return _Borrowers.Contract.BalanceOfUnderlying(&_Borrowers.CallOpts, owner)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) view returns(uint256)
func (_Borrowers *BorrowersCallerSession) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	return _Borrowers.Contract.BalanceOfUnderlying(&_Borrowers.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Borrowers *BorrowersCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Borrowers *BorrowersSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Borrowers.Contract.BorrowBalanceStored(&_Borrowers.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Borrowers *BorrowersCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Borrowers.Contract.BorrowBalanceStored(&_Borrowers.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Borrowers *BorrowersCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Borrowers *BorrowersSession) BorrowIndex() (*big.Int, error) {
	return _Borrowers.Contract.BorrowIndex(&_Borrowers.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) BorrowIndex() (*big.Int, error) {
	return _Borrowers.Contract.BorrowIndex(&_Borrowers.CallOpts)
}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x52609750.
//
// Solidity: function borrowRatePerSecond() view returns(uint256)
func (_Borrowers *BorrowersCaller) BorrowRatePerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "borrowRatePerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x52609750.
//
// Solidity: function borrowRatePerSecond() view returns(uint256)
func (_Borrowers *BorrowersSession) BorrowRatePerSecond() (*big.Int, error) {
	return _Borrowers.Contract.BorrowRatePerSecond(&_Borrowers.CallOpts)
}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x52609750.
//
// Solidity: function borrowRatePerSecond() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) BorrowRatePerSecond() (*big.Int, error) {
	return _Borrowers.Contract.BorrowRatePerSecond(&_Borrowers.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xc6788bdd.
//
// Solidity: function claims(address ) view returns(bool)
func (_Borrowers *BorrowersCaller) Claims(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "claims", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claims is a free data retrieval call binding the contract method 0xc6788bdd.
//
// Solidity: function claims(address ) view returns(bool)
func (_Borrowers *BorrowersSession) Claims(arg0 common.Address) (bool, error) {
	return _Borrowers.Contract.Claims(&_Borrowers.CallOpts, arg0)
}

// Claims is a free data retrieval call binding the contract method 0xc6788bdd.
//
// Solidity: function claims(address ) view returns(bool)
func (_Borrowers *BorrowersCallerSession) Claims(arg0 common.Address) (bool, error) {
	return _Borrowers.Contract.Claims(&_Borrowers.CallOpts, arg0)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_Borrowers *BorrowersCaller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "compliance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_Borrowers *BorrowersSession) Compliance() (common.Address, error) {
	return _Borrowers.Contract.Compliance(&_Borrowers.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_Borrowers *BorrowersCallerSession) Compliance() (common.Address, error) {
	return _Borrowers.Contract.Compliance(&_Borrowers.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Borrowers *BorrowersCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Borrowers *BorrowersSession) Decimals() (uint8, error) {
	return _Borrowers.Contract.Decimals(&_Borrowers.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Borrowers *BorrowersCallerSession) Decimals() (uint8, error) {
	return _Borrowers.Contract.Decimals(&_Borrowers.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Borrowers *BorrowersCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Borrowers *BorrowersSession) ExchangeRateStored() (*big.Int, error) {
	return _Borrowers.Contract.ExchangeRateStored(&_Borrowers.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Borrowers.Contract.ExchangeRateStored(&_Borrowers.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256)
func (_Borrowers *BorrowersCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256)
func (_Borrowers *BorrowersSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Borrowers.Contract.GetAccountSnapshot(&_Borrowers.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256)
func (_Borrowers *BorrowersCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Borrowers.Contract.GetAccountSnapshot(&_Borrowers.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Borrowers *BorrowersCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Borrowers *BorrowersSession) GetCash() (*big.Int, error) {
	return _Borrowers.Contract.GetCash(&_Borrowers.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) GetCash() (*big.Int, error) {
	return _Borrowers.Contract.GetCash(&_Borrowers.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Borrowers *BorrowersCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Borrowers *BorrowersSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Borrowers.Contract.GetRoleAdmin(&_Borrowers.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Borrowers *BorrowersCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Borrowers.Contract.GetRoleAdmin(&_Borrowers.CallOpts, role)
}

// GrantBorrows is a free data retrieval call binding the contract method 0x6bf4940c.
//
// Solidity: function grantBorrows(address , address ) view returns(uint256)
func (_Borrowers *BorrowersCaller) GrantBorrows(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "grantBorrows", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GrantBorrows is a free data retrieval call binding the contract method 0x6bf4940c.
//
// Solidity: function grantBorrows(address , address ) view returns(uint256)
func (_Borrowers *BorrowersSession) GrantBorrows(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Borrowers.Contract.GrantBorrows(&_Borrowers.CallOpts, arg0, arg1)
}

// GrantBorrows is a free data retrieval call binding the contract method 0x6bf4940c.
//
// Solidity: function grantBorrows(address , address ) view returns(uint256)
func (_Borrowers *BorrowersCallerSession) GrantBorrows(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Borrowers.Contract.GrantBorrows(&_Borrowers.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Borrowers *BorrowersCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Borrowers *BorrowersSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Borrowers.Contract.HasRole(&_Borrowers.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Borrowers *BorrowersCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Borrowers.Contract.HasRole(&_Borrowers.CallOpts, role, account)
}

// HyController is a free data retrieval call binding the contract method 0xb6f64810.
//
// Solidity: function hyController() view returns(address)
func (_Borrowers *BorrowersCaller) HyController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "hyController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HyController is a free data retrieval call binding the contract method 0xb6f64810.
//
// Solidity: function hyController() view returns(address)
func (_Borrowers *BorrowersSession) HyController() (common.Address, error) {
	return _Borrowers.Contract.HyController(&_Borrowers.CallOpts)
}

// HyController is a free data retrieval call binding the contract method 0xb6f64810.
//
// Solidity: function hyController() view returns(address)
func (_Borrowers *BorrowersCallerSession) HyController() (common.Address, error) {
	return _Borrowers.Contract.HyController(&_Borrowers.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Borrowers *BorrowersCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Borrowers *BorrowersSession) InterestRateModel() (common.Address, error) {
	return _Borrowers.Contract.InterestRateModel(&_Borrowers.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Borrowers *BorrowersCallerSession) InterestRateModel() (common.Address, error) {
	return _Borrowers.Contract.InterestRateModel(&_Borrowers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Borrowers *BorrowersCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Borrowers *BorrowersSession) Name() (string, error) {
	return _Borrowers.Contract.Name(&_Borrowers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Borrowers *BorrowersCallerSession) Name() (string, error) {
	return _Borrowers.Contract.Name(&_Borrowers.CallOpts)
}

// PassivePool is a free data retrieval call binding the contract method 0x87f6c1e0.
//
// Solidity: function passivePool() view returns(address)
func (_Borrowers *BorrowersCaller) PassivePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "passivePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PassivePool is a free data retrieval call binding the contract method 0x87f6c1e0.
//
// Solidity: function passivePool() view returns(address)
func (_Borrowers *BorrowersSession) PassivePool() (common.Address, error) {
	return _Borrowers.Contract.PassivePool(&_Borrowers.CallOpts)
}

// PassivePool is a free data retrieval call binding the contract method 0x87f6c1e0.
//
// Solidity: function passivePool() view returns(address)
func (_Borrowers *BorrowersCallerSession) PassivePool() (common.Address, error) {
	return _Borrowers.Contract.PassivePool(&_Borrowers.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Borrowers *BorrowersCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Borrowers *BorrowersSession) Paused() (bool, error) {
	return _Borrowers.Contract.Paused(&_Borrowers.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Borrowers *BorrowersCallerSession) Paused() (bool, error) {
	return _Borrowers.Contract.Paused(&_Borrowers.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Borrowers *BorrowersCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Borrowers *BorrowersSession) Rate() (*big.Int, error) {
	return _Borrowers.Contract.Rate(&_Borrowers.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) Rate() (*big.Int, error) {
	return _Borrowers.Contract.Rate(&_Borrowers.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Borrowers *BorrowersCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Borrowers *BorrowersSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Borrowers.Contract.ReserveFactorMantissa(&_Borrowers.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Borrowers.Contract.ReserveFactorMantissa(&_Borrowers.CallOpts)
}

// StopAccruing is a free data retrieval call binding the contract method 0x7c1e6231.
//
// Solidity: function stopAccruing() view returns(bool)
func (_Borrowers *BorrowersCaller) StopAccruing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "stopAccruing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StopAccruing is a free data retrieval call binding the contract method 0x7c1e6231.
//
// Solidity: function stopAccruing() view returns(bool)
func (_Borrowers *BorrowersSession) StopAccruing() (bool, error) {
	return _Borrowers.Contract.StopAccruing(&_Borrowers.CallOpts)
}

// StopAccruing is a free data retrieval call binding the contract method 0x7c1e6231.
//
// Solidity: function stopAccruing() view returns(bool)
func (_Borrowers *BorrowersCallerSession) StopAccruing() (bool, error) {
	return _Borrowers.Contract.StopAccruing(&_Borrowers.CallOpts)
}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0xb1d38974.
//
// Solidity: function supplyRatePerSecond() view returns(uint256)
func (_Borrowers *BorrowersCaller) SupplyRatePerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "supplyRatePerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0xb1d38974.
//
// Solidity: function supplyRatePerSecond() view returns(uint256)
func (_Borrowers *BorrowersSession) SupplyRatePerSecond() (*big.Int, error) {
	return _Borrowers.Contract.SupplyRatePerSecond(&_Borrowers.CallOpts)
}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0xb1d38974.
//
// Solidity: function supplyRatePerSecond() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) SupplyRatePerSecond() (*big.Int, error) {
	return _Borrowers.Contract.SupplyRatePerSecond(&_Borrowers.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Borrowers *BorrowersCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Borrowers *BorrowersSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Borrowers.Contract.SupportsInterface(&_Borrowers.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Borrowers *BorrowersCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Borrowers.Contract.SupportsInterface(&_Borrowers.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Borrowers *BorrowersCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Borrowers *BorrowersSession) Symbol() (string, error) {
	return _Borrowers.Contract.Symbol(&_Borrowers.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Borrowers *BorrowersCallerSession) Symbol() (string, error) {
	return _Borrowers.Contract.Symbol(&_Borrowers.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Borrowers *BorrowersCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Borrowers *BorrowersSession) TotalBorrows() (*big.Int, error) {
	return _Borrowers.Contract.TotalBorrows(&_Borrowers.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) TotalBorrows() (*big.Int, error) {
	return _Borrowers.Contract.TotalBorrows(&_Borrowers.CallOpts)
}

// TotalCanClaim is a free data retrieval call binding the contract method 0x3d7fe5dc.
//
// Solidity: function totalCanClaim() view returns(uint256)
func (_Borrowers *BorrowersCaller) TotalCanClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "totalCanClaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCanClaim is a free data retrieval call binding the contract method 0x3d7fe5dc.
//
// Solidity: function totalCanClaim() view returns(uint256)
func (_Borrowers *BorrowersSession) TotalCanClaim() (*big.Int, error) {
	return _Borrowers.Contract.TotalCanClaim(&_Borrowers.CallOpts)
}

// TotalCanClaim is a free data retrieval call binding the contract method 0x3d7fe5dc.
//
// Solidity: function totalCanClaim() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) TotalCanClaim() (*big.Int, error) {
	return _Borrowers.Contract.TotalCanClaim(&_Borrowers.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Borrowers *BorrowersCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Borrowers *BorrowersSession) TotalReserves() (*big.Int, error) {
	return _Borrowers.Contract.TotalReserves(&_Borrowers.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) TotalReserves() (*big.Int, error) {
	return _Borrowers.Contract.TotalReserves(&_Borrowers.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Borrowers *BorrowersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Borrowers *BorrowersSession) TotalSupply() (*big.Int, error) {
	return _Borrowers.Contract.TotalSupply(&_Borrowers.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) TotalSupply() (*big.Int, error) {
	return _Borrowers.Contract.TotalSupply(&_Borrowers.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Borrowers *BorrowersCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Borrowers *BorrowersSession) UnderlyingToken() (common.Address, error) {
	return _Borrowers.Contract.UnderlyingToken(&_Borrowers.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Borrowers *BorrowersCallerSession) UnderlyingToken() (common.Address, error) {
	return _Borrowers.Contract.UnderlyingToken(&_Borrowers.CallOpts)
}

// UnwindClaimedCollateral is a free data retrieval call binding the contract method 0xb395d5a9.
//
// Solidity: function unwindClaimedCollateral() view returns(uint256)
func (_Borrowers *BorrowersCaller) UnwindClaimedCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Borrowers.contract.Call(opts, &out, "unwindClaimedCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnwindClaimedCollateral is a free data retrieval call binding the contract method 0xb395d5a9.
//
// Solidity: function unwindClaimedCollateral() view returns(uint256)
func (_Borrowers *BorrowersSession) UnwindClaimedCollateral() (*big.Int, error) {
	return _Borrowers.Contract.UnwindClaimedCollateral(&_Borrowers.CallOpts)
}

// UnwindClaimedCollateral is a free data retrieval call binding the contract method 0xb395d5a9.
//
// Solidity: function unwindClaimedCollateral() view returns(uint256)
func (_Borrowers *BorrowersCallerSession) UnwindClaimedCollateral() (*big.Int, error) {
	return _Borrowers.Contract.UnwindClaimedCollateral(&_Borrowers.CallOpts)
}

// HYTokenInit is a paid mutator transaction binding the contract method 0x5376e4ae.
//
// Solidity: function __HYToken_init(address interestRateModel_, uint256 initialExchangeRateMantissa_) returns()
func (_Borrowers *BorrowersTransactor) HYTokenInit(opts *bind.TransactOpts, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "__HYToken_init", interestRateModel_, initialExchangeRateMantissa_)
}

// HYTokenInit is a paid mutator transaction binding the contract method 0x5376e4ae.
//
// Solidity: function __HYToken_init(address interestRateModel_, uint256 initialExchangeRateMantissa_) returns()
func (_Borrowers *BorrowersSession) HYTokenInit(interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.HYTokenInit(&_Borrowers.TransactOpts, interestRateModel_, initialExchangeRateMantissa_)
}

// HYTokenInit is a paid mutator transaction binding the contract method 0x5376e4ae.
//
// Solidity: function __HYToken_init(address interestRateModel_, uint256 initialExchangeRateMantissa_) returns()
func (_Borrowers *BorrowersTransactorSession) HYTokenInit(interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.HYTokenInit(&_Borrowers.TransactOpts, interestRateModel_, initialExchangeRateMantissa_)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns()
func (_Borrowers *BorrowersTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns()
func (_Borrowers *BorrowersSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.AddReserves(&_Borrowers.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns()
func (_Borrowers *BorrowersTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.AddReserves(&_Borrowers.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns()
func (_Borrowers *BorrowersTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns()
func (_Borrowers *BorrowersSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.ReduceReserves(&_Borrowers.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns()
func (_Borrowers *BorrowersTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.ReduceReserves(&_Borrowers.TransactOpts, reduceAmount)
}

// SetHYController is a paid mutator transaction binding the contract method 0x104fe6e7.
//
// Solidity: function _setHYController(address newController) returns()
func (_Borrowers *BorrowersTransactor) SetHYController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "_setHYController", newController)
}

// SetHYController is a paid mutator transaction binding the contract method 0x104fe6e7.
//
// Solidity: function _setHYController(address newController) returns()
func (_Borrowers *BorrowersSession) SetHYController(newController common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetHYController(&_Borrowers.TransactOpts, newController)
}

// SetHYController is a paid mutator transaction binding the contract method 0x104fe6e7.
//
// Solidity: function _setHYController(address newController) returns()
func (_Borrowers *BorrowersTransactorSession) SetHYController(newController common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetHYController(&_Borrowers.TransactOpts, newController)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns()
func (_Borrowers *BorrowersTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns()
func (_Borrowers *BorrowersSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetInterestRateModel(&_Borrowers.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns()
func (_Borrowers *BorrowersTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetInterestRateModel(&_Borrowers.TransactOpts, newInterestRateModel)
}

// SetPassivePool is a paid mutator transaction binding the contract method 0x0c66c0a4.
//
// Solidity: function _setPassivePool(address newPassivePool) returns()
func (_Borrowers *BorrowersTransactor) SetPassivePool(opts *bind.TransactOpts, newPassivePool common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "_setPassivePool", newPassivePool)
}

// SetPassivePool is a paid mutator transaction binding the contract method 0x0c66c0a4.
//
// Solidity: function _setPassivePool(address newPassivePool) returns()
func (_Borrowers *BorrowersSession) SetPassivePool(newPassivePool common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetPassivePool(&_Borrowers.TransactOpts, newPassivePool)
}

// SetPassivePool is a paid mutator transaction binding the contract method 0x0c66c0a4.
//
// Solidity: function _setPassivePool(address newPassivePool) returns()
func (_Borrowers *BorrowersTransactorSession) SetPassivePool(newPassivePool common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetPassivePool(&_Borrowers.TransactOpts, newPassivePool)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns()
func (_Borrowers *BorrowersTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns()
func (_Borrowers *BorrowersSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.SetReserveFactor(&_Borrowers.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns()
func (_Borrowers *BorrowersTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.SetReserveFactor(&_Borrowers.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_Borrowers *BorrowersTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_Borrowers *BorrowersSession) AccrueInterest() (*types.Transaction, error) {
	return _Borrowers.Contract.AccrueInterest(&_Borrowers.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_Borrowers *BorrowersTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Borrowers.Contract.AccrueInterest(&_Borrowers.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Borrowers *BorrowersTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Borrowers *BorrowersSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Approve(&_Borrowers.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Borrowers *BorrowersTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Approve(&_Borrowers.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns()
func (_Borrowers *BorrowersTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns()
func (_Borrowers *BorrowersSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Borrow(&_Borrowers.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns()
func (_Borrowers *BorrowersTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Borrow(&_Borrowers.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Borrowers *BorrowersTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Borrowers *BorrowersSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.BorrowBalanceCurrent(&_Borrowers.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Borrowers *BorrowersTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.BorrowBalanceCurrent(&_Borrowers.TransactOpts, account)
}

// BorrowFrom is a paid mutator transaction binding the contract method 0x4196aaaf.
//
// Solidity: function borrowFrom(address from, uint256 borrowAmount) returns()
func (_Borrowers *BorrowersTransactor) BorrowFrom(opts *bind.TransactOpts, from common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "borrowFrom", from, borrowAmount)
}

// BorrowFrom is a paid mutator transaction binding the contract method 0x4196aaaf.
//
// Solidity: function borrowFrom(address from, uint256 borrowAmount) returns()
func (_Borrowers *BorrowersSession) BorrowFrom(from common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.BorrowFrom(&_Borrowers.TransactOpts, from, borrowAmount)
}

// BorrowFrom is a paid mutator transaction binding the contract method 0x4196aaaf.
//
// Solidity: function borrowFrom(address from, uint256 borrowAmount) returns()
func (_Borrowers *BorrowersTransactorSession) BorrowFrom(from common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.BorrowFrom(&_Borrowers.TransactOpts, from, borrowAmount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Borrowers *BorrowersTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Borrowers *BorrowersSession) Claim() (*types.Transaction, error) {
	return _Borrowers.Contract.Claim(&_Borrowers.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Borrowers *BorrowersTransactorSession) Claim() (*types.Transaction, error) {
	return _Borrowers.Contract.Claim(&_Borrowers.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Borrowers *BorrowersTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Borrowers *BorrowersSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.DecreaseAllowance(&_Borrowers.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Borrowers *BorrowersTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.DecreaseAllowance(&_Borrowers.TransactOpts, spender, subtractedValue)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Borrowers *BorrowersTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Borrowers *BorrowersSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Borrowers.Contract.ExchangeRateCurrent(&_Borrowers.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Borrowers *BorrowersTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Borrowers.Contract.ExchangeRateCurrent(&_Borrowers.TransactOpts)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address from_, address to_, uint256 amount_) returns(bool)
func (_Borrowers *BorrowersTransactor) ForcedTransfer(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "forcedTransfer", from_, to_, amount_)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address from_, address to_, uint256 amount_) returns(bool)
func (_Borrowers *BorrowersSession) ForcedTransfer(from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.ForcedTransfer(&_Borrowers.TransactOpts, from_, to_, amount_)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address from_, address to_, uint256 amount_) returns(bool)
func (_Borrowers *BorrowersTransactorSession) ForcedTransfer(from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.ForcedTransfer(&_Borrowers.TransactOpts, from_, to_, amount_)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Borrowers *BorrowersTransactor) FreezePartialTokens(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "freezePartialTokens", _userAddress, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Borrowers *BorrowersSession) FreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.FreezePartialTokens(&_Borrowers.TransactOpts, _userAddress, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Borrowers *BorrowersTransactorSession) FreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.FreezePartialTokens(&_Borrowers.TransactOpts, _userAddress, _amount)
}

// GrantBorrow is a paid mutator transaction binding the contract method 0x6463d80e.
//
// Solidity: function grantBorrow(address to, uint256 amount) returns()
func (_Borrowers *BorrowersTransactor) GrantBorrow(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "grantBorrow", to, amount)
}

// GrantBorrow is a paid mutator transaction binding the contract method 0x6463d80e.
//
// Solidity: function grantBorrow(address to, uint256 amount) returns()
func (_Borrowers *BorrowersSession) GrantBorrow(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.GrantBorrow(&_Borrowers.TransactOpts, to, amount)
}

// GrantBorrow is a paid mutator transaction binding the contract method 0x6463d80e.
//
// Solidity: function grantBorrow(address to, uint256 amount) returns()
func (_Borrowers *BorrowersTransactorSession) GrantBorrow(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.GrantBorrow(&_Borrowers.TransactOpts, to, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.GrantRole(&_Borrowers.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.GrantRole(&_Borrowers.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Borrowers *BorrowersTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Borrowers *BorrowersSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.IncreaseAllowance(&_Borrowers.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Borrowers *BorrowersTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.IncreaseAllowance(&_Borrowers.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x1d4a8855.
//
// Solidity: function initialize(address interestRateModel_, uint256 initialExchangeRateMantissa_, address underlyingToken_, string name_, string symbol_, address hyController_, address passivePool_, address _rewards, address poolOwner_) returns()
func (_Borrowers *BorrowersTransactor) Initialize(opts *bind.TransactOpts, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, underlyingToken_ common.Address, name_ string, symbol_ string, hyController_ common.Address, passivePool_ common.Address, _rewards common.Address, poolOwner_ common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "initialize", interestRateModel_, initialExchangeRateMantissa_, underlyingToken_, name_, symbol_, hyController_, passivePool_, _rewards, poolOwner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1d4a8855.
//
// Solidity: function initialize(address interestRateModel_, uint256 initialExchangeRateMantissa_, address underlyingToken_, string name_, string symbol_, address hyController_, address passivePool_, address _rewards, address poolOwner_) returns()
func (_Borrowers *BorrowersSession) Initialize(interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, underlyingToken_ common.Address, name_ string, symbol_ string, hyController_ common.Address, passivePool_ common.Address, _rewards common.Address, poolOwner_ common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.Initialize(&_Borrowers.TransactOpts, interestRateModel_, initialExchangeRateMantissa_, underlyingToken_, name_, symbol_, hyController_, passivePool_, _rewards, poolOwner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1d4a8855.
//
// Solidity: function initialize(address interestRateModel_, uint256 initialExchangeRateMantissa_, address underlyingToken_, string name_, string symbol_, address hyController_, address passivePool_, address _rewards, address poolOwner_) returns()
func (_Borrowers *BorrowersTransactorSession) Initialize(interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, underlyingToken_ common.Address, name_ string, symbol_ string, hyController_ common.Address, passivePool_ common.Address, _rewards common.Address, poolOwner_ common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.Initialize(&_Borrowers.TransactOpts, interestRateModel_, initialExchangeRateMantissa_, underlyingToken_, name_, symbol_, hyController_, passivePool_, _rewards, poolOwner_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Borrowers *BorrowersTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Borrowers *BorrowersSession) Pause() (*types.Transaction, error) {
	return _Borrowers.Contract.Pause(&_Borrowers.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Borrowers *BorrowersTransactorSession) Pause() (*types.Transaction, error) {
	return _Borrowers.Contract.Pause(&_Borrowers.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns()
func (_Borrowers *BorrowersTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns()
func (_Borrowers *BorrowersSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Redeem(&_Borrowers.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns()
func (_Borrowers *BorrowersTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Redeem(&_Borrowers.TransactOpts, redeemTokens)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.RenounceRole(&_Borrowers.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.RenounceRole(&_Borrowers.TransactOpts, role, account)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns()
func (_Borrowers *BorrowersTransactor) Repay(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "repay", repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns()
func (_Borrowers *BorrowersSession) Repay(repayAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Repay(&_Borrowers.TransactOpts, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns()
func (_Borrowers *BorrowersTransactorSession) Repay(repayAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Repay(&_Borrowers.TransactOpts, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address borrower, uint256 repayAmount) returns()
func (_Borrowers *BorrowersTransactor) RepayOnBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "repayOnBehalf", borrower, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address borrower, uint256 repayAmount) returns()
func (_Borrowers *BorrowersSession) RepayOnBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.RepayOnBehalf(&_Borrowers.TransactOpts, borrower, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address borrower, uint256 repayAmount) returns()
func (_Borrowers *BorrowersTransactorSession) RepayOnBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.RepayOnBehalf(&_Borrowers.TransactOpts, borrower, repayAmount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.RevokeRole(&_Borrowers.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Borrowers *BorrowersTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.RevokeRole(&_Borrowers.TransactOpts, role, account)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_Borrowers *BorrowersTransactor) SetAddressFrozen(opts *bind.TransactOpts, _userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "setAddressFrozen", _userAddress, _freeze)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_Borrowers *BorrowersSession) SetAddressFrozen(_userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _Borrowers.Contract.SetAddressFrozen(&_Borrowers.TransactOpts, _userAddress, _freeze)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_Borrowers *BorrowersTransactorSession) SetAddressFrozen(_userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _Borrowers.Contract.SetAddressFrozen(&_Borrowers.TransactOpts, _userAddress, _freeze)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address compliance_) returns()
func (_Borrowers *BorrowersTransactor) SetCompliance(opts *bind.TransactOpts, compliance_ common.Address) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "setCompliance", compliance_)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address compliance_) returns()
func (_Borrowers *BorrowersSession) SetCompliance(compliance_ common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetCompliance(&_Borrowers.TransactOpts, compliance_)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address compliance_) returns()
func (_Borrowers *BorrowersTransactorSession) SetCompliance(compliance_ common.Address) (*types.Transaction, error) {
	return _Borrowers.Contract.SetCompliance(&_Borrowers.TransactOpts, compliance_)
}

// StopAccruingInterest is a paid mutator transaction binding the contract method 0xdd507fa9.
//
// Solidity: function stopAccruingInterest() returns()
func (_Borrowers *BorrowersTransactor) StopAccruingInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "stopAccruingInterest")
}

// StopAccruingInterest is a paid mutator transaction binding the contract method 0xdd507fa9.
//
// Solidity: function stopAccruingInterest() returns()
func (_Borrowers *BorrowersSession) StopAccruingInterest() (*types.Transaction, error) {
	return _Borrowers.Contract.StopAccruingInterest(&_Borrowers.TransactOpts)
}

// StopAccruingInterest is a paid mutator transaction binding the contract method 0xdd507fa9.
//
// Solidity: function stopAccruingInterest() returns()
func (_Borrowers *BorrowersTransactorSession) StopAccruingInterest() (*types.Transaction, error) {
	return _Borrowers.Contract.StopAccruingInterest(&_Borrowers.TransactOpts)
}

// Supply is a paid mutator transaction binding the contract method 0x35403023.
//
// Solidity: function supply(uint256 mintAmount) returns()
func (_Borrowers *BorrowersTransactor) Supply(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "supply", mintAmount)
}

// Supply is a paid mutator transaction binding the contract method 0x35403023.
//
// Solidity: function supply(uint256 mintAmount) returns()
func (_Borrowers *BorrowersSession) Supply(mintAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Supply(&_Borrowers.TransactOpts, mintAmount)
}

// Supply is a paid mutator transaction binding the contract method 0x35403023.
//
// Solidity: function supply(uint256 mintAmount) returns()
func (_Borrowers *BorrowersTransactorSession) Supply(mintAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Supply(&_Borrowers.TransactOpts, mintAmount)
}

// SupplyFor is a paid mutator transaction binding the contract method 0x21d18e4f.
//
// Solidity: function supplyFor(address minter, uint256 mintAmount) returns()
func (_Borrowers *BorrowersTransactor) SupplyFor(opts *bind.TransactOpts, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "supplyFor", minter, mintAmount)
}

// SupplyFor is a paid mutator transaction binding the contract method 0x21d18e4f.
//
// Solidity: function supplyFor(address minter, uint256 mintAmount) returns()
func (_Borrowers *BorrowersSession) SupplyFor(minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.SupplyFor(&_Borrowers.TransactOpts, minter, mintAmount)
}

// SupplyFor is a paid mutator transaction binding the contract method 0x21d18e4f.
//
// Solidity: function supplyFor(address minter, uint256 mintAmount) returns()
func (_Borrowers *BorrowersTransactorSession) SupplyFor(minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.SupplyFor(&_Borrowers.TransactOpts, minter, mintAmount)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Borrowers *BorrowersTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Borrowers *BorrowersSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Borrowers.Contract.TotalBorrowsCurrent(&_Borrowers.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Borrowers *BorrowersTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Borrowers.Contract.TotalBorrowsCurrent(&_Borrowers.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Borrowers *BorrowersTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Borrowers *BorrowersSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Transfer(&_Borrowers.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Borrowers *BorrowersTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.Transfer(&_Borrowers.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Borrowers *BorrowersTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Borrowers *BorrowersSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.TransferFrom(&_Borrowers.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Borrowers *BorrowersTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.TransferFrom(&_Borrowers.TransactOpts, from, to, amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Borrowers *BorrowersTransactor) UnfreezePartialTokens(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "unfreezePartialTokens", _userAddress, _amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Borrowers *BorrowersSession) UnfreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.UnfreezePartialTokens(&_Borrowers.TransactOpts, _userAddress, _amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Borrowers *BorrowersTransactorSession) UnfreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.UnfreezePartialTokens(&_Borrowers.TransactOpts, _userAddress, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Borrowers *BorrowersTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Borrowers *BorrowersSession) Unpause() (*types.Transaction, error) {
	return _Borrowers.Contract.Unpause(&_Borrowers.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Borrowers *BorrowersTransactorSession) Unpause() (*types.Transaction, error) {
	return _Borrowers.Contract.Unpause(&_Borrowers.TransactOpts)
}

// UnwindCollateral is a paid mutator transaction binding the contract method 0xda0dfe18.
//
// Solidity: function unwindCollateral(uint256 unwindAmount_, uint256 badDebt_) returns()
func (_Borrowers *BorrowersTransactor) UnwindCollateral(opts *bind.TransactOpts, unwindAmount_ *big.Int, badDebt_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.contract.Transact(opts, "unwindCollateral", unwindAmount_, badDebt_)
}

// UnwindCollateral is a paid mutator transaction binding the contract method 0xda0dfe18.
//
// Solidity: function unwindCollateral(uint256 unwindAmount_, uint256 badDebt_) returns()
func (_Borrowers *BorrowersSession) UnwindCollateral(unwindAmount_ *big.Int, badDebt_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.UnwindCollateral(&_Borrowers.TransactOpts, unwindAmount_, badDebt_)
}

// UnwindCollateral is a paid mutator transaction binding the contract method 0xda0dfe18.
//
// Solidity: function unwindCollateral(uint256 unwindAmount_, uint256 badDebt_) returns()
func (_Borrowers *BorrowersTransactorSession) UnwindCollateral(unwindAmount_ *big.Int, badDebt_ *big.Int) (*types.Transaction, error) {
	return _Borrowers.Contract.UnwindCollateral(&_Borrowers.TransactOpts, unwindAmount_, badDebt_)
}

// BorrowersAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Borrowers contract.
type BorrowersAccrueInterestIterator struct {
	Event *BorrowersAccrueInterest // Event containing the contract specifics and raw log

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
func (it *BorrowersAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersAccrueInterest)
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
		it.Event = new(BorrowersAccrueInterest)
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
func (it *BorrowersAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersAccrueInterest represents a AccrueInterest event raised by the Borrowers contract.
type BorrowersAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Borrowers *BorrowersFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*BorrowersAccrueInterestIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &BorrowersAccrueInterestIterator{contract: _Borrowers.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Borrowers *BorrowersFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *BorrowersAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersAccrueInterest)
				if err := _Borrowers.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Borrowers *BorrowersFilterer) ParseAccrueInterest(log types.Log) (*BorrowersAccrueInterest, error) {
	event := new(BorrowersAccrueInterest)
	if err := _Borrowers.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersAddressFrozenIterator is returned from FilterAddressFrozen and is used to iterate over the raw logs and unpacked data for AddressFrozen events raised by the Borrowers contract.
type BorrowersAddressFrozenIterator struct {
	Event *BorrowersAddressFrozen // Event containing the contract specifics and raw log

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
func (it *BorrowersAddressFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersAddressFrozen)
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
		it.Event = new(BorrowersAddressFrozen)
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
func (it *BorrowersAddressFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersAddressFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersAddressFrozen represents a AddressFrozen event raised by the Borrowers contract.
type BorrowersAddressFrozen struct {
	Arg0 common.Address
	Arg1 bool
	Arg2 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressFrozen is a free log retrieval operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address arg0, bool arg1, address arg2)
func (_Borrowers *BorrowersFilterer) FilterAddressFrozen(opts *bind.FilterOpts) (*BorrowersAddressFrozenIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "AddressFrozen")
	if err != nil {
		return nil, err
	}
	return &BorrowersAddressFrozenIterator{contract: _Borrowers.contract, event: "AddressFrozen", logs: logs, sub: sub}, nil
}

// WatchAddressFrozen is a free log subscription operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address arg0, bool arg1, address arg2)
func (_Borrowers *BorrowersFilterer) WatchAddressFrozen(opts *bind.WatchOpts, sink chan<- *BorrowersAddressFrozen) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "AddressFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersAddressFrozen)
				if err := _Borrowers.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
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

// ParseAddressFrozen is a log parse operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address arg0, bool arg1, address arg2)
func (_Borrowers *BorrowersFilterer) ParseAddressFrozen(log types.Log) (*BorrowersAddressFrozen, error) {
	event := new(BorrowersAddressFrozen)
	if err := _Borrowers.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Borrowers contract.
type BorrowersApprovalIterator struct {
	Event *BorrowersApproval // Event containing the contract specifics and raw log

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
func (it *BorrowersApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersApproval)
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
		it.Event = new(BorrowersApproval)
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
func (it *BorrowersApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersApproval represents a Approval event raised by the Borrowers contract.
type BorrowersApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Borrowers *BorrowersFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BorrowersApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BorrowersApprovalIterator{contract: _Borrowers.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Borrowers *BorrowersFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BorrowersApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersApproval)
				if err := _Borrowers.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Borrowers *BorrowersFilterer) ParseApproval(log types.Log) (*BorrowersApproval, error) {
	event := new(BorrowersApproval)
	if err := _Borrowers.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Borrowers contract.
type BorrowersBorrowIterator struct {
	Event *BorrowersBorrow // Event containing the contract specifics and raw log

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
func (it *BorrowersBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersBorrow)
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
		it.Event = new(BorrowersBorrow)
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
func (it *BorrowersBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersBorrow represents a Borrow event raised by the Borrowers contract.
type BorrowersBorrow struct {
	Borrower       common.Address
	To             common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x10a0132d3bf8c82a7fb93a86160f3074ca5c3e5706fa2bcdf0e2b5fd495af09b.
//
// Solidity: event Borrow(address borrower, address to, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Borrowers *BorrowersFilterer) FilterBorrow(opts *bind.FilterOpts) (*BorrowersBorrowIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &BorrowersBorrowIterator{contract: _Borrowers.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x10a0132d3bf8c82a7fb93a86160f3074ca5c3e5706fa2bcdf0e2b5fd495af09b.
//
// Solidity: event Borrow(address borrower, address to, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Borrowers *BorrowersFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *BorrowersBorrow) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersBorrow)
				if err := _Borrowers.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x10a0132d3bf8c82a7fb93a86160f3074ca5c3e5706fa2bcdf0e2b5fd495af09b.
//
// Solidity: event Borrow(address borrower, address to, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Borrowers *BorrowersFilterer) ParseBorrow(log types.Log) (*BorrowersBorrow, error) {
	event := new(BorrowersBorrow)
	if err := _Borrowers.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Borrowers contract.
type BorrowersClaimIterator struct {
	Event *BorrowersClaim // Event containing the contract specifics and raw log

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
func (it *BorrowersClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersClaim)
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
		it.Event = new(BorrowersClaim)
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
func (it *BorrowersClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersClaim represents a Claim event raised by the Borrowers contract.
type BorrowersClaim struct {
	Claimer     common.Address
	ClaimAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address claimer, uint256 claimAmount)
func (_Borrowers *BorrowersFilterer) FilterClaim(opts *bind.FilterOpts) (*BorrowersClaimIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &BorrowersClaimIterator{contract: _Borrowers.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address claimer, uint256 claimAmount)
func (_Borrowers *BorrowersFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *BorrowersClaim) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersClaim)
				if err := _Borrowers.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address claimer, uint256 claimAmount)
func (_Borrowers *BorrowersFilterer) ParseClaim(log types.Log) (*BorrowersClaim, error) {
	event := new(BorrowersClaim)
	if err := _Borrowers.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersGrantBorrowIterator is returned from FilterGrantBorrow and is used to iterate over the raw logs and unpacked data for GrantBorrow events raised by the Borrowers contract.
type BorrowersGrantBorrowIterator struct {
	Event *BorrowersGrantBorrow // Event containing the contract specifics and raw log

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
func (it *BorrowersGrantBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersGrantBorrow)
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
		it.Event = new(BorrowersGrantBorrow)
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
func (it *BorrowersGrantBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersGrantBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersGrantBorrow represents a GrantBorrow event raised by the Borrowers contract.
type BorrowersGrantBorrow struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGrantBorrow is a free log retrieval operation binding the contract event 0x697fc025640e5064ceb4eb8f9cfa1e8cfe5b14d9132fc46199b4a6bb3e422492.
//
// Solidity: event GrantBorrow(address from, address to, uint256 amount)
func (_Borrowers *BorrowersFilterer) FilterGrantBorrow(opts *bind.FilterOpts) (*BorrowersGrantBorrowIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "GrantBorrow")
	if err != nil {
		return nil, err
	}
	return &BorrowersGrantBorrowIterator{contract: _Borrowers.contract, event: "GrantBorrow", logs: logs, sub: sub}, nil
}

// WatchGrantBorrow is a free log subscription operation binding the contract event 0x697fc025640e5064ceb4eb8f9cfa1e8cfe5b14d9132fc46199b4a6bb3e422492.
//
// Solidity: event GrantBorrow(address from, address to, uint256 amount)
func (_Borrowers *BorrowersFilterer) WatchGrantBorrow(opts *bind.WatchOpts, sink chan<- *BorrowersGrantBorrow) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "GrantBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersGrantBorrow)
				if err := _Borrowers.contract.UnpackLog(event, "GrantBorrow", log); err != nil {
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

// ParseGrantBorrow is a log parse operation binding the contract event 0x697fc025640e5064ceb4eb8f9cfa1e8cfe5b14d9132fc46199b4a6bb3e422492.
//
// Solidity: event GrantBorrow(address from, address to, uint256 amount)
func (_Borrowers *BorrowersFilterer) ParseGrantBorrow(log types.Log) (*BorrowersGrantBorrow, error) {
	event := new(BorrowersGrantBorrow)
	if err := _Borrowers.contract.UnpackLog(event, "GrantBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Borrowers contract.
type BorrowersInitializedIterator struct {
	Event *BorrowersInitialized // Event containing the contract specifics and raw log

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
func (it *BorrowersInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersInitialized)
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
		it.Event = new(BorrowersInitialized)
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
func (it *BorrowersInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersInitialized represents a Initialized event raised by the Borrowers contract.
type BorrowersInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Borrowers *BorrowersFilterer) FilterInitialized(opts *bind.FilterOpts) (*BorrowersInitializedIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BorrowersInitializedIterator{contract: _Borrowers.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Borrowers *BorrowersFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BorrowersInitialized) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersInitialized)
				if err := _Borrowers.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Borrowers *BorrowersFilterer) ParseInitialized(log types.Log) (*BorrowersInitialized, error) {
	event := new(BorrowersInitialized)
	if err := _Borrowers.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersNewComplianceIterator is returned from FilterNewCompliance and is used to iterate over the raw logs and unpacked data for NewCompliance events raised by the Borrowers contract.
type BorrowersNewComplianceIterator struct {
	Event *BorrowersNewCompliance // Event containing the contract specifics and raw log

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
func (it *BorrowersNewComplianceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersNewCompliance)
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
		it.Event = new(BorrowersNewCompliance)
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
func (it *BorrowersNewComplianceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersNewComplianceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersNewCompliance represents a NewCompliance event raised by the Borrowers contract.
type BorrowersNewCompliance struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewCompliance is a free log retrieval operation binding the contract event 0xa61d2921544fe9dcd3c393e2d1836b67591de613770c9df0283f64de58075784.
//
// Solidity: event NewCompliance(address arg0)
func (_Borrowers *BorrowersFilterer) FilterNewCompliance(opts *bind.FilterOpts) (*BorrowersNewComplianceIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "NewCompliance")
	if err != nil {
		return nil, err
	}
	return &BorrowersNewComplianceIterator{contract: _Borrowers.contract, event: "NewCompliance", logs: logs, sub: sub}, nil
}

// WatchNewCompliance is a free log subscription operation binding the contract event 0xa61d2921544fe9dcd3c393e2d1836b67591de613770c9df0283f64de58075784.
//
// Solidity: event NewCompliance(address arg0)
func (_Borrowers *BorrowersFilterer) WatchNewCompliance(opts *bind.WatchOpts, sink chan<- *BorrowersNewCompliance) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "NewCompliance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersNewCompliance)
				if err := _Borrowers.contract.UnpackLog(event, "NewCompliance", log); err != nil {
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

// ParseNewCompliance is a log parse operation binding the contract event 0xa61d2921544fe9dcd3c393e2d1836b67591de613770c9df0283f64de58075784.
//
// Solidity: event NewCompliance(address arg0)
func (_Borrowers *BorrowersFilterer) ParseNewCompliance(log types.Log) (*BorrowersNewCompliance, error) {
	event := new(BorrowersNewCompliance)
	if err := _Borrowers.contract.UnpackLog(event, "NewCompliance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersNewHYControllerIterator is returned from FilterNewHYController and is used to iterate over the raw logs and unpacked data for NewHYController events raised by the Borrowers contract.
type BorrowersNewHYControllerIterator struct {
	Event *BorrowersNewHYController // Event containing the contract specifics and raw log

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
func (it *BorrowersNewHYControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersNewHYController)
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
		it.Event = new(BorrowersNewHYController)
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
func (it *BorrowersNewHYControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersNewHYControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersNewHYController represents a NewHYController event raised by the Borrowers contract.
type BorrowersNewHYController struct {
	OldHYController common.Address
	NewHYController common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewHYController is a free log retrieval operation binding the contract event 0x7e4293f869952b2a81c7cc0500d57d22639e1a521328b360f917ba75628630d8.
//
// Solidity: event NewHYController(address oldHYController, address newHYController)
func (_Borrowers *BorrowersFilterer) FilterNewHYController(opts *bind.FilterOpts) (*BorrowersNewHYControllerIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "NewHYController")
	if err != nil {
		return nil, err
	}
	return &BorrowersNewHYControllerIterator{contract: _Borrowers.contract, event: "NewHYController", logs: logs, sub: sub}, nil
}

// WatchNewHYController is a free log subscription operation binding the contract event 0x7e4293f869952b2a81c7cc0500d57d22639e1a521328b360f917ba75628630d8.
//
// Solidity: event NewHYController(address oldHYController, address newHYController)
func (_Borrowers *BorrowersFilterer) WatchNewHYController(opts *bind.WatchOpts, sink chan<- *BorrowersNewHYController) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "NewHYController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersNewHYController)
				if err := _Borrowers.contract.UnpackLog(event, "NewHYController", log); err != nil {
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

// ParseNewHYController is a log parse operation binding the contract event 0x7e4293f869952b2a81c7cc0500d57d22639e1a521328b360f917ba75628630d8.
//
// Solidity: event NewHYController(address oldHYController, address newHYController)
func (_Borrowers *BorrowersFilterer) ParseNewHYController(log types.Log) (*BorrowersNewHYController, error) {
	event := new(BorrowersNewHYController)
	if err := _Borrowers.contract.UnpackLog(event, "NewHYController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Borrowers contract.
type BorrowersNewMarketInterestRateModelIterator struct {
	Event *BorrowersNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *BorrowersNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersNewMarketInterestRateModel)
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
		it.Event = new(BorrowersNewMarketInterestRateModel)
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
func (it *BorrowersNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Borrowers contract.
type BorrowersNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Borrowers *BorrowersFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*BorrowersNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &BorrowersNewMarketInterestRateModelIterator{contract: _Borrowers.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Borrowers *BorrowersFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *BorrowersNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersNewMarketInterestRateModel)
				if err := _Borrowers.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Borrowers *BorrowersFilterer) ParseNewMarketInterestRateModel(log types.Log) (*BorrowersNewMarketInterestRateModel, error) {
	event := new(BorrowersNewMarketInterestRateModel)
	if err := _Borrowers.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersNewPassivePoolIterator is returned from FilterNewPassivePool and is used to iterate over the raw logs and unpacked data for NewPassivePool events raised by the Borrowers contract.
type BorrowersNewPassivePoolIterator struct {
	Event *BorrowersNewPassivePool // Event containing the contract specifics and raw log

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
func (it *BorrowersNewPassivePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersNewPassivePool)
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
		it.Event = new(BorrowersNewPassivePool)
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
func (it *BorrowersNewPassivePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersNewPassivePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersNewPassivePool represents a NewPassivePool event raised by the Borrowers contract.
type BorrowersNewPassivePool struct {
	OldPassivePool common.Address
	NewPassivePool common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPassivePool is a free log retrieval operation binding the contract event 0x65f6bed0ea739216c154efefd0bfa796108ffb9682935199fd2280f54280df3f.
//
// Solidity: event NewPassivePool(address oldPassivePool, address newPassivePool)
func (_Borrowers *BorrowersFilterer) FilterNewPassivePool(opts *bind.FilterOpts) (*BorrowersNewPassivePoolIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "NewPassivePool")
	if err != nil {
		return nil, err
	}
	return &BorrowersNewPassivePoolIterator{contract: _Borrowers.contract, event: "NewPassivePool", logs: logs, sub: sub}, nil
}

// WatchNewPassivePool is a free log subscription operation binding the contract event 0x65f6bed0ea739216c154efefd0bfa796108ffb9682935199fd2280f54280df3f.
//
// Solidity: event NewPassivePool(address oldPassivePool, address newPassivePool)
func (_Borrowers *BorrowersFilterer) WatchNewPassivePool(opts *bind.WatchOpts, sink chan<- *BorrowersNewPassivePool) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "NewPassivePool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersNewPassivePool)
				if err := _Borrowers.contract.UnpackLog(event, "NewPassivePool", log); err != nil {
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

// ParseNewPassivePool is a log parse operation binding the contract event 0x65f6bed0ea739216c154efefd0bfa796108ffb9682935199fd2280f54280df3f.
//
// Solidity: event NewPassivePool(address oldPassivePool, address newPassivePool)
func (_Borrowers *BorrowersFilterer) ParseNewPassivePool(log types.Log) (*BorrowersNewPassivePool, error) {
	event := new(BorrowersNewPassivePool)
	if err := _Borrowers.contract.UnpackLog(event, "NewPassivePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Borrowers contract.
type BorrowersNewReserveFactorIterator struct {
	Event *BorrowersNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *BorrowersNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersNewReserveFactor)
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
		it.Event = new(BorrowersNewReserveFactor)
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
func (it *BorrowersNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersNewReserveFactor represents a NewReserveFactor event raised by the Borrowers contract.
type BorrowersNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Borrowers *BorrowersFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*BorrowersNewReserveFactorIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &BorrowersNewReserveFactorIterator{contract: _Borrowers.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Borrowers *BorrowersFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *BorrowersNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersNewReserveFactor)
				if err := _Borrowers.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Borrowers *BorrowersFilterer) ParseNewReserveFactor(log types.Log) (*BorrowersNewReserveFactor, error) {
	event := new(BorrowersNewReserveFactor)
	if err := _Borrowers.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Borrowers contract.
type BorrowersPausedIterator struct {
	Event *BorrowersPaused // Event containing the contract specifics and raw log

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
func (it *BorrowersPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersPaused)
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
		it.Event = new(BorrowersPaused)
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
func (it *BorrowersPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersPaused represents a Paused event raised by the Borrowers contract.
type BorrowersPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Borrowers *BorrowersFilterer) FilterPaused(opts *bind.FilterOpts) (*BorrowersPausedIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BorrowersPausedIterator{contract: _Borrowers.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Borrowers *BorrowersFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BorrowersPaused) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersPaused)
				if err := _Borrowers.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Borrowers *BorrowersFilterer) ParsePaused(log types.Log) (*BorrowersPaused, error) {
	event := new(BorrowersPaused)
	if err := _Borrowers.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Borrowers contract.
type BorrowersRedeemIterator struct {
	Event *BorrowersRedeem // Event containing the contract specifics and raw log

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
func (it *BorrowersRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersRedeem)
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
		it.Event = new(BorrowersRedeem)
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
func (it *BorrowersRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersRedeem represents a Redeem event raised by the Borrowers contract.
type BorrowersRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Borrowers *BorrowersFilterer) FilterRedeem(opts *bind.FilterOpts) (*BorrowersRedeemIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &BorrowersRedeemIterator{contract: _Borrowers.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Borrowers *BorrowersFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *BorrowersRedeem) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersRedeem)
				if err := _Borrowers.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Borrowers *BorrowersFilterer) ParseRedeem(log types.Log) (*BorrowersRedeem, error) {
	event := new(BorrowersRedeem)
	if err := _Borrowers.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Borrowers contract.
type BorrowersRepayBorrowIterator struct {
	Event *BorrowersRepayBorrow // Event containing the contract specifics and raw log

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
func (it *BorrowersRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersRepayBorrow)
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
		it.Event = new(BorrowersRepayBorrow)
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
func (it *BorrowersRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersRepayBorrow represents a RepayBorrow event raised by the Borrowers contract.
type BorrowersRepayBorrow struct {
	Payer          common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Borrower       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0xf2e05b26b4020b61a94725f7a9bd18e24b72eec2c9fa3592040e872f65dd172a.
//
// Solidity: event RepayBorrow(address payer, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows, address borrower)
func (_Borrowers *BorrowersFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*BorrowersRepayBorrowIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &BorrowersRepayBorrowIterator{contract: _Borrowers.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0xf2e05b26b4020b61a94725f7a9bd18e24b72eec2c9fa3592040e872f65dd172a.
//
// Solidity: event RepayBorrow(address payer, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows, address borrower)
func (_Borrowers *BorrowersFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *BorrowersRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersRepayBorrow)
				if err := _Borrowers.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0xf2e05b26b4020b61a94725f7a9bd18e24b72eec2c9fa3592040e872f65dd172a.
//
// Solidity: event RepayBorrow(address payer, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows, address borrower)
func (_Borrowers *BorrowersFilterer) ParseRepayBorrow(log types.Log) (*BorrowersRepayBorrow, error) {
	event := new(BorrowersRepayBorrow)
	if err := _Borrowers.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Borrowers contract.
type BorrowersReservesAddedIterator struct {
	Event *BorrowersReservesAdded // Event containing the contract specifics and raw log

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
func (it *BorrowersReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersReservesAdded)
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
		it.Event = new(BorrowersReservesAdded)
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
func (it *BorrowersReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersReservesAdded represents a ReservesAdded event raised by the Borrowers contract.
type BorrowersReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Borrowers *BorrowersFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*BorrowersReservesAddedIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &BorrowersReservesAddedIterator{contract: _Borrowers.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Borrowers *BorrowersFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *BorrowersReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersReservesAdded)
				if err := _Borrowers.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Borrowers *BorrowersFilterer) ParseReservesAdded(log types.Log) (*BorrowersReservesAdded, error) {
	event := new(BorrowersReservesAdded)
	if err := _Borrowers.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Borrowers contract.
type BorrowersReservesReducedIterator struct {
	Event *BorrowersReservesReduced // Event containing the contract specifics and raw log

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
func (it *BorrowersReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersReservesReduced)
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
		it.Event = new(BorrowersReservesReduced)
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
func (it *BorrowersReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersReservesReduced represents a ReservesReduced event raised by the Borrowers contract.
type BorrowersReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Borrowers *BorrowersFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*BorrowersReservesReducedIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &BorrowersReservesReducedIterator{contract: _Borrowers.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Borrowers *BorrowersFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *BorrowersReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersReservesReduced)
				if err := _Borrowers.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Borrowers *BorrowersFilterer) ParseReservesReduced(log types.Log) (*BorrowersReservesReduced, error) {
	event := new(BorrowersReservesReduced)
	if err := _Borrowers.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Borrowers contract.
type BorrowersRoleAdminChangedIterator struct {
	Event *BorrowersRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BorrowersRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersRoleAdminChanged)
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
		it.Event = new(BorrowersRoleAdminChanged)
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
func (it *BorrowersRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersRoleAdminChanged represents a RoleAdminChanged event raised by the Borrowers contract.
type BorrowersRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Borrowers *BorrowersFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BorrowersRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BorrowersRoleAdminChangedIterator{contract: _Borrowers.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Borrowers *BorrowersFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BorrowersRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersRoleAdminChanged)
				if err := _Borrowers.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Borrowers *BorrowersFilterer) ParseRoleAdminChanged(log types.Log) (*BorrowersRoleAdminChanged, error) {
	event := new(BorrowersRoleAdminChanged)
	if err := _Borrowers.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Borrowers contract.
type BorrowersRoleGrantedIterator struct {
	Event *BorrowersRoleGranted // Event containing the contract specifics and raw log

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
func (it *BorrowersRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersRoleGranted)
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
		it.Event = new(BorrowersRoleGranted)
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
func (it *BorrowersRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersRoleGranted represents a RoleGranted event raised by the Borrowers contract.
type BorrowersRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Borrowers *BorrowersFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BorrowersRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BorrowersRoleGrantedIterator{contract: _Borrowers.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Borrowers *BorrowersFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BorrowersRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersRoleGranted)
				if err := _Borrowers.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Borrowers *BorrowersFilterer) ParseRoleGranted(log types.Log) (*BorrowersRoleGranted, error) {
	event := new(BorrowersRoleGranted)
	if err := _Borrowers.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Borrowers contract.
type BorrowersRoleRevokedIterator struct {
	Event *BorrowersRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BorrowersRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersRoleRevoked)
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
		it.Event = new(BorrowersRoleRevoked)
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
func (it *BorrowersRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersRoleRevoked represents a RoleRevoked event raised by the Borrowers contract.
type BorrowersRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Borrowers *BorrowersFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BorrowersRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BorrowersRoleRevokedIterator{contract: _Borrowers.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Borrowers *BorrowersFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BorrowersRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersRoleRevoked)
				if err := _Borrowers.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Borrowers *BorrowersFilterer) ParseRoleRevoked(log types.Log) (*BorrowersRoleRevoked, error) {
	event := new(BorrowersRoleRevoked)
	if err := _Borrowers.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Borrowers contract.
type BorrowersSupplyIterator struct {
	Event *BorrowersSupply // Event containing the contract specifics and raw log

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
func (it *BorrowersSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersSupply)
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
		it.Event = new(BorrowersSupply)
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
func (it *BorrowersSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersSupply represents a Supply event raised by the Borrowers contract.
type BorrowersSupply struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x9d624b01705090e0c402c84f2bdef9c498399a009cee9a4e08ab004bae447121.
//
// Solidity: event Supply(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Borrowers *BorrowersFilterer) FilterSupply(opts *bind.FilterOpts) (*BorrowersSupplyIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &BorrowersSupplyIterator{contract: _Borrowers.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x9d624b01705090e0c402c84f2bdef9c498399a009cee9a4e08ab004bae447121.
//
// Solidity: event Supply(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Borrowers *BorrowersFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *BorrowersSupply) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersSupply)
				if err := _Borrowers.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x9d624b01705090e0c402c84f2bdef9c498399a009cee9a4e08ab004bae447121.
//
// Solidity: event Supply(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Borrowers *BorrowersFilterer) ParseSupply(log types.Log) (*BorrowersSupply, error) {
	event := new(BorrowersSupply)
	if err := _Borrowers.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersTokensFrozenIterator is returned from FilterTokensFrozen and is used to iterate over the raw logs and unpacked data for TokensFrozen events raised by the Borrowers contract.
type BorrowersTokensFrozenIterator struct {
	Event *BorrowersTokensFrozen // Event containing the contract specifics and raw log

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
func (it *BorrowersTokensFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersTokensFrozen)
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
		it.Event = new(BorrowersTokensFrozen)
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
func (it *BorrowersTokensFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersTokensFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersTokensFrozen represents a TokensFrozen event raised by the Borrowers contract.
type BorrowersTokensFrozen struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTokensFrozen is a free log retrieval operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address arg0, uint256 arg1)
func (_Borrowers *BorrowersFilterer) FilterTokensFrozen(opts *bind.FilterOpts) (*BorrowersTokensFrozenIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "TokensFrozen")
	if err != nil {
		return nil, err
	}
	return &BorrowersTokensFrozenIterator{contract: _Borrowers.contract, event: "TokensFrozen", logs: logs, sub: sub}, nil
}

// WatchTokensFrozen is a free log subscription operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address arg0, uint256 arg1)
func (_Borrowers *BorrowersFilterer) WatchTokensFrozen(opts *bind.WatchOpts, sink chan<- *BorrowersTokensFrozen) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "TokensFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersTokensFrozen)
				if err := _Borrowers.contract.UnpackLog(event, "TokensFrozen", log); err != nil {
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

// ParseTokensFrozen is a log parse operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address arg0, uint256 arg1)
func (_Borrowers *BorrowersFilterer) ParseTokensFrozen(log types.Log) (*BorrowersTokensFrozen, error) {
	event := new(BorrowersTokensFrozen)
	if err := _Borrowers.contract.UnpackLog(event, "TokensFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersTokensUnfrozenIterator is returned from FilterTokensUnfrozen and is used to iterate over the raw logs and unpacked data for TokensUnfrozen events raised by the Borrowers contract.
type BorrowersTokensUnfrozenIterator struct {
	Event *BorrowersTokensUnfrozen // Event containing the contract specifics and raw log

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
func (it *BorrowersTokensUnfrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersTokensUnfrozen)
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
		it.Event = new(BorrowersTokensUnfrozen)
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
func (it *BorrowersTokensUnfrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersTokensUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersTokensUnfrozen represents a TokensUnfrozen event raised by the Borrowers contract.
type BorrowersTokensUnfrozen struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTokensUnfrozen is a free log retrieval operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address arg0, uint256 arg1)
func (_Borrowers *BorrowersFilterer) FilterTokensUnfrozen(opts *bind.FilterOpts) (*BorrowersTokensUnfrozenIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "TokensUnfrozen")
	if err != nil {
		return nil, err
	}
	return &BorrowersTokensUnfrozenIterator{contract: _Borrowers.contract, event: "TokensUnfrozen", logs: logs, sub: sub}, nil
}

// WatchTokensUnfrozen is a free log subscription operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address arg0, uint256 arg1)
func (_Borrowers *BorrowersFilterer) WatchTokensUnfrozen(opts *bind.WatchOpts, sink chan<- *BorrowersTokensUnfrozen) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "TokensUnfrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersTokensUnfrozen)
				if err := _Borrowers.contract.UnpackLog(event, "TokensUnfrozen", log); err != nil {
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

// ParseTokensUnfrozen is a log parse operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address arg0, uint256 arg1)
func (_Borrowers *BorrowersFilterer) ParseTokensUnfrozen(log types.Log) (*BorrowersTokensUnfrozen, error) {
	event := new(BorrowersTokensUnfrozen)
	if err := _Borrowers.contract.UnpackLog(event, "TokensUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Borrowers contract.
type BorrowersTransferIterator struct {
	Event *BorrowersTransfer // Event containing the contract specifics and raw log

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
func (it *BorrowersTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersTransfer)
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
		it.Event = new(BorrowersTransfer)
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
func (it *BorrowersTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersTransfer represents a Transfer event raised by the Borrowers contract.
type BorrowersTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Borrowers *BorrowersFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BorrowersTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BorrowersTransferIterator{contract: _Borrowers.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Borrowers *BorrowersFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BorrowersTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersTransfer)
				if err := _Borrowers.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Borrowers *BorrowersFilterer) ParseTransfer(log types.Log) (*BorrowersTransfer, error) {
	event := new(BorrowersTransfer)
	if err := _Borrowers.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowersUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Borrowers contract.
type BorrowersUnpausedIterator struct {
	Event *BorrowersUnpaused // Event containing the contract specifics and raw log

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
func (it *BorrowersUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowersUnpaused)
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
		it.Event = new(BorrowersUnpaused)
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
func (it *BorrowersUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowersUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowersUnpaused represents a Unpaused event raised by the Borrowers contract.
type BorrowersUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Borrowers *BorrowersFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BorrowersUnpausedIterator, error) {

	logs, sub, err := _Borrowers.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BorrowersUnpausedIterator{contract: _Borrowers.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Borrowers *BorrowersFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BorrowersUnpaused) (event.Subscription, error) {

	logs, sub, err := _Borrowers.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowersUnpaused)
				if err := _Borrowers.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Borrowers *BorrowersFilterer) ParseUnpaused(log types.Log) (*BorrowersUnpaused, error) {
	event := new(BorrowersUnpaused)
	if err := _Borrowers.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
