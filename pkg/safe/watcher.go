package safe

import (
	"log"
	"sync"
	"time"

	"github.com/dialecticch/medici-go/pkg/contracts/safe"
	"github.com/dialecticch/medici-go/pkg/safe/writers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Watcher watches for safe events and pipes them to an output channel.
type Watcher struct {
	mux sync.RWMutex

	backend bind.ContractBackend

	safes []common.Address

	subscriptions map[string]event.Subscription

	moduleEnabledSink    chan *safe.GnosisSafeEnabledModule
	moduleDisabledSink   chan *safe.GnosisSafeDisabledModule
	addedOwnerSink       chan *safe.GnosisSafeAddedOwner
	removedOwnerSink     chan *safe.GnosisSafeRemovedOwner
	changedThresholdSink chan *safe.GnosisSafeChangedThreshold
}

// NewWatcher returns a watcher for the specified safes
func NewWatcher(safes []common.Address, backend bind.ContractBackend) *Watcher {
	return &Watcher{
		backend:              backend,
		subscriptions:        make(map[string]event.Subscription),
		safes:                safes,
		moduleEnabledSink:    make(chan *safe.GnosisSafeEnabledModule, len(safes)),
		moduleDisabledSink:   make(chan *safe.GnosisSafeDisabledModule, len(safes)),
		addedOwnerSink:       make(chan *safe.GnosisSafeAddedOwner, len(safes)),
		removedOwnerSink:     make(chan *safe.GnosisSafeRemovedOwner, len(safes)),
		changedThresholdSink: make(chan *safe.GnosisSafeChangedThreshold, len(safes)),
	}
}

// Watch is a log subscription binding to a set of safes.
func (w *Watcher) Watch(sink chan *writers.Event) (event.Subscription, error) {
	for _, addr := range w.safes {
		w.register(addr)
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		for {
			select {
			case evt := <-w.moduleEnabledSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					UpdateType:  writers.EnabledModule,
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Raw.Address,
					Block:       evt.Raw.BlockNumber,
					Arguments:   map[string]string{"module": evt.Module.String()},
				}
			case evt := <-w.moduleDisabledSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					UpdateType:  writers.DisabledModule,
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Raw.Address,
					Block:       evt.Raw.BlockNumber,
					Arguments:   map[string]string{"module": evt.Module.String()},
				}
			case evt := <-w.addedOwnerSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					UpdateType:  writers.AddedOwner,
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Raw.Address,
					Block:       evt.Raw.BlockNumber,
					Arguments:   map[string]string{"owner": evt.Owner.String()},
				}
			case evt := <-w.removedOwnerSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					UpdateType:  writers.RemovedOwner,
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Raw.Address,
					Block:       evt.Raw.BlockNumber,
					Arguments:   map[string]string{"owner": evt.Owner.String()},
				}
			case evt := <-w.changedThresholdSink:
				sink <- &writers.Event{
					UpdatedAt:   time.Now(),
					UpdateType:  writers.ChangedThreshold,
					Transaction: evt.Raw.TxHash,
					LogIndex:    uint64(evt.Raw.Index),
					Safe:        evt.Raw.Address,
					Block:       evt.Raw.BlockNumber,
					Arguments:   map[string]string{"threshold": evt.Threshold.String()},
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

func (w *Watcher) register(addr common.Address) {
	key := addr.String()

	w.mux.RLock()
	_, ok := w.subscriptions[key+"_enabled"]
	w.mux.RUnlock()

	if ok {
		return
	}

	module, err := safe.NewGnosisSafe(addr, w.backend)
	if err != nil {
		log.Printf("safe.NewGnosisSafe err: %s", err)
		return
	}

	log.Printf("watching %s", addr.String())

	enabledModule, err := module.WatchEnabledModule(
		&bind.WatchOpts{},
		w.moduleEnabledSink,
	)

	if err != nil {
		log.Printf("module.WatchEnabledModule err: %s", err)
		w.unsubscribe()
		return
	}

	disabledModule, err := module.WatchDisabledModule(
		&bind.WatchOpts{},
		w.moduleDisabledSink,
	)

	if err != nil {
		log.Printf("module.WatchDisabledModule err: %s", err)
		w.unsubscribe()
		return
	}

	addedOwner, err := module.WatchAddedOwner(
		&bind.WatchOpts{},
		w.addedOwnerSink,
	)

	if err != nil {
		log.Printf("module.WatchAddedOwner err: %s", err)
		w.unsubscribe()
		return
	}

	removedOwner, err := module.WatchRemovedOwner(
		&bind.WatchOpts{},
		w.removedOwnerSink,
	)

	if err != nil {
		log.Printf("module.RemovedOwner err: %s", err)
		w.unsubscribe()
		return
	}

	changedThreshold, err := module.WatchChangedThreshold(
		&bind.WatchOpts{},
		w.changedThresholdSink,
	)

	if err != nil {
		log.Printf("module.WatchChangedThreshold err: %s", err)
		w.unsubscribe()
		return
	}

	w.mux.Lock()
	w.subscriptions[key+"_enabled"] = enabledModule
	w.subscriptions[key+"_disabled"] = disabledModule
	w.subscriptions[key+"_addedOwner"] = addedOwner
	w.subscriptions[key+"_removedOwner"] = removedOwner
	w.subscriptions[key+"_changedThreshold"] = changedThreshold
	w.mux.Unlock()
}

func (w *Watcher) unsubscribe() {
	for _, sub := range w.subscriptions {
		sub.Unsubscribe()
	}
}
