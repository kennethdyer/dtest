package main

import (
	"github.com/kennethdyer/dtest"
	"github.com/spf13/cobra"
)

var vm = &cobra.Command{
	Use:   "vm",
	Short: "Wrapper interface to test and administer the Podman VM",
	Run: func(_ *cobra.Command, args []string) {
		dtest.RunVM(args)
	},
}

var vmList = &cobra.Command{
	Use:   "ls",
	Short: "List Podman virtual machines that are relevant to DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunVMList()
	},
}
