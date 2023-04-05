package strategy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Strategy struct {
	addr, safe common.Address

	pool *big.Int

	transactionBuilder TransactionBuilderFunc
	threshold          ThresholdFunc
}

func (s *Strategy) Safe() common.Address {
	return s.safe
}

func (s *Strategy) Address() common.Address {
	return s.addr
}

func (s *Strategy) Pool() *big.Int {
	return s.pool
}

func (s *Strategy) HasReachedThreshold(tx *types.Transaction) (bool, error) {
	return s.threshold(tx)
}

func (s *Strategy) GetTransaction() (*types.Transaction, error) {
	return s.transactionBuilder()
}
