package config

import (
	"github.com/spf13/viper"
	"github.com/useloopso/transaction-relayer/types"
)

var c types.Config

func Load() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}

func Get() types.Config {
	return c
}
