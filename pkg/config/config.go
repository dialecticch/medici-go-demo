package config

import (
	"os"

	"github.com/spf13/viper"
)

type PostgresConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSL      string `mapstructure:"ssl"`
}

type SlackConfig struct {
	Webhook string `mapstructure:"webhook"`
}

type Network struct {
	Node               string   `mapstructure:"node"`
	SenderRPC          string   `mapstructure:"sender_rpc"`
	Symbol             string   `mapstructure:"symbol"`
	Stablecoin         string   `mapstructure:"stablecoin"`
	StablecoinDecimals int      `mapstructure:"stablecoin_decimals"`
	WrappedToken       string   `mapstructure:"wrapped_token"`
	NoSellTokens       []string `mapstructure:"no_sell_tokens"`
	Explorer           string   `mapstructure:"explorer"`
	Epoch              uint64   `mapstructure:"epoch"`
	Aggregator         string   `mapstructur:"aggregator"`
}

type AggregatorConfig struct {
	Name    string `mapstructure:"name"`
	BaseUrl string `mapstructure:"base_url"`
}

type Config struct {
	PrivateKey  string                      `mapstructure:"private_key"`
	DB          PostgresConf                `mapstructure:"db"`
	Slack       SlackConfig                 `mapstructure:"slack"`
	Networks    map[uint64]Network          `mapstructure:"network"`
	Aggregators map[string]AggregatorConfig `mapstructure:"aggregator"`
}

// Load opens and parses a configuration file.
func Load(file string, conf interface{}) error {
	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	viper.SetConfigFile(file)
	viper.SetConfigType("toml")
	viper.SetEnvPrefix("medici")
	err = viper.BindEnv("private_key")
	if err != nil {
		return err
	}

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.GetViper().Unmarshal(conf)
	if err != nil {
		return err
	}

	return nil
}
