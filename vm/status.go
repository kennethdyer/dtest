package vm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

var vmLineSplit = regexp.MustCompile("  +")

// Status is run from the command-line and is used to print the
// status of the virtual machine in Podman configured for use in
// this project by DTest.
func Status() {
	name, vmLine, ok := findRelevantVM()
	if !ok {
		lgr.Errorf("DTest VM (%s) does not exist", name)
		return
	}
	line := vmLineSplit.Split(vmLine, -1)

	con := []string{}
	if viper.GetBool("verbosity") {
		con = []string{
			"Podman Virtual Machine Status (DTest)",
			fmt.Sprintf("Name: %s", line[0]),
			fmt.Sprintf("Hardware: %s (%s CPU / %s mem / %s size)", line[1], line[4], line[5], line[6]),
			fmt.Sprintf("Created: %s", line[2]),
			fmt.Sprintf("Last Active: %s", line[3]),
		}
	} else {
		active := line[3]
		switch active {
		case "Never":
			active = "never started"
		default:
			active = fmt.Sprintf("last up %s", active)
		}
		con = []string{
			fmt.Sprintf("%s (%s)", line[0], active),
		}
	}

	fmt.Println(strings.Join(con, "\n  "))
}
