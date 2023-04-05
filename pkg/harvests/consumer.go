package harvests

import (
	"log"
	"time"

	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/dialecticch/medici-go/pkg/harvests/writers"
)

// Consumer receives Watcher events and forwards them to Writers.
type Consumer struct {
	watcher *Watcher

	sink chan *strategy.StrategyHarvested

	writers []writers.Writer
}

// NewConsumer returns a consumer for the specified Watcher that writes to the given Writers.
func NewConsumer(w *Watcher, writers []writers.Writer) *Consumer {
	return &Consumer{
		watcher: w,
		sink:    make(chan *strategy.StrategyHarvested),
		writers: writers,
	}
}

// Run runs the event processing and forwarding.
func (c *Consumer) Run() error {
	sub, err := c.watcher.Watch(c.sink)
	if err != nil {
		return err
	}

	for {
		select {
		case evt := <-c.sink:
			go c.handle(evt)
		case err := <-sub.Err():
			sub.Unsubscribe()
			return err
		}
	}
}

func (c *Consumer) handle(evt *strategy.StrategyHarvested) {
	out := &writers.Event{
		HarvestedAt: time.Now(),
		Transaction: evt.Raw.TxHash,
		LogIndex:    uint64(evt.Raw.Index),
		Safe:        evt.Safe,
		Strategy:    evt.Raw.Address,
		Token:       evt.Token,
		Amount:      evt.Amount,
		Block:       evt.Raw.BlockNumber,
		Pool:        evt.Pool,
	}

	for _, w := range c.writers {
		err := w.Write(out)
		if err != nil {
			log.Printf("w.Write err: %s", err)
		}
	}
}
