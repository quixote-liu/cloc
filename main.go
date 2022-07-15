package main

import (
	"os"
)

const (
	ExitCodeSuccess = iota
	ExitCodeFailed
)

func main() {
	_ = parseArguments(os.Args)

	os.Exit(ExitCodeFailed)
}
