package aggregators_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dialecticch/medici-go/pkg/aggregators"
	"github.com/ethereum/go-ethereum/common"
)

var (
	oneInchBaseUrl = "https://api.1inch.exchange/v4.0"
	oneInchName    = "1Inch Aggregator"
	chainID        = "1"
	oneInchAddress = common.HexToAddress("0x1111111254fb6c44bac0bed2854e76f90643097d")
	safeAddress    = common.HexToAddress("0x00000000000000000000000000000000000054FE")
	wrapped        = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") // wrapped native token
	stableCoin     = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // mainnet usdc address
)

func TestOneInch_RunGetQuote(t *testing.T) {
	expectedQuote := aggregators.OneInchQuote{
		FromToken: aggregators.OneInchToken{
			Symbol:   "WETH",
			Name:     "Wrapped Ether",
			Address:  wrapped.String(),
			Decimals: 18,
			LogoURI:  "",
			Tags:     nil,
		},
		ToToken: aggregators.OneInchToken{
			Symbol:   "USDC",
			Name:     "USD Coin",
			Address:  stableCoin.String(),
			Decimals: 6,
			LogoURI:  "",
			Tags:     nil,
		},
		ToTokenAmount:   "1926426408",
		FromTokenAmount: "1000000000000000000",
		Protocols:       nil,
		EstimatedGas:    202393,
	}

	expectedQuoteBytes, err := json.Marshal(expectedQuote)
	if err != nil {
		t.Fatal("Fatal error while encoding expected quote: ", err)
	}

	mockOneInchServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(expectedQuoteBytes))
	}))

	oneInch, err := aggregators.NewOneInchAggregator(oneInchName, mockOneInchServer.URL, chainID)
	if err != nil {
		t.Fatal("Fatal error for OneInchAggregator: ", err)
	}

	quote, err := oneInch.GetQuote(aggregators.QuoteArgs{
		ChainID: 1,
		From:    wrapped,
		To:      stableCoin,
		Amount:  big.NewInt(1e18),
		Opts:    aggregators.OneInchQuoteOpts{},
	})

	if quote.FromToken != wrapped {
		t.Fatal("Expected ", wrapped.String(), " got ", quote.FromToken, " for FromToken")
	}

	if quote.ToToken != stableCoin {
		t.Fatal("Expected ", stableCoin.String(), " got ", quote.ToToken, " for ToToken")
	}

	expectedAmountIn := "1000000000000000000"
	if quote.FromTokenAmount.String() != expectedAmountIn {
		t.Fatal("Expected ", expectedAmountIn, " got ", quote.FromTokenAmount, " for FromTokenAmount")
	}

	expectedAmountOut := "1926426408"
	if quote.ToTokenAmount.String() != expectedAmountOut {
		t.Fatal("Expected ", expectedAmountOut, " got ", quote.ToTokenAmount, " for ToTokenAmount")
	}

	expectedGasEstimated := 202393
	if quote.EstimatedGas != int64(expectedGasEstimated) {
		t.Fatal("Expected ", expectedGasEstimated, " got ", quote.EstimatedGas, " for EstimatedGas")
	}

	if err != nil {
		t.Fatal("Fatal error for OneInchAggregator: ", err)
	}
}

func TestOneInch_RunGetSwap(t *testing.T) {
	expectedSwap := aggregators.OneInchSwap{
		FromToken: aggregators.OneInchToken{
			Symbol:   "WETH",
			Name:     "Wrapped Ether",
			Address:  wrapped.String(),
			Decimals: 18,
			LogoURI:  "",
			Tags:     nil,
		},
		ToToken: aggregators.OneInchToken{
			Symbol:   "USDC",
			Name:     "USD Coin",
			Address:  stableCoin.String(),
			Decimals: 6,
			LogoURI:  "",
			Tags:     nil,
		},
		ToTokenAmount:   "1926426408",
		FromTokenAmount: "1000000000000000000",
		Protocols:       nil,
		Tx: aggregators.OneInchTx{
			From:     safeAddress.String(),
			To:       oneInchAddress.String(),
			Data:     "",
			Value:    "0",
			GasPrice: "21261016567",
			Gas:      0,
		},
	}

	expectedSwapBytes, err := json.Marshal(expectedSwap)
	if err != nil {
		t.Fatal("Fatal error while encoding expected quote: ", err)
	}

	mockOneInchServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(expectedSwapBytes))
	}))

	oneInch, err := aggregators.NewOneInchAggregator(oneInchName, mockOneInchServer.URL, chainID)
	if err != nil {
		t.Fatal("Fatal error for OneInchAggregator: ", err)
	}

	swap, err := oneInch.GetSwap(aggregators.SwapArgs{
		ChainID:     1,
		From:        wrapped,
		To:          stableCoin,
		FromAddress: safeAddress,
		Amount:      big.NewInt(1e18),
		Slippage:    big.NewFloat(0.1),
		Opts: aggregators.OneInchSwapOpts{
			DisableEstimate: true,
		},
	})

	if swap.FromToken != wrapped {
		t.Fatal("Expected ", wrapped, " got ", swap.FromToken, " for FromToken")
	}

	if swap.ToToken != stableCoin {
		t.Fatal("Expected ", stableCoin, " got ", swap.ToToken, " for ToToken")
	}

	expectedAmountIn := "1000000000000000000"
	if swap.FromTokenAmount.String() != expectedAmountIn {
		t.Fatal("Expected ", expectedAmountIn, " got ", swap.FromTokenAmount, " for FromTokenAmount")
	}

	expectedAmountOut := "1926426408"
	if swap.ToTokenAmount.String() != expectedAmountOut {
		t.Fatal("Expected ", expectedAmountOut, " got ", swap.ToTokenAmount, " for ToTokenAmount")
	}

	if err != nil {
		t.Fatal("Fatal error for OneInchAggregator: ", err)
	}
}

func TestOneInch_RunGetApprovalSpender(t *testing.T) {
	oneInch, err := aggregators.NewOneInchAggregator(oneInchName, oneInchBaseUrl, chainID)
	if err != nil {
		t.Fatal("Fatal error for OneInchAggregator: ", err)
	}

	_, err = oneInch.GetApprovalSpender(1)

	if err != nil {
		t.Fatal("Fatal error for OneInchAggregator: ", err)
	}
}
