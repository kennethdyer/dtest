package main

import (
	"github.com/spf13/cobra"
)

var runCmd = cobra.Command{
	Use:   "run",
	Short: "Run configured or specified tests",
	Run: func(_ *cobra.Command, args []string) {
		logInit()
		unimpl("test")
	},
}
