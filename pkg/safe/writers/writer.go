package writers

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type UpdateType int

const (
	EnabledModule UpdateType = iota
	DisabledModule
	AddedOwner
	RemovedOwner
	ChangedThreshold
)

type Event struct {
	UpdatedAt   time.Time
	UpdateType  UpdateType
	Transaction common.Hash
	LogIndex    uint64
	Safe        common.Address
	Block       uint64
	Arguments   map[string]string
}

type Writer interface {
	Write(*Event) error
}
