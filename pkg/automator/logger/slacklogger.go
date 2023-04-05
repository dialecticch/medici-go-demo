package logger

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/slack-go/slack"
)

type SlackLogger struct {
	url      string
	explorer string
}

func NewSlackLogger(url, explorer string) *SlackLogger {
	return &SlackLogger{url: url, explorer: explorer}
}

func (s *SlackLogger) Failed(receipt *types.Receipt) error {
	return slack.PostWebhook(
		s.url,
		&slack.WebhookMessage{
			Text: fmt.Sprintf(
				"*Failed Transaction:* <%[1]s/tx/%[2]s|%[2]s>",
				s.explorer,
				strings.ToLower(receipt.TxHash.String()),
			),
		},
	)
}

func (s *SlackLogger) SendError(err error, address common.Address) error {
	return slack.PostWebhook(
		s.url,
		&slack.WebhookMessage{
			Text: fmt.Sprintf(
				"*Failed Send Transaction:* \n Error: %[1]s \n Sender: \n <%[2]s/address/%[3]s|%[3]s>",
				err,
				s.explorer,
				address.String(),
			),
		},
	)
}

func (s *SlackLogger) MaxRetriesReached(strategy common.Address, safe common.Address, pool *big.Int) error {
	return slack.PostWebhook(
		s.url,
		&slack.WebhookMessage{
			Text: fmt.Sprintf(
				"*Harvesting has reached max retries:* \n Strategy: <%[1]s/address/%[2]s|%[2]s> \n Safe: <%[1]s/address/%[2]s|%[2]s> \n Pool: %[3]s",
				s.explorer,
				strings.ToLower(strategy.String()),
				strings.ToLower(safe.String()),
				pool.String(),
			),
		},
	)
}
