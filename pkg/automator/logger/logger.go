package logger

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Logger interface {
	Failed(*types.Receipt) error
	SendError(error, common.Address) error
	MaxRetriesReached(common.Address, common.Address, *big.Int) error
}
