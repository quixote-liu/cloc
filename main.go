package main

import (
	"fmt"
	"os"
)

const (
	ExitCodeSuccess = iota
	ExitCodeFailed
)

func main() {
	args := os.Args
	if len(args) == 1 {
		args = append(args, "help")
	}

	path := args[1]

	// if subcommand is 'help'
	if path == "help" {
		fmt.Printf("%v\n", helptext)
		os.Exit(ExitCodeSuccess)
	}

	// check the correctness of file path
	s, err := os.Stat(path)
	if err != nil {
		printfErr(err)
		os.Exit(ExitCodeFailed)
	}

	opts, code, err := parseRawOptions(args[2:])
	if err != nil {
		printfErr(err)
		os.Exit(code)
	}

	// get the instance of command
	var cmd cmder
	if s.IsDir() {
		cmd = newDirCmd(path)
	} else {
		cmd = newPageCmd(path)
	}

	// run command
	code, err = cmd.run(opts)
	if err != nil {
		printfErr(err)
	}

	os.Exit(code)
}

const helptext = `if you want count code file, like JavaScript file, you can input:

	cloc ./demo.js -sort code
	
	the cloc tool support these types of file:
	-----    JavaScript
	-----    JSON
	-----    TypeScript
	-----    HTML
	-----    SCSS
	-----    CSS
	-----    Golang

	if you want count directory, you can input:

	cloc ./dirdemo
`

type cmder interface {
	run(opts map[string]string) (code int, err error)
}
