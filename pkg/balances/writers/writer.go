package writers

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type UpdateType string

const (
	DEPOSIT  = "deposit"
	WITHDRAW = "withdraw"
)

type Event struct {
	UpdatedAt   time.Time
	Transaction common.Hash
	LogIndex    uint64
	Safe        common.Address
	Strategy    common.Address
	UpdateType  UpdateType
	Amount      *big.Int
	Pool        *big.Int
	Block       uint64
}

type Writer interface {
	Write(*Event) error
}
