package main

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func logInit() {
	debug := viper.GetBool("debug")
	verbose := viper.GetBool("verbosity")
	quiet := viper.GetBool("quiet")
	log.SetReportTimestamp(false)
	log.SetPrefix("DTest")
	if quiet {
		log.SetLevel(log.ErrorLevel)
	} else if debug && verbose {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	} else if debug {
		log.SetLevel(log.DebugLevel)
	} else if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
