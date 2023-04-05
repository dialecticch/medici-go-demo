package balances_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/dialecticch/medici-go/pkg/balances"
	"github.com/dialecticch/medici-go/pkg/balances/writers"
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

	w := balances.NewWatcher(repo, b)

	sink := make(chan *writers.Event)
	s, err := w.Watch(sink)
	if err != nil {
		t.Fatal(err)
	}

	opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	pool := common.Big1

	_, err = emitter.LogDeposited(opts, pool, safe, amount)
	if err != nil {
		t.Fatal(err)
	}
	b.Commit()

	timeout := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- true
	}()

	select {
	case evt := <-sink:
		if evt.Safe != safe {
			t.Fatalf("safe is not same as expected actual: %s expected: %s", evt.Safe.String(), safe.String())
		}

		if evt.UpdateType != writers.DEPOSIT {
			t.Fatalf("update type is not same as expected actual: %s expected: %s", evt.UpdateType, writers.DEPOSIT)
		}

		if evt.Amount.Cmp(amount) != 0 {
			t.Fatalf("amount is not same as expected actual: %s expected: %s", evt.Amount.String(), amount.String())
		}
	case <-timeout:
		t.Fatal("timed out")
	}

	_, err = emitter.LogWithdrew(opts, pool, safe, amount)
	if err != nil {
		t.Fatal(err)
	}
	b.Commit()

	go func() {
		time.Sleep(5 * time.Second)
		timeout <- true
	}()

	select {
	case evt := <-sink:
		if evt.Safe != safe {
			t.Fatalf("safe is not same as expected actual: %s expected: %s", evt.Safe.String(), safe.String())
		}

		if evt.UpdateType != writers.WITHDRAW {
			t.Fatalf("update type is not same as expected actual: %s expected: %s", evt.UpdateType, writers.WITHDRAW)
		}

		if evt.Amount.Cmp(amount) != 0 {
			t.Fatalf("amount is not same as expected actual: %s expected: %s", evt.Amount.String(), amount.String())
		}
	case <-timeout:
		t.Fatal("timed out")
	}

	s.Unsubscribe()
}
