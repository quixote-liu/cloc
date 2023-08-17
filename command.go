package main

import (
	"fmt"
	"os"

	fileparser "github.com/quixote-liu/cloc/file_parser"
	"github.com/quixote-liu/cloc/option"
)

type Command struct {
	isOutputHelpText bool

	targetFilePath string
	fileParser     fileparser.FileParser
}

func NewCommand(args []string) (*Command, error) {
	var isOutputHelpText bool
	if len(args) == 1 || (len(args) == 2 && (args[1] == "-help" || args[1] == "-h")) {
		isOutputHelpText = true
	}
	if isOutputHelpText {
		return &Command{isOutputHelpText: true}, nil
	}

	options := option.New()
	var targetFilePath string

	args = args[1:]
	var err error
	for len(args) > 0 {
		originLength := len(args)

		if originLength == 1 {
			targetFilePath = args[0]
			if _, err = os.Stat(targetFilePath); err != nil {
				return nil, fmt.Errorf("get file information failed: %v", err)
			}
			break
		}

		args, err := options.ExtractArguments(args)
		if err != nil {
			return nil, err
		}
		if len(args) == originLength {
			return nil, fmt.Errorf("the command %s is undefined", args[0])
		}
	}

	if targetFilePath == "" {
		return nil, fmt.Errorf("missing the target file")
	}

	return &Command{
		targetFilePath: targetFilePath,
		fileParser:     fileparser.NewFileParser(targetFilePath, options),
	}, nil
}

func (c *Command) Run() error {
	if c.isOutputHelpText {
		fmt.Print(helptext)
		return nil
	}

	if !c.targetFilePath.IsDir() {
		// TODO: optimize parse single page and output
	} else {
		// TODO: optimize parse dir and output
	}

	return nil
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
