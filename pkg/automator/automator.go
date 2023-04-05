package automator

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dialecticch/medici-go/pkg/automator/logger"
	"github.com/dialecticch/medici-go/pkg/automator/strategy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/sync/semaphore"

	"github.com/dialecticch/medici-go/pkg/automator/sender"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Automator struct {
	mux sync.Mutex

	sink chan *types.Header

	repository *repositories.StrategyRepository

	semaphores map[string]*semaphore.Weighted

	retries    map[string]int
	maxRetries int

	sender sender.Sender

	key *ecdsa.PrivateKey

	epoch uint64

	builder *strategy.Builder

	logger logger.Logger

	backend bind.ContractBackend

	hasReportedInsufficientFunds int32
}

func NewAutomator(
	sink chan *types.Header,
	repository *repositories.StrategyRepository,
	builder *strategy.Builder,
	sender sender.Sender,
	key *ecdsa.PrivateKey,
	epoch uint64,
	maxRetries int,
	l logger.Logger,
	backend bind.ContractBackend,
) *Automator {
	return &Automator{
		sink:       sink,
		repository: repository,
		builder:    builder,
		semaphores: make(map[string]*semaphore.Weighted),
		retries:    make(map[string]int),
		maxRetries: maxRetries,
		sender:     sender,
		key:        key,
		epoch:      epoch,
		logger:     l,
		backend:    backend,
	}
}

func (a *Automator) Run() error {
	elapsed := uint64(0)

	for header := range a.sink {
		// We only wanna check every N blocks.
		if elapsed == 0 || elapsed%a.epoch == 0 {
			strategies, err := a.repository.Query(repositories.Active)
			if err != nil {
				return err
			}

			log.Printf("running checks for block %s", header.Number.String())

			for _, s := range strategies {
				strat, err := a.builder.Build(s, header.Number)
				if err != nil {
					log.Printf("a.builder.Build err: %s", err)
					continue
				}

				log.Printf("checking %s (pool: %s) on %s", s.Strategy.Name, s.Pool, s.Safe)

				go a.handle(strat)
			}
		}

		elapsed++
	}

	return nil
}

func (a *Automator) handle(strategy *strategy.Strategy) {
	sem := a.semaphore(strategy.Safe(), strategy.Address(), strategy.Pool())

	if !sem.TryAcquire(1) {
		return
	}

	defer sem.Release(1)

	key := strategy.Safe().String() + strategy.Address().String() + strategy.Pool().String()

	if a.retries[key] == a.maxRetries {
		return
	}

	tx, err := strategy.GetTransaction()
	if err != nil {
		log.Printf("a.txForAction err: %s", err)
		return
	}

	ok, err := strategy.HasReachedThreshold(tx)
	if err != nil {
		log.Printf("hasReachedThreshold err: %s to: %s txdata: %s", err, tx.To().Hex(), common.Bytes2Hex(tx.Data()))
		return
	}

	if !ok {
		return
	}

	log.Printf("running action for %s on %s", strategy.Address().Hex(), strategy.Safe().Hex())

	receipt, err := a.sender.Send(tx)
	if err != nil {
		log.Printf("a.sender.Send err: %s", err)

		if strings.Contains(err.Error(), core.ErrInsufficientFunds.Error()) {

			if atomic.LoadInt32(&a.hasReportedInsufficientFunds) == 1 {
				return
			}

			atomic.StoreInt32(&a.hasReportedInsufficientFunds, 1)
		}

		err = a.logger.SendError(err, crypto.PubkeyToAddress(a.key.PublicKey))
		if err != nil {
			log.Printf("a.logger.SendError err: %s", err)
		}

		return
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		a.retries[key] = 0
		atomic.StoreInt32(&a.hasReportedInsufficientFunds, 0)
		return
	}

	err = a.logger.Failed(receipt)
	if err != nil {
		log.Printf("a.logger.Failed err: %s", err)
	}

	a.retries[key] = a.retries[key] + 1
	if a.retries[key] != a.maxRetries {
		return
	}

	err = a.logger.MaxRetriesReached(strategy.Address(), strategy.Safe(), strategy.Pool())
	if err != nil {
		log.Printf("a.logger.Failed err: %s", err)
	}
}

func (a *Automator) semaphore(safe, strategy common.Address, pool *big.Int) *semaphore.Weighted {
	a.mux.Lock()
	defer a.mux.Unlock()

	key := safe.String() + strategy.String() + pool.String()
	s, ok := a.semaphores[key]
	if !ok {
		s = semaphore.NewWeighted(1)
		a.semaphores[key] = s
	}

	return s
}
