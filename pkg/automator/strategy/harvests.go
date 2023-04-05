package strategy

import (
	"context"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func filterHarvests(harvests []strategy.AbstractStrategyHarvest) (ret []strategy.AbstractStrategyHarvest) {
	zero := common.HexToAddress("0x0000000000000000000000000000000000000000")

	for _, h := range harvests {
		if h.Token == zero {
			continue
		}

		if h.Amount.Int64() == 0 {
			continue
		}

		ret = append(ret, h)
	}

	return
}

func removeDuplicates(harvests []strategy.AbstractStrategyHarvest) (ret []strategy.AbstractStrategyHarvest) {
	seen := make(map[common.Address]bool)

	for _, h := range harvests {
		if seen[h.Token] {
			continue
		}

		seen[h.Token] = true
		ret = append(ret, h)
	}

	return
}

func SimulateClaim(
	backend bind.ContractBackend,
	strategyAddress common.Address,
	pool *big.Int,
	safe common.Address,
) ([]strategy.AbstractStrategyHarvest, error) {

	callData, err := utils.EncodeClaimCalldata(pool, safe, make([]byte, 0))
	if err != nil {
		return nil, err
	}

	callMsg := ethereum.CallMsg{
		To:   &strategyAddress,
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Data: callData,
	}

	response, err := backend.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}

	harvests, err := utils.DecodeHarvests(response)
	if err != nil {
		return nil, err
	}

	filteredHarvests := filterHarvests(harvests)

	return removeDuplicates(filteredHarvests), nil
}
