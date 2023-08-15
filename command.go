package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Command struct {
	isOutputHelpText bool

	sortOpt  *sortOption
	orderOpt *orderOption
	fileOpt  *fileOption

	outer io.WriteCloser

	targetFileInfo os.FileInfo
}

func NewCommand(args []string) (*Command, error) {
	var isOutputHelpText bool
	if len(args) == 1 || (len(args) == 2 && (args[1] == "-help" || args[1] == "-h")) {
		isOutputHelpText = true
	}
	if isOutputHelpText {
		return &Command{isOutputHelpText: true}, nil
	}

	sortOpt := newSortOption()
	orderOpt := newOrderOption()
	fileOpt := newFileOption()
	var fileInfo os.FileInfo

	args = args[1:]
	var err error
	for len(args) > 0 {
		originLength := len(args)

		args, err = sortOpt.extract(args)
		if err != nil {
			return nil, err
		}

		args, err = orderOpt.extract(args)
		if err != nil {
			return nil, err
		}

		args, err = fileOpt.extract(args)
		if err != nil {
			return nil, err
		}

		if len(args) == 1 {
			fileInfo, err = os.Stat(args[0])
			if err != nil {
				return nil, fmt.Errorf("get file information failed: %v", err)
			}
		}

		if len(args) == originLength {
			return nil, fmt.Errorf("the command %s is undefined", args[0])
		}
	}

	var outer io.WriteCloser
	if fileOpt.isMatched {
		file, err := os.Create(fileOpt.path)
		if err != nil {
			return nil, fmt.Errorf("output file %s error: %v", fileOpt.path, err)
		}
		outer = file
	} else {
		outer = os.Stdout
	}

	return &Command{
		sortOpt:        sortOpt,
		orderOpt:       orderOpt,
		fileOpt:        fileOpt,
		outer:          outer,
		targetFileInfo: fileInfo,
	}, nil
}

func (c *Command) Run() error {
	if c.isOutputHelpText {
		fmt.Print(helptext)
		return nil
	}

	if !c.targetFileInfo.IsDir() {
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

type sortOption struct {
	isMatched bool
	names     []string

	isCode    bool
	isComment bool
	isBlank   bool
}

func newSortOption() *sortOption {
	return &sortOption{
		names: []string{"-sort", "-s"},
	}
}

func (s *sortOption) extract(args []string) ([]string, error) {
	mat := stringsContains(s.names, args[0])
	if !mat {
		return args, nil
	}
	if s.isMatched {
		return nil, errors.New("the option of [sort] is duplication")
	}
	s.isMatched = true

	if len(args) <= 1 {
		return nil, errors.New("the option of [sort] missing the value")
	}
	value := args[1]
	switch value {
	case "code":
		s.isCode = true
	case "comment":
		s.isComment = true
	case "blank":
		s.isBlank = true
	default:
		return nil, fmt.Errorf("the value(%s) of option [sort] is error", value)
	}

	return args[2:], nil
}

type orderOption struct {
	isMatched bool
	names     []string

	isDesc bool
	isAsc  bool
}

func newOrderOption() *orderOption {
	return &orderOption{
		names: []string{"-order", "-o"},
	}
}

func (o *orderOption) extract(args []string) ([]string, error) {
	mat := stringsContains(o.names, args[0])
	if !mat {
		return args, nil
	}
	if o.isMatched {
		return nil, errors.New("the option of [order] is duplication")
	}
	o.isMatched = true

	if len(args) <= 1 {
		return nil, errors.New("the option of [order] missing the value")
	}
	value := args[1]
	switch value {
	case "desc":
		o.isDesc = true
	case "asc":
		o.isAsc = true
	default:
		return nil, fmt.Errorf("the value(%s) of option [sort] is error", value)
	}

	return args[2:], nil
}

type fileOption struct {
	isMatched bool
	names     []string

	path string
}

func newFileOption() *fileOption {
	return &fileOption{
		names: []string{"-file", "-f"},
	}
}

func (f *fileOption) extract(args []string) ([]string, error) {
	mat := stringsContains(f.names, args[0])
	if !mat {
		return args, nil
	}
	if f.isMatched {
		return nil, errors.New("the option of [file] is duplication")
	}
	f.isMatched = true

	if len(args) <= 1 {
		return nil, errors.New("the option of [file] missing the value")
	}
	path := args[1]
	if _, err := os.Stat(path); os.IsExist(err) {
		return nil, fmt.Errorf("the path(%s) of file is exist, please re-input", path)
	}

	f.path = path

	return args[2:], nil
}
