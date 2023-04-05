package sender

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/dialecticch/medici-go/pkg/testutils"
)

func TestTransactionSend(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	b := testutils.SetupBackend(privateKey)
	ts := NewTransactionSender(b)

	addr := common.HexToAddress("0xffe5a180e035a5f8c5f1201c587042d2cea1584a")

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(875000000),
		Gas:      21000,
		To:       &addr,
		Value:    big.NewInt(1),
		Data:     nil,
	})

	signer := types.NewLondonSigner(big.NewInt(1337))
	signedTx, _ := types.SignTx(tx, signer, privateKey)

	c := make(chan *types.Receipt)
	go func() {
		receipt, err := ts.Send(signedTx)
		if err != nil {
			t.Error(err)
			return
		}
		c <- receipt
	}()

	time.Sleep(1 * time.Second)

	b.Commit()

	<-c
}
