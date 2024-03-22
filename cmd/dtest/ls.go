package main

import (
	"github.com/spf13/cobra"
)

var lsCmd = cobra.Command{
	Use:   "ls",
	Short: "Lists configured images, deployment targets, and topologies",
	Run: func(_ *cobra.Command, args []string) {
		logInit()
		unimpl("ls")
	},
}

var imgCmd = cobra.Command{
	Use:   "images",
	Short: "Lists configured images",
	Run: func(_ *cobra.Command, args []string) {
		logInit()
		unimpl("ls images")
	},
}

var targCmd = cobra.Command{
	Use:   "targets",
	Short: "Lists configured deployment targets",
	Run: func(_ *cobra.Command, args []string) {
		logInit()
		unimpl("ls targets")
	},
}

var topCmd = cobra.Command{
	Use:   "topologies",
	Short: "Lists configured deployment topologies",
	Run: func(_ *cobra.Command, args []string) {
		logInit()
		unimpl("ls topologies")
	},
}
