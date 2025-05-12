package main

import (
	"github.com/kennethdyer/dtest"
	"github.com/spf13/cobra"
)

var vm = &cobra.Command{
	Use:   "vm",
	Short: "Wrapper interface to test and administer the Podman VM used by DTest",
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

var vmStart = &cobra.Command{
	Use:   "start",
	Short: "Explicitly starts the Podman virtual machine for use by DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunVMStart()
	},
}

var vmStatus = &cobra.Command{
	Use:   "status",
	Short: "Reports the status of the Podman virtual machine used by DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunVMStatus()
	},
}

var vmStop = &cobra.Command{
	Use:   "stop",
	Short: "Explicitly stops the Podman virtual machine used by DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunVMStop()
	},
}

var vmInit = &cobra.Command{
	Use:   "init",
	Short: "Initializes a Podman virtual machine for use by DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunVMInit()
	},
}

var vmDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys the Podman virtual machine used by DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunVMDestroy()
	},
}
