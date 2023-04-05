package writers

import (
	"fmt"

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
	t, err := s.text(event)
	if err != nil {
		return err
	}

	return slack.PostWebhook(
		s.url,
		&slack.WebhookMessage{
			Text: t,
		},
	)
}

func (s *SlackWriter) text(event *Event) (string, error) {
	switch event.UpdateType {
	case EnabledModule:
		return fmt.Sprintf(
			"*Module Enabled*\n        Safe: <%[1]s/address/%[2]s|%[2]s>\n        Module: <%[1]s/address/%[3]s|%[3]s>",
			s.explorer,
			event.Safe.String(),
			event.Arguments["module"],
		), nil
	case DisabledModule:
		return fmt.Sprintf(
			"*Module Disabled*\n        Safe: <%[1]s/address/%[2]s|%[2]s>\n        Module: <%[1]s/address/%[3]s|%[3]s>",
			s.explorer,
			event.Safe.String(),
			event.Arguments["module"],
		), nil
	case AddedOwner:
		return fmt.Sprintf(
			"*Added Owner*\n        Safe: <%[1]s/address/%[2]s|%[2]s>\n        Owner: <%[1]s/address/%[3]s|%[3]s>",
			s.explorer,
			event.Safe.String(),
			event.Arguments["owner"],
		), nil
	case RemovedOwner:
		return fmt.Sprintf(
			"*Removed Owner*\n        Safe: <%[1]s/address/%[2]s|%[2]s>\n        Owner: <%[1]s/address/%[3]s|%[3]s>",
			s.explorer,
			event.Safe.String(),
			event.Arguments["owner"],
		), nil
	case ChangedThreshold:
		return fmt.Sprintf(
			"*Changed Threshold*\n        Safe: <%[1]s/address/%[2]s|%[2]s>\n        Threshold: %[3]s",
			s.explorer,
			event.Safe.String(),
			event.Arguments["threshold"],
		), nil
	default:
		return "", fmt.Errorf("unknown event type: %d", event.UpdateType)
	}
}
