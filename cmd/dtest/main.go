package main

import (
	// Vendor Imports
	"github.com/spf13/cobra"

	// Local Imports
	"github.com/kennethdyer/dtest"
)

var cmd = cobra.Command{
	Use:     "dtest",
	Short:   "Deployment Testing Tool",
	Long:    "A Deployment Testing Tool for those who Detest Fouled Installation Instructions",
	Version: dtest.Version,
}

func main() {

	// dtest ls command and subcommands
	lsCmd.AddCommand(&topCmd)
	lsCmd.AddCommand(&targCmd)
	lsCmd.AddCommand(&imgCmd)
	cmd.AddCommand(&lsCmd)

	// dtest run command
	cmd.AddCommand(&runCmd)
	cmd.Execute()
}
