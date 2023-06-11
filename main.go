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

	cmd := NewCommand(args)

	err := cmd.Run()
	if err != nil {
		printfErr(err)
		os.Exit(ExitCodeFailed)
	}

	os.Exit(ExitCodeSuccess)
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
	-----    rust
	-----    C#
	-----    java
	-----    C/C++

	if you want count directory, you can input:

	cloc ./dirdemo
`

type cmder interface {
	run(opts map[string]string) (code int, err error)
}
