package balances

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/dialecticch/medici-go/pkg/balances/writers"
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/types"
)

// Watcher watches for Withdraw / Deposit events and pipes them to an output channel.
type Watcher struct {
	mux sync.RWMutex

	repo    *repositories.StrategyRepository
	backend bind.ContractBackend

	subscriptions map[string]event.Subscription

	depositSink  chan *strategy.StrategyDeposited
	withdrawSink chan *strategy.StrategyWithdrew
}

// NewWatcher returns a watcher for the specified safes and modules.
func NewWatcher(repo *repositories.StrategyRepository, backend bind.ContractBackend) *Watcher {
	return &Watcher{
		repo:          repo,
		backend:       backend,
		subscriptions: make(map[string]event.Subscription),
		depositSink:   make(chan *strategy.StrategyDeposited, 10),
		withdrawSink:  make(chan *strategy.StrategyWithdrew, 10),
	}
}

// Watch is a Deposit / Withdraw log subscription binding to a set of contracts.
func (w *Watcher) Watch(sink chan *writers.Event) (event.Subscription, error) {
	w.repo.Map(w.register)

	w.repo.OnUpdate(func(strategies []*types.StrategyConfig) {
		for _, s := range strategies {
			w.register(s)
		}
	})

	return event.NewSubscription(func(quit <-chan struct{}) error {
		for {
			select {
			case evt := <-w.depositSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Safe,
					Strategy:    evt.Raw.Address,
					UpdateType:  writers.DEPOSIT,
					Amount:      evt.Amount,
					Block:       evt.Raw.BlockNumber,
					Pool:        evt.Pool,
				}
			case evt := <-w.withdrawSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Safe,
					Strategy:    evt.Raw.Address,
					UpdateType:  writers.WITHDRAW,
					Amount:      evt.Amount,
					Block:       evt.Raw.BlockNumber,
					Pool:        evt.Pool,
				}
			case <-quit:
				w.unsubscribe()
				return nil
			default:
				w.mux.RLock()
				subs := w.subscriptions
				w.mux.RUnlock()

				for _, sub := range subs {
					select {
					case err := <-sub.Err():
						return err
					default:
						continue
					}
				}
			}
		}
	}), nil
}

func (w *Watcher) register(s *types.StrategyConfig) {
	key := key(s)

	w.mux.RLock()
	_, ok := w.subscriptions[key+"_deposit"]
	w.mux.RUnlock()

	if ok {
		return
	}

	module, err := strategy.NewStrategy(s.Strategy.Address, w.backend)
	if err != nil {
		log.Printf("strategy.NewStrategy err: %s", err)
		return
	}

	name, _ := module.NAME(nil)
	log.Printf("watching %s (pool: %s) on %s", name, s.Pool.String(), s.Safe)

	deposits, err := module.WatchDeposited(
		&bind.WatchOpts{},
		w.depositSink,
		[]common.Address{s.Safe},
	)

	if err != nil {
		log.Printf("module.WatchDeposited err: %s", err)
		w.unsubscribe()
		return
	}

	withdraw, err := module.WatchWithdrew(
		&bind.WatchOpts{},
		w.withdrawSink,
		[]common.Address{s.Safe},
	)

	if err != nil {
		log.Printf("module.WatchWithdrew err: %s", err)
		w.unsubscribe()
		return
	}

	w.mux.Lock()
	w.subscriptions[key+"_deposit"] = deposits
	w.subscriptions[key+"_withdraw"] = withdraw
	w.mux.Unlock()
}

func (w *Watcher) unsubscribe() {
	for _, sub := range w.subscriptions {
		sub.Unsubscribe()
	}
}

func key(s *types.StrategyConfig) string {
	return fmt.Sprintf("%s_%s", s.Safe.String(), s.Strategy.Address.String())
}
