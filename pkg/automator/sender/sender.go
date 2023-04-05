package sender

import "github.com/ethereum/go-ethereum/core/types"

type Sender interface {
	Send(*types.Transaction) (*types.Receipt, error)
}
