package strategy

import (
	"math/big"
	"testing"

	"github.com/dialecticch/medici-go/pkg/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSimulateClaim(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	b := testutils.SetupBackend(privateKey)

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	strat, _, _, err := testutils.DeployMockStrategy(transactOpts, b)
	if err != nil {
		t.Fatal(err)
	}

	b.Commit()

	claims, err := SimulateClaim(
		b,
		strat,
		big.NewInt(0),
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(claims) != 2 {
		t.Fatal("wrong amount")
	}
}
