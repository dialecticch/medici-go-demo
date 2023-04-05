package balances_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"

	"github.com/dialecticch/medici-go/pkg/balances"
	"github.com/dialecticch/medici-go/pkg/balances/testdata"
	"github.com/dialecticch/medici-go/pkg/balances/testdata/match"
	"github.com/dialecticch/medici-go/pkg/balances/writers"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/testutils"
)

func TestBackfiller_RunDeposits(t *testing.T) {
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
		WithArgs(1337).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow("0xffe5a180e035a5f8c5f1201c587042d2cea1584a", addr.String(), "foo", "0.15", "harvest", "gas_percentage", "0", "1", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))

	repo := repositories.NewStrategyRepository(db, 1337)
	strats, err := repo.Query(repositories.Any)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)

	writer := testdata.NewMockWriter(ctrl)

	bf := balances.NewBackfiller(b, []writers.Writer{writer})

	opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	pool := common.Big1

	_, err = emitter.LogDeposited(opts, pool, safe, amount)
	if err != nil {
		t.Fatal(err)
	}
	b.Commit()

	errs := make(chan error)
	go func() {
		err := bf.RunDeposits(1, strats)
		errs <- err
	}()

	success := make(chan bool)
	writer.EXPECT().Write(match.Event(pool, amount, safe, addr)).DoAndReturn(func(*writers.Event) error {
		success <- true
		return nil
	})

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(15 * time.Second)
		timeout <- true
	}()

	select {
	case <-success:
		return
	case <-timeout:
		t.Fatal("timed out")
	case err := <-errs:
		t.Fatalf("errored %s", err)
	}
}

func TestBackfiller_RunWithdraws(t *testing.T) {
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
		WithArgs(1337).
		WillReturnRows(sqlmock.NewRows([]string{"safe", "strategy", "name", "threshold", "action", "threshold_type", "last_harvested", "pool", "deposit_token"}).AddRow("0xffe5a180e035a5f8c5f1201c587042d2cea1584a", addr.String(), "foo", "0.15", "harvest", "gas_percentage", "0", "1", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))

	repo := repositories.NewStrategyRepository(db, 1337)
	strats, err := repo.Query(repositories.Any)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)

	writer := testdata.NewMockWriter(ctrl)

	bf := balances.NewBackfiller(b, []writers.Writer{writer})

	pool := common.Big1

	opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	_, err = emitter.LogWithdrew(opts, pool, safe, amount)
	if err != nil {
		t.Fatal(err)
	}
	b.Commit()

	errs := make(chan error)
	go func() {
		err := bf.RunWithdraws(1, strats)
		errs <- err
	}()

	success := make(chan bool)
	writer.EXPECT().Write(match.Event(pool, amount, safe, addr)).DoAndReturn(func(*writers.Event) error {
		success <- true
		return nil
	})

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(15 * time.Second)
		timeout <- true
	}()

	select {
	case <-success:
		return
	case <-timeout:
		t.Fatal("timed out")
	case err := <-errs:
		t.Fatalf("errored %s", err)
	}
}
