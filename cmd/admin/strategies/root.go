package strategies

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
	Use: "strategies",
}

func init() {
	Command.AddCommand(addCmd)
	Command.AddCommand(listCmd)
}
