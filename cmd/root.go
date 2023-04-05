package cmd

import (
	"github.com/dialecticch/medici-go/cmd/admin"
	"github.com/dialecticch/medici-go/cmd/control"
	"github.com/spf13/cobra"

	"github.com/dialecticch/medici-go/pkg/config"
)

var cfgFile string
var cfg *config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "medici",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".medici.toml", "config file (default is $HOME/.medici.toml)")
	rootCmd.AddCommand(control.Command)
	rootCmd.AddCommand(admin.Command)
}

// initConfig reads in config file.
func initConfig() {
	cfg = &config.Config{}
	cobra.CheckErr(config.Load(cfgFile, cfg))

	control.SetConfig(cfg)
	admin.SetConfig(cfg)
}
