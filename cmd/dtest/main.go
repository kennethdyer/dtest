package main

import (
	"github.com/kennethdyer/dtest"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "dtest",
	Short:   "The Deployment Testing Tool",
	Version: dtest.Version,
}

func main() {

	// Add vm Subcommand
	vm.AddCommand(vmList)
	vm.AddCommand(vmStart)
	vm.AddCommand(vmStatus)
	vm.AddCommand(vmStop)
	vm.AddCommand(vmInit)
	vm.AddCommand(vmDestroy)
	cmd.AddCommand(vm)

	// Execute Command
	cmd.Execute()
}
