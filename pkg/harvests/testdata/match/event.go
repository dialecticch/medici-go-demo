package match

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"

	"github.com/dialecticch/medici-go/pkg/harvests/writers"
)

type eventMatcher struct {
	strategy common.Address
	token    common.Address
	safe     common.Address
	amount   *big.Int
	pool     *big.Int
}

func (e eventMatcher) Matches(x interface{}) bool {
	evt, ok := x.(*writers.Event)
	if !ok {
		return false
	}

	return evt.Strategy.String() == e.strategy.String() &&
		evt.Safe.String() == e.safe.String() &&
		evt.Token.String() == e.token.String() &&
		e.amount.Cmp(evt.Amount) == 0 &&
		e.pool.Cmp(evt.Pool) == 0
}

func (e eventMatcher) String() string {
	return "is of type writers.Event"
}

// Event returns a matcher for a writers.Event
func Event(pool, amount *big.Int, safe, strategy, token common.Address) gomock.Matcher {
	return &eventMatcher{
		pool:     pool,
		amount:   amount,
		safe:     safe,
		strategy: strategy,
		token:    token,
	}
}
