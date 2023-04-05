package tokens

import (
	"github.com/dialecticch/medici-go/pkg/config"
	"github.com/spf13/cobra"
)

var (
	Cfg *config.Config

	chainID uint64
)

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	Use: "tokens",
}

func init() {
	Command.AddCommand(addCmd)
}
