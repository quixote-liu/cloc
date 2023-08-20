package main

import (
	"fmt"
	"os"
)

const (
	ExitCodeSuccess = iota
	ExitCodeFailed
)

func printfErr(err error) {
	fmt.Printf("[ERROR]: %v\n", err)
}

func main() {
	cmd, err := NewCommand(os.Args)
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
