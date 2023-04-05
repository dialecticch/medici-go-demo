package writers

import (
	"fmt"

	"github.com/dialecticch/medici-go/pkg/strings"
	"github.com/slack-go/slack"
)

type SlackWriter struct {
	url      string
	explorer string
}

func NewSlackWriter(url, explorer string) *SlackWriter {
	return &SlackWriter{
		url:      url,
		explorer: explorer,
	}
}

func (s *SlackWriter) Write(event *Event) error {
	data := strings.EventData{
		Title: "New Harvest",
		Fields: map[string]string{
			"Transaction": strings.Transaction(s.explorer, event.Transaction.String()),
			"Harvested":   fmt.Sprintf("%s of %s", event.Amount.String(), strings.Token(s.explorer, event.Token.String())),
			"Safe":        strings.Address(s.explorer, event.Safe.String()),
			"Pool":        event.Pool.String(),
		},
	}

	return slack.PostWebhook(s.url, &slack.WebhookMessage{Text: data.String()})
}
