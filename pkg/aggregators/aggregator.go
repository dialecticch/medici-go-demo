package aggregators

import (
	"fmt"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/config"
	"github.com/ethereum/go-ethereum/common"
)

type AggregatorTransaction struct {
	From     common.Address
	To       common.Address
	Data     []byte
	Value    *big.Int
	GasPrice *big.Int
	Gas      int64
}

type SwapArgs struct {
	ChainID     uint64
	From        common.Address
	To          common.Address
	FromAddress common.Address
	Amount      *big.Int
	Slippage    *big.Float
	Opts        interface{}
}

type Swap struct {
	FromToken       common.Address
	ToToken         common.Address
	ToTokenAmount   *big.Int
	FromTokenAmount *big.Int
	Tx              *AggregatorTransaction
}

type QuoteArgs struct {
	ChainID uint64
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Opts    interface{}
}

type Quote struct {
	FromToken       common.Address
	ToToken         common.Address
	ToTokenAmount   *big.Int
	FromTokenAmount *big.Int
	EstimatedGas    int64
}

type Aggregator interface {
	GetSwap(SwapArgs) (*Swap, error)
	GetQuote(QuoteArgs) (*Quote, error)
	GetApprovalSpender(uint64) (*common.Address, error)
}

func NewAggregator(networkCfg *config.Network, aggregatorsCfg *map[string]config.AggregatorConfig, chainID *big.Int) (Aggregator, error) {
	aggregatorCfg := (*aggregatorsCfg)[networkCfg.Aggregator]

	switch networkCfg.Aggregator {
	case "one_inch":
		return NewOneInchAggregator(aggregatorCfg.Name, aggregatorCfg.BaseUrl, chainID.String())
	default:
		return nil, fmt.Errorf("no supported aggregator for %s", networkCfg.Aggregator)
	}
}
