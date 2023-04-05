package harvests

import (
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/types"
)

// Watcher watches for harvested events and pipes them to an output channel.
type Watcher struct {
	mux sync.RWMutex

	repo    *repositories.StrategyRepository
	backend bind.ContractBackend

	subscriptions map[string]event.Subscription
	sink          chan *strategy.StrategyHarvested
}

// NewWatcher returns a watcher for the specified safes and modules.
func NewWatcher(repo *repositories.StrategyRepository, backend bind.ContractBackend) *Watcher {
	return &Watcher{
		repo:          repo,
		backend:       backend,
		subscriptions: make(map[string]event.Subscription),
		sink:          make(chan *strategy.StrategyHarvested, 10),
	}
}

// Watch is a Harvested log subscription binding to a set of contracts.
func (w *Watcher) Watch(sink chan *strategy.StrategyHarvested) (event.Subscription, error) {
	w.repo.Map(w.register)

	w.repo.OnUpdate(func(strategies []*types.StrategyConfig) {
		for _, s := range strategies {
			w.register(s)
		}
	})

	return event.NewSubscription(func(quit <-chan struct{}) error {
		for {
			select {
			case evt := <-w.sink:
				sink <- evt
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
	_, ok := w.subscriptions[key]
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

	sub, err := module.WatchHarvested(
		nil,
		w.sink,
		[]common.Address{s.Safe},
		nil,
	)

	if err != nil {
		log.Printf("module.WatchHarvested err: %s", err)
		w.unsubscribe()
		return
	}

	w.mux.Lock()
	w.subscriptions[key] = sub
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
