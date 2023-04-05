package aggregators

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type OneInchToken struct {
	Symbol   string   `json:"symbol"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Decimals int64    `json:"decimals"`
	LogoURI  string   `json:"logoURI"`
	Tags     []string `json:"tags"`
}

type OneInchProtocol [][]struct {
	Name             string  `json:"name"`
	Part             float64 `json:"part"`
	FromTokenAddress string  `json:"fromTokenAddress"`
	ToTokenAddress   string  `json:"toTokenAddress"`
}

type OneInchTx struct {
	// transactions will be sent from this address
	From string `json:"from"`
	// transactions will be sent to our contract address
	To string `json:"to"`
	// call data
	Data string `json:"data"`
	// amount of ETH (in wei) will be sent to the contract address
	Value string `json:"value"`
	// gas price in wei
	GasPrice string `json:"gasPrice"`
	// estimated amount of the gas limit, increase this value by 25%
	Gas int64 `json:"gas"`
}

type OneInchSwapOpts struct {
	// protocols that can be used in a swap
	Protocols string
	// address that will receive a purchased token
	// Receiver of destination currency. default: fromAddress
	DestReceiver string
	// referrer's address
	ReferrerAddress string
	// referrer's fee in percentage
	// Ethereum: min: 0; max: 3; Binance: min: 0; max: 3; default: 0; !should be the same for quote and swap!
	Fee string
	// gas price
	// default: fast from network
	GasPrice string
	// if true, CHI will be burned from fromAddress to compensate gas
	// default: false; Suggest to check user's balance and allowance before set this flag; CHI should be approved to spender address
	BurnChi bool
	// how many connectorTokens can be used
	// min: 0; max: 3; default: 2; !should be the same for quote and swap!
	ComplexityLevel string
	// contract addresses of connector tokens
	// max: 5; !should be the same for quote and swap!
	ConnectorTokens string
	// if true, accept the partial order execution
	AllowPartialFill bool
	// if true, checks of the required quantities are disabled
	DisableEstimate bool
	// maximum amount of gas for a swap
	GasLimit string
	// virtual split parts. default: 50; max: 500; !should be the same for quote and swap!
	VirtualParts string
	// maximum number of parts each main route part can be split into
	// split parts. default: Ethereum: 50; Binance: 40 max: Ethereum: 100; Binance: 100; !should be the same for quote and swap!
	Parts string
	// maximum number of main route parts
	// default: Ethereum: 10, Binance: 10; max: Ethereum: 50, Binance: 50 !should be the same for quote and swap!
	MainRouteParts string
}

type OneInchSwap struct {
	// parameters of a token to sell
	FromToken OneInchToken `json:"fromToken"`
	// parameters of a token to buy
	ToToken OneInchToken `json:"ToToken"`
	// input amount of fromToken in minimal divisible units
	ToTokenAmount string `json:"toTokenAmount"`
	// result amount of toToken in minimal divisible units
	FromTokenAmount string `json:"fromTokenAmount"`
	// route of the trade
	Protocols []OneInchProtocol `json:"protocols"`
	// transaction data
	Tx OneInchTx `json:"tx"`
}

type OneInchQuoteOpts struct {
	// Ethereum: min: 0; max: 3; Binance: min: 0; max: 3; default: 0; !should be the same for quote and swap!
	Fee string
	// Liquidity protocols that can be used in a swap
	Protocols string
	// gas price
	// default: fast from network
	GasPrice string
	// how many connectorTokens can be used
	// min: 0; max: 3; default: 2; !should be the same for quote and swap!
	ComplexityLevel string
	// contract addresses of connector tokens
	// max: 5; !should be the same for quote and swap!
	ConnectorTokens string
	// maximum amount of gas for a swap
	GasLimit string
	// virtual split parts. default: 50; max: 500; !should be the same for quote and swap!
	VirtualParts string
	// maximum number of parts each main route part can be split into
	// split parts. default: Ethereum: 50; Binance: 40 max: Ethereum: 100; Binance: 100; !should be the same for quote and swap!
	Parts string
	// maximum number of main route parts
	// default: Ethereum: 10, Binance: 10; max: Ethereum: 50, Binance: 50 !should be the same for quote and swap!
	MainRouteParts string
}

type OneInchQuote struct {
	// parameters of a token to sell
	FromToken OneInchToken `json:"fromToken"`
	// parameters of a token to buy
	ToToken OneInchToken `json:"toToken"`
	// input amount of fromToken in minimal divisible units
	ToTokenAmount string `json:"toTokenAmount"`
	// result amount of toToken in minimal divisible units
	FromTokenAmount string `json:"fromTokenAmount"`
	// route of the trade
	Protocols []OneInchProtocol `json:"protocols"`
	// rough estimated amount of the gas limit for used protocols;
	// do not use estimatedGas from the quote method as the gas limit of a transaction
	EstimatedGas int64 `json:"estimatedGas"`
}

type OneInchApproveSpender struct {
	// address of 1inch contract
	Address string `json:"address"`
}

type OneInchAggregator struct {
	Name    string
	BaseUrl url.URL
	Client  *http.Client
}

func NewOneInchAggregator(name string, baseUrl string, chainID string) (*OneInchAggregator, error) {
	oneInchBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, errors.New("cannot parse base url")
	}

	oneInchBaseUrl.Path = path.Join(oneInchBaseUrl.Path, chainID)

	return &OneInchAggregator{
		Name:    name,
		BaseUrl: *oneInchBaseUrl,
		Client:  &http.Client{},
	}, nil
}

func (a OneInchAggregator) GetQuote(args QuoteArgs) (*Quote, error) {
	endpoint := "/quote"

	opts, err := getQuoteOpts(&args)
	if err != nil {
		return nil, err
	}

	var oneInchQuote OneInchQuote
	err = a.doRequest(args.ChainID, endpoint, "GET", nil, &oneInchQuote, opts)
	if err != nil {
		return nil, err
	}

	quote, err := convertQuote(oneInchQuote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func (a OneInchAggregator) GetSwap(args SwapArgs) (*Swap, error) {
	endpoint := "/swap"

	queries, err := getSwapOpts(&args)
	if err != nil {
		return nil, err
	}

	var oneInchSwap OneInchSwap
	err = a.doRequest(args.ChainID, endpoint, "GET", nil, &oneInchSwap, queries)
	if err != nil {
		return nil, err
	}

	swap, err := convertSwap(oneInchSwap)
	if err != nil {
		return nil, err
	}

	return swap, nil
}

func (a OneInchAggregator) GetApprovalSpender(chainID uint64) (*common.Address, error) {
	endpoint := "/approve/spender"

	var oneInchApproveSpender OneInchApproveSpender

	err := a.doRequest(chainID, endpoint, "GET", &oneInchApproveSpender, nil, nil)
	if err != nil {
		return nil, err
	}

	approveSpender := common.HexToAddress(oneInchApproveSpender.Address)

	return &approveSpender, nil
}

func (a OneInchAggregator) doRequest(chainId uint64, endpoint string, method string, reqData interface{}, expRes interface{}, queries url.Values) error {
	baseUrl := a.BaseUrl
	baseUrl.Path = path.Join(baseUrl.Path, endpoint)
	baseUrl.RawQuery = filterQuery(&queries).Encode()

	var reqBytes []byte
	var err error

	if reqData != nil {
		reqBytes, err = json.Marshal(reqData)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		if expRes == nil {
			return nil
		}

		err = json.Unmarshal(body, &expRes)

		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("returned %d: %s", resp.StatusCode, body)
	}
}

func filterQuery(query *url.Values) url.Values {
	filtered := url.Values{}
	for k, v := range *query {
		if len(v) > 0 && len(v[0]) > 0 {
			filtered[k] = v
		}
	}

	return filtered
}

func getQuoteOpts(args *QuoteArgs) (url.Values, error) {
	if args.From.String() == "" || args.To.String() == "" || args.Amount.String() == "" {
		return nil, errors.New("OneInchAggregator: missing arguments")
	}

	opts := args.Opts.(OneInchQuoteOpts)

	return url.Values{
		"amount":           {args.Amount.String()},
		"toTokenAddress":   {args.To.String()},
		"fromTokenAddress": {args.From.String()},
		"fee":              {opts.Fee},
		"protocols":        {opts.Protocols},
		"gasPrice":         {opts.GasPrice},
		"complexityLevel":  {opts.ComplexityLevel},
		"connectorTokens":  {opts.ConnectorTokens},
		"gasLimit":         {opts.GasLimit},
		"mainRouteParts":   {opts.MainRouteParts},
		"parts":            {opts.Parts},
	}, nil

}

func getSwapOpts(args *SwapArgs) (url.Values, error) {
	if args.From.String() == "" || args.To.String() == "" || args.Amount.String() == "" || args.FromAddress.String() == "" || args.Slippage.String() == "" {
		return nil, errors.New("OneInchAggregator: missing arguments")
	}

	opts := args.Opts.(OneInchSwapOpts)

	return url.Values{
		"amount":           {args.Amount.String()},
		"slippage":         {args.Slippage.String()},
		"toTokenAddress":   {args.To.String()},
		"fromTokenAddress": {args.From.String()},
		"fromAddress":      {args.FromAddress.String()},
		"burnChi":          {strconv.FormatBool(opts.BurnChi)},
		"allowPartialFill": {strconv.FormatBool(opts.AllowPartialFill)},
		"disableEstimate":  {strconv.FormatBool(opts.DisableEstimate)},
		"protocols":        {opts.Protocols},
		"destReceiver":     {opts.DestReceiver},
		"referrerAddress":  {opts.ReferrerAddress},
		"fee":              {opts.Fee},
		"gasPrice":         {opts.GasPrice},
		"complexityLevel":  {opts.ComplexityLevel},
		"connectorTokens":  {opts.ConnectorTokens},
		"gasLimit":         {opts.GasLimit},
		"parts":            {opts.Parts},
		"virtualParts":     {opts.VirtualParts},
		"mainRouteParts":   {opts.MainRouteParts},
	}, nil
}

func convertQuote(oneInchQuote OneInchQuote) (*Quote, error) {
	fromTokenAmount, success := new(big.Int).SetString(oneInchQuote.FromTokenAmount, 10)
	if !success {
		return nil, errors.New("cannot convert FromTokenAmount")
	}

	toTokenAmount, success := new(big.Int).SetString(oneInchQuote.ToTokenAmount, 10)
	if !success {
		return nil, errors.New("cannot convert ToTokenAmount")
	}

	quote := &Quote{
		FromToken:       common.HexToAddress(oneInchQuote.FromToken.Address),
		ToToken:         common.HexToAddress(oneInchQuote.ToToken.Address),
		FromTokenAmount: fromTokenAmount,
		ToTokenAmount:   toTokenAmount,
		EstimatedGas:    oneInchQuote.EstimatedGas,
	}

	return quote, nil
}

func convertSwap(oneInchSwap OneInchSwap) (*Swap, error) {
	fromTokenAmount, success := new(big.Int).SetString(oneInchSwap.FromTokenAmount, 10)
	if !success {
		return nil, errors.New("cannot convert FromTokenAmount")
	}

	toTokenAmount, success := new(big.Int).SetString(oneInchSwap.ToTokenAmount, 10)
	if !success {
		return nil, errors.New("cannot convert ToTokenAmount")
	}

	data := oneInchSwap.Tx.Data
	if len(data) > 2 && data[0:2] == "0x" {
		data = data[2:]
	}

	swap := &Swap{
		FromToken:       common.HexToAddress(oneInchSwap.FromToken.Address),
		ToToken:         common.HexToAddress(oneInchSwap.ToToken.Address),
		ToTokenAmount:   toTokenAmount,
		FromTokenAmount: fromTokenAmount,
		Tx: &AggregatorTransaction{
			From:     common.HexToAddress(oneInchSwap.Tx.From),
			To:       common.HexToAddress(oneInchSwap.Tx.To),
			Data:     common.Hex2Bytes(data),
			Value:    new(big.Int),
			GasPrice: new(big.Int),
			Gas:      oneInchSwap.Tx.Gas,
		},
	}

	return swap, nil
}
