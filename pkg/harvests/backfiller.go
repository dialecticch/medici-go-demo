package harvests

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/harvests/writers"
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

func (b *Backfiller) Run(startBlock uint64, strategies []*types.StrategyConfig) error {
	log.Println("Starting backfill from block", startBlock)
	for _, s := range strategies {
		m, err := strategy.NewStrategy(s.Strategy.Address, b.backend)
		if err != nil {
			log.Printf("strategy.NewStrategy err: %s", err)
			return err
		}

		log.Println("Finding events from", s.Strategy.Address)

		logs, err := m.FilterHarvested(&bind.FilterOpts{
			Start: startBlock,
		}, nil, nil)
		if err != nil {
			log.Printf("FilterHarvested err: %s", err)
			return err
		}

		for logs.Next() {
			l := logs.Event
			log.Printf("block: %d safe: %s token: %s amount: %s", l.Raw.BlockNumber, l.Safe, l.Token, l.Amount)

			header, err := b.backend.HeaderByNumber(context.Background(), new(big.Int).SetUint64(l.Raw.BlockNumber))
			if err != nil {
				log.Printf("Error getting block header %s", err)
				return err
			}

			for _, w := range b.writers {
				err = w.Write(&writers.Event{
					HarvestedAt: time.Unix(int64(header.Time), 0),
					Transaction: l.Raw.TxHash,
					LogIndex:    uint64(l.Raw.Index),
					Safe:        l.Safe,
					Strategy:    l.Raw.Address,
					Token:       l.Token,
					Amount:      l.Amount,
					Block:       l.Raw.BlockNumber,
					Pool:        l.Pool,
				})
				if err != nil {
					if !strings.Contains(err.Error(), "duplicate key value") {
						log.Println("Writing failed", err)
						fmt.Printf("w.Write err: %s safe: %s strategy: %s gracefully continuing\n",
							err,
							l.Safe.String(),
							l.Raw.Address.String(),
						)
					}
				}
			}

		}
	}

	return nil
}
