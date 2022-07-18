package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

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

func (d *dirCmd) run(opts map[string]string) (int, error) {
	opts = d.orderOpt.extract(opts)
	if len(opts) != 0 {
		return ExitCodeFailed, fmt.Errorf("the count of directory does not support options: [%s]", serializeMap(opts))
	}
	d.readFileNames(d.path, "")
	return ExitCodeSuccess, nil
}

func (d *dirCmd) readFileNames(path, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("[ERROR]: failed to read file: %v\n", err)
		os.Exit(ExitCodeFailed)
	}
	entries = d.sortEntries(entries)
	prefix += "│----"
	for _, e := range entries {
		fmt.Println(prefix + e.Name())
		if e.IsDir() {
			p := filepath.Join(path, e.Name())
			d.readFileNames(p, prefix)
			continue
		}
	}
}

func (d *dirCmd) sortEntries(values []os.DirEntry) []os.DirEntry {
	var less func(i, j int) bool
	switch d.orderOpt.value() {
	case orderValueDesc:
		less = func(i, j int) bool {
			return values[i].Name() > values[j].Name()
		}
	case orderValueAsc:
		less = func(i, j int) bool {
			return values[i].Name() < values[j].Name()
		}
	default:
		return values
	}
	sort.SliceStable(values, less)
	return values
}
