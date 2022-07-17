package main

import "fmt"

type file struct{}

func newFile(path string) cmder {
	return &file{}
}

func (f *file) run(opts map[string]string) (code int, err error) {
	fmt.Println("options:", opts)
	code = 0
	fmt.Println("in file running")
	return
}
