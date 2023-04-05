package oracles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"sync"
	"time"
)

// PriceOracle returns the price of an asset as a big int.
type PriceOracle interface {
	Decimals() int
	Price() *big.Int
}

type CoinmarketCapPriceOracle struct {
	mux sync.RWMutex

	price *big.Int

	from string
	url  string
}

func NewCoinmarketCapPriceOracle(from, to string) (*CoinmarketCapPriceOracle, error) {
	oracle := &CoinmarketCapPriceOracle{
		from: from,
		url:  fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s", from, to),
	}

	price, err := oracle.query()
	if err != nil {
		return nil, err
	}

	oracle.mux.Lock()
	oracle.price = price
	oracle.mux.Unlock()

	go oracle.run()

	return oracle, nil
}

func (c *CoinmarketCapPriceOracle) Price() *big.Int {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return c.price
}

func (c *CoinmarketCapPriceOracle) Decimals() int {
	return 6
}

func (c *CoinmarketCapPriceOracle) run() {
	ticker := time.NewTicker(10 * time.Minute)

	for range ticker.C {
		price, err := c.query()
		if err != nil {
			log.Printf("failed to query price err: %s", err)
			continue
		}

		c.mux.Lock()
		c.price = price
		c.mux.Unlock()

		log.Printf("New %s price %s", c.from, price.String())
	}
}

func (c *CoinmarketCapPriceOracle) query() (*big.Int, error) {
	response, err := http.Get(c.url)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	price := &struct {
		USD float64 `json:"USD"`
	}{}

	err = json.Unmarshal(responseData, price)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetUint64(uint64(price.USD * math.Pow(10, float64(c.Decimals())))), nil
}
