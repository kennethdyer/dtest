package main

import (
	"github.com/charmbracelet/log"
)

func unimpl(c string) {
	log.Fatalf("Unimplemented Command: %s", c)
}
