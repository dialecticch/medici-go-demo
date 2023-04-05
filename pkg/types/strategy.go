package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// StrategyMetadata represents the metadata for a medici strategy
type StrategyMetadata struct {
	Address common.Address
	Name    string
}

// StrategyConfig represents the configuration of a specific medici strategy for a specific safe
type StrategyConfig struct {
	Safe          common.Address
	DepositToken  common.Address
	Strategy      StrategyMetadata
	Threshold     float64
	Action        string
	ThresholdType string
	LastHarvested big.Int
	Pool          *big.Int
}
