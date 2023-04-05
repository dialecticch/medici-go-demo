package balances

import (
	"log"

	"github.com/dialecticch/medici-go/pkg/balances/writers"
)

// Consumer receives Watcher events and forwards them to Writers.
type Consumer struct {
	watcher *Watcher

	sink chan *writers.Event

	writers []writers.Writer
}

// NewConsumer returns a consumer for the specified Watcher that writes to the given Writers.
func NewConsumer(w *Watcher, ws []writers.Writer) *Consumer {
	return &Consumer{
		watcher: w,
		sink:    make(chan *writers.Event),
		writers: ws,
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

func (c *Consumer) handle(evt *writers.Event) {
	for _, w := range c.writers {
		err := w.Write(evt)
		if err != nil {
			log.Printf("w.Write err: %s", err)
		}
	}
}
