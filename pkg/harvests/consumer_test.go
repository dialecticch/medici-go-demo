package harvests_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"

	"github.com/dialecticch/medici-go/pkg/harvests"
	"github.com/dialecticch/medici-go/pkg/harvests/testdata"
	"github.com/dialecticch/medici-go/pkg/harvests/testdata/match"
	"github.com/dialecticch/medici-go/pkg/harvests/writers"
	"github.com/dialecticch/medici-go/pkg/repositories"
	"github.com/dialecticch/medici-go/pkg/testutils"
)

func TestConsumer_Run(t *testing.T) {
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

	ctrl := gomock.NewController(t)

	writer := testdata.NewMockWriter(ctrl)

	c := harvests.NewConsumer(w, []writers.Writer{writer})

	errs := make(chan error)
	go func() {
		err := c.Run()
		errs <- err
	}()

	time.Sleep(5 * time.Second)

	success := make(chan bool)

	pool := common.Big1

	writer.EXPECT().Write(match.Event(pool, amount, safe, addr, token)).DoAndReturn(func(*writers.Event) error {
		success <- true
		return nil
	})

	opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

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
	case <-success:
		return
	case <-timeout:
		t.Fatal("timed out")
	case err := <-errs:
		t.Fatalf("errored %s", err)
	}
}
