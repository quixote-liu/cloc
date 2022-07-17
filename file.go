package main

import "fmt"

type fileCmd struct {
	path string
}

func newFileCmd(path string) cmder {
	return &fileCmd{
		path: path,
	}
}

func (f *fileCmd) run(opts map[string]string) (code int, err error) {
	fmt.Println("options:", opts)
	code = 0
	fmt.Println("in file running")
	return
}
