package writers

import (
	"github.com/dialecticch/medici-go/pkg/strings"
	"github.com/slack-go/slack"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		Title: "New " + cases.Title(language.English).String(string(event.UpdateType)),
		Fields: map[string]string{
			"Transaction": strings.Transaction(s.explorer, event.Transaction.String()),
			"Amount":      event.Amount.String(),
			"Safe":        strings.Address(s.explorer, event.Safe.String()),
			"Pool":        event.Pool.String(),
		},
	}

	return slack.PostWebhook(s.url, &slack.WebhookMessage{Text: data.String()})
}
