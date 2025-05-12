package main

import (
	"github.com/kennethdyer/dtest"
	"github.com/spf13/cobra"
)

var cmdImg = &cobra.Command{
	Use:   "img",
	Short: "Administers DTest images for use in deployment testing",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunImage()
	},
}

var imgLs = &cobra.Command{
	Use:   "ls",
	Short: "Lists images configured for use by DTest",
	Run: func(_ *cobra.Command, _ []string) {
		dtest.RunImageLs()
	},
}
