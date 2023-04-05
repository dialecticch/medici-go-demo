package main

import (
	"os"

	"github.com/dialecticch/medici-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
