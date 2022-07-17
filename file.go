package main

type file struct {}

func newFile(path string) cmder {
	return &file{}
}

func (f *file) run(opts map[string]string) (code int, err error) {
	return
}