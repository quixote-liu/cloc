package main

type Dir struct {
	path string
}

func newDir(path string) cmder {
	return &Dir{
		path: path,
	}
}

func (d *Dir) run(opts map[string]string) (code int, err error) {
	return
}
