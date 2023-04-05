package balances

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/dialecticch/medici-go/pkg/balances/writers"
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Backfiller struct {
	backend bind.ContractBackend
	writers []writers.Writer
}

// NewBackfiller returns a backfiller for the specified safes and modules.
func NewBackfiller(backend bind.ContractBackend, writers []writers.Writer) *Backfiller {
	return &Backfiller{
		backend: backend,
		writers: writers,
	}
}

func (b *Backfiller) RunDeposits(startBlock uint64, strategies []*types.StrategyConfig) error {
	log.Println("Starting backfill from block", startBlock)
	for _, s := range strategies {
		m, err := strategy.NewStrategy(s.Strategy.Address, b.backend)
		if err != nil {
			log.Printf("contracts.NewModule err: %s", err)
			return err
		}

		log.Println("Finding events from", s.Strategy.Address)

		logs, err := m.FilterDeposited(&bind.FilterOpts{
			Start: startBlock,
		}, nil)
		if err != nil {
			log.Printf("FilterDeposited err: %s", err)
			return err
		}

		for logs.Next() {
			l := logs.Event
			log.Printf("block: %d safe: %s amount: %s", l.Raw.BlockNumber, l.Safe, l.Amount)

			header, err := b.backend.HeaderByNumber(context.Background(), new(big.Int).SetUint64(l.Raw.BlockNumber))
			if err != nil {
				log.Printf("Error getting block header %s", err)
				return err
			}

			for _, w := range b.writers {
				err = w.Write(&writers.Event{
					UpdatedAt:   time.Unix(int64(header.Time), 0),
					Transaction: l.Raw.TxHash,
					LogIndex:    uint64(l.Raw.Index),
					Safe:        l.Safe,
					Strategy:    l.Raw.Address,
					Amount:      l.Amount,
					UpdateType:  writers.DEPOSIT,
					Block:       l.Raw.BlockNumber,
					Pool:        l.Pool,
				})
				if err != nil {
					if !strings.Contains(err.Error(), "duplicate key value") {
						log.Println("Writing failed", err)
						return err
					}
				}
			}

		}
	}

	return nil
}

func (b *Backfiller) RunWithdraws(startBlock uint64, strategies []*types.StrategyConfig) error {
	log.Println("Starting backfill from block", startBlock)
	for _, s := range strategies {
		m, err := strategy.NewStrategy(s.Strategy.Address, b.backend)
		if err != nil {
			log.Printf("contracts.NewModule err: %s", err)
			return err
		}

		log.Println("Finding events from", s.Strategy.Address)

		logs, err := m.FilterWithdrew(&bind.FilterOpts{
			Start: startBlock,
		}, nil)
		if err != nil {
			log.Printf("FilterWithdrew err: %s", err)
			return err
		}

		for logs.Next() {
			l := logs.Event
			log.Printf("block: %d safe: %s amount: %s", l.Raw.BlockNumber, l.Safe, l.Amount)

			header, err := b.backend.HeaderByNumber(context.Background(), new(big.Int).SetUint64(l.Raw.BlockNumber))
			if err != nil {
				log.Printf("Error getting block header %s", err)
				return err
			}

			for _, w := range b.writers {
				err = w.Write(&writers.Event{
					UpdatedAt:   time.Unix(int64(header.Time), 0),
					Transaction: l.Raw.TxHash,
					LogIndex:    uint64(l.Raw.Index),
					Safe:        l.Safe,
					Strategy:    l.Raw.Address,
					Amount:      l.Amount,
					UpdateType:  writers.WITHDRAW,
					Block:       l.Raw.BlockNumber,
					Pool:        l.Pool,
				})
				if err != nil {
					if !strings.Contains(err.Error(), "duplicate key value") {
						log.Println("Writing failed", err)
						return err
					}
				}
			}

		}
	}

	return nil
}
