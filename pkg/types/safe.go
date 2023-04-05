package types

import (
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/ethereum/go-ethereum/common"
)

type Safe struct {
	Address    common.Address
	Strategies []*strategy.Strategy
}
