package main

import "fmt"

type Dir struct {
	path string
}

func newDir(path string) cmder {
	return &Dir{
		path: path,
	}
}

func (d *Dir) run(opts map[string]string) (code int, err error) {
	fmt.Println("options:", opts)
	code = 0
	fmt.Println("in dir running")
	return
}
