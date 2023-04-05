package fees

import (
	"bytes"
	"math"
	"math/big"

	"github.com/dialecticch/medici-go/pkg/oracles"
	"github.com/ethereum-optimism/optimism/gas-oracle/bindings"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type Calculator interface {
	Fee(tx *types.Transaction) (*big.Float, error)
}

type EthereumFeeCalculator struct {
	oracle oracles.PriceOracle
}

func NewEthereumFeeCalculator(oracle oracles.PriceOracle) *EthereumFeeCalculator {
	return &EthereumFeeCalculator{oracle: oracle}
}

func (c *EthereumFeeCalculator) Fee(tx *types.Transaction) (*big.Float, error) {
	gas := quo(tx.Cost(), params.Ether)
	price := quo(c.oracle.Price(), math.Pow(10, float64(c.oracle.Decimals())))

	return new(big.Float).Mul(
		gas,
		price,
	), nil
}

type OptimismFeeCalculator struct {
	oracle    oracles.PriceOracle
	gasOracle *bindings.GasPriceOracle
}

func NewOptimismFeeCalculator(oracle oracles.PriceOracle, gasOracle *bindings.GasPriceOracle) *OptimismFeeCalculator {
	return &OptimismFeeCalculator{oracle: oracle, gasOracle: gasOracle}
}

func (c *OptimismFeeCalculator) Fee(tx *types.Transaction) (*big.Float, error) {
	// see https://community.optimism.io/docs/developers/build/transaction-fees/#displaying-fees-to-users for context
	l2gas := quo(tx.Cost(), params.Ether)

	var buf bytes.Buffer
	err := tx.EncodeRLP(&buf)
	if err != nil {
		return nil, err
	}

	l1cost, err := c.gasOracle.GetL1Fee(nil, buf.Bytes())
	if err != nil {
		return nil, err
	}

	gas := new(big.Float).Add(l2gas, new(big.Float).Mul(quo(l1cost, params.Ether), new(big.Float).SetFloat64(1.25)))
	price := quo(c.oracle.Price(), math.Pow(10, float64(c.oracle.Decimals())))

	return new(big.Float).Mul(
		gas,
		price,
	), nil
}

func quo(i *big.Int, f float64) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(i), big.NewFloat(f))
}
