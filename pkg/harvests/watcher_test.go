package harvests_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dialecticch/medici-go/pkg/contracts/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/dialecticch/medici-go/pkg/harvests"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/testutils"
)

func TestWatcher_Watch(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	b := testutils.SetupBackend(privateKey)

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	addr, _, emitter, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal(err)
	}
	b.Commit()

	safe := common.HexToAddress("0xffe5a180e035a5f8c5f1201c587042d2cea1584a")
	token := common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	amount := big.NewInt(12)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("^SELECT (.+)").
		WithArgs(1337, true).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow("0xffe5a180e035a5f8c5f1201c587042d2cea1584a", addr.String(), "foo", "0.15", "harvest", "gas_percentage", "0", "1", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))

	repo := repositories.NewStrategyRepository(db, 1337)
	err = repo.Run(repositories.Active)
	if err != nil {
		t.Fatal(err)
	}

	w := harvests.NewWatcher(repo, b)

	sink := make(chan *strategy.StrategyHarvested)
	s, err := w.Watch(sink)
	if err != nil {
		t.Fatal(err)
	}

	opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	pool := common.Big1

	_, err = emitter.LogHarvested(opts, pool, safe, token, amount)
	if err != nil {
		t.Fatal(err)
	}
	b.Commit()

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(15 * time.Second)
		timeout <- true
	}()

	select {
	case evt := <-sink:
		if evt.Safe != safe {
			t.Fatalf("safe is not same as expected actual: %s expected: %s", evt.Safe.String(), safe.String())
		}

		if evt.Token != token {
			t.Fatalf("token is not same as expected actual: %s expected: %s", evt.Token.String(), token.String())
		}

		if evt.Amount.Cmp(amount) != 0 {
			t.Fatalf("amount is not same as expected actual: %s expected: %s", evt.Amount.String(), amount.String())
		}
	case <-timeout:
		t.Fatal("timed out")
	}

	s.Unsubscribe()
}
