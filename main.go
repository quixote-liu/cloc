package main

import (
	"os"
)

const (
	ExitCodeSuccess = iota
	ExitCodeFailed
)

func main() {
	args := os.Args

	cmd, err := NewCommand(args)
	if err != nil {
		printfErr(err)
		os.Exit(ExitCodeFailed)
	}

	if err = cmd.Run(); err != nil {
		printfErr(err)
		os.Exit(ExitCodeFailed)
	}

	os.Exit(ExitCodeSuccess)
}


