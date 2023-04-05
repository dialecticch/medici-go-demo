package admin

import (
	"github.com/dialecticch/medici-go/cmd/admin/pools"
	"github.com/dialecticch/medici-go/cmd/admin/strategies"
	"github.com/dialecticch/medici-go/cmd/admin/tokens"
	"github.com/dialecticch/medici-go/pkg/config"
	"github.com/spf13/cobra"
)

var (
	Cfg *config.Config
)

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	Use:   "admin",
	Short: "Server side admin cli",
}

func init() {
	Command.AddCommand(strategies.Command)
	Command.AddCommand(pools.Command)
	Command.AddCommand(tokens.Command)
	Command.AddCommand(backfillCmd)
	Command.AddCommand(cleanerCmd)
}

func SetConfig(cfg *config.Config) {
	Cfg = cfg
	pools.Cfg = cfg
	strategies.Cfg = cfg
	tokens.Cfg = cfg
}
