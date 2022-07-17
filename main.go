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
		fmt.Printf("[ERROR]: missing arguments, see 'cloc help'")
		os.Exit(ExitCodeFailed)
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
		fmt.Printf("[ERROR]: %v", err)
		os.Exit(ExitCodeFailed)
	}

	// parse directory and file
	if s.IsDir() {
		// parse directory
	} else {
		// parse single file
	}

	os.Exit(ExitCodeFailed)
}

const helptext = `if you want count file, like JavaScript file, you can input:

	cloc ./demo.js -sort code -order desc
	
	the cloc tool support these types of file:
	-----    JavaScript
	-----    JSON
	-----    TypeScript
	-----    HTML
	-----    SCSS

	if you want count directory, you can input:

	cloc ./dirdemo

	the directory census does not support options.
`

func transformRawOptions(raws []string) (map[string]string, error) {

}
