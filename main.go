package main

import (
	"os"

	"github.com/quixote-liu/cloc/util"
)

const (
	ExitCodeSuccess = iota
	ExitCodeFailed
)

func main() {
	cmd, err := NewCommand(os.Args)
	if err != nil {
		util.PrintfErr(err)
		os.Exit(ExitCodeFailed)
	}

	if err = cmd.Run(); err != nil {
		util.PrintfErr(err)
		os.Exit(ExitCodeFailed)
	}

	os.Exit(ExitCodeSuccess)
}
