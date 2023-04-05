package sender

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractTransactorAndMiner interface {
	bind.ContractTransactor
	bind.DeployBackend
}

// TransactionSender sends arbitrary transactions on the Ethereum blockchain and returns their receipt.
type TransactionSender struct {
	backend ContractTransactorAndMiner
}

func NewTransactionSender(backend ContractTransactorAndMiner) *TransactionSender {
	return &TransactionSender{
		backend: backend,
	}
}

func (t *TransactionSender) Send(transaction *types.Transaction) (*types.Receipt, error) {
	err := t.backend.SendTransaction(context.Background(), transaction)
	if err != nil {
		return nil, err
	}
	return bind.WaitMined(context.Background(), t.backend, transaction)
}
