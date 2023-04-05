package control

import (
	"github.com/dialecticch/medici-go/pkg/config"
	"github.com/spf13/cobra"
)

var (
	Cfg     *config.Config
	ChainID uint64
)

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	Use:   "ctrl",
	Short: "Control command for starting and stopping medici services.",
}

func init() {
	Command.AddCommand(automateCmd)
	Command.AddCommand(harvestsCmd)
	Command.AddCommand(balancesCmd)
	Command.AddCommand(safeCmd)
}

func SetConfig(cfg *config.Config) {
	Cfg = cfg
}
