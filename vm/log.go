package vm

import (
	"os"

	"github.com/charmbracelet/log"
)

var lgr = log.NewWithOptions(
	os.Stderr,
	log.Options{
		ReportTimestamp: false,
		Level:           log.DebugLevel,
		Prefix:          "VM",
	},
)
