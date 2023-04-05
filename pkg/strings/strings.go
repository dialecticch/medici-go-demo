package strings

import "fmt"

func Transaction(explorer, hash string) string {
	return fmt.Sprintf("<%[1]s/tx/%[2]s|%[2]s>", explorer, hash)
}

func Address(explorer, address string) string {
	return fmt.Sprintf("<%[1]s/address/%[2]s|%[2]s>", explorer, address)
}

func Token(explorer, address string) string {
	return fmt.Sprintf("<%[1]s/token/%[2]s|%[2]s>", explorer, address)
}

type EventData struct {
	Title  string
	Fields map[string]string
}

func (d *EventData) String() string {
	msg := "*" + d.Title + "*"

	for k, v := range d.Fields {
		msg += "\n    " + k + ": " + v
	}

	return msg
}
