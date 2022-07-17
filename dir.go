package main

import "fmt"

type dirCmd struct {
	path     string
	orderOpt Optioner
}

func newDirCmd(path string) cmder {
	return &dirCmd{
		path:     path,
		orderOpt: newOrderOption(),
	}
}

func (d *dirCmd) run(opts map[string]string) (code int, err error) {
	if len(opts) == 0 {
		// read directory messages

	}

	opts = d.orderOpt.extract(opts)
	if len(opts) != 0 {
		return ExitCodeFailed, fmt.Errorf("the count of directory does not support options: [%s]", serializeMap(opts))
	}

	return
}

func (d *dirCmd) readFileNames() error {
	
	return nil
}
