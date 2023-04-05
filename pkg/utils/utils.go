package utils

import (
	"math/big"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// SlippageMillis adjusts an amount by a given slippage
func SlippageMillis(amount *big.Int, slippage int64) *big.Int {
	return new(big.Int).Div(
		new(big.Int).Mul(amount, big.NewInt(1000-slippage)),
		big.NewInt(1000),
	)
}

var (
	uint256Ty, _ = abi.NewType("uint256", "", nil)
	addressTy, _ = abi.NewType("address", "", nil)
	bytesTy, _   = abi.NewType("bytes", "", nil)
	claimsArgs   = abi.Arguments{
		{Name: "pool", Type: uint256Ty},
		{Name: "safe", Type: addressTy},
		{Name: "_data", Type: bytesTy},
	}

	harvestType, _ = abi.NewType("tuple[]", "StrategyHarvest[]", []abi.ArgumentMarshaling{
		{Name: "Token", Type: "address"},
		{Name: "Amount", Type: "uint256"},
	})
	harvestArgs = abi.Arguments{
		{Name: "harvest", Type: harvestType},
	}
)

func EncodeClaimCalldata(pool *big.Int, safe common.Address, data []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash([]byte("simulateClaim(uint256,address,bytes)"))
	selector := hash[:4]
	argData, err := claimsArgs.Pack(pool, safe, data)
	if err != nil {
		return nil, err
	}
	calldata := append(selector, argData...)
	return calldata, nil
}

func DecodeHarvests(resp []byte) ([]strategy.AbstractStrategyHarvest, error) {
	decoded, err := harvestArgs.UnpackValues(resp)
	if err != nil {
		return nil, err
	}

	harvests := make([]strategy.AbstractStrategyHarvest, 0)

	for _, t := range decoded[0].([]struct {
		Token  common.Address `json:"Token"`
		Amount *big.Int       `json:"Amount"`
	}) {
		harvests = append(harvests, strategy.AbstractStrategyHarvest{
			Token:  t.Token,
			Amount: t.Amount,
		})
	}

	return harvests, err
}

func StringsToAddrs(strs []string) []common.Address {
	ret := make([]common.Address, 0)

	for _, str := range strs {
		ret = append(ret, common.HexToAddress(str))
	}

	return ret
}
