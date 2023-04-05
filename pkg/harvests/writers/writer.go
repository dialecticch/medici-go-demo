package writers

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Event struct {
	HarvestedAt time.Time
	Transaction common.Hash
	LogIndex    uint64
	Safe        common.Address
	Strategy    common.Address
	Token       common.Address
	Amount      *big.Int
	Pool        *big.Int
	Block       uint64
}

type Writer interface {
	Write(*Event) error
}
