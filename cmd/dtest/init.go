package main

import (
	"github.com/spf13/viper"
)

func init() {
	// Set Configuration Defaults
	viper.SetDefault("debug", true)
	viper.SetDefault("verbosity", false)
	viper.SetDefault("quiet", false)

	// Read in Viper Configuration
	viper.SetConfigName("dtest")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath("$HOME/.local/etc")
	viper.AddConfigPath("./tests")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cmd.PersistentFlags().BoolP(
		"verbose", "v",
		viper.GetBool("verbosity"),
		"Enables verbose logging messages")
	viper.BindPFlag("verbosity", cmd.PersistentFlags().Lookup("verbose"))

	cmd.PersistentFlags().BoolP(
		"debug", "D",
		viper.GetBool("debug"),
		"Enables debug logging messages")
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	cmd.PersistentFlags().BoolP(
		"quiet", "q",
		viper.GetBool("quiet"),
		"Quiets logging messages")
	viper.BindPFlag("quiet", cmd.PersistentFlags().Lookup("quiet"))

}
