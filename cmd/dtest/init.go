package main

import "github.com/spf13/viper"

func init() {
	// Configure Defaults
	viper.SetDefault("dtest.machine.name", "mongodb-dtest")

}
