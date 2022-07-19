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

func (cmd *dirCmd) run(opts map[string]string) (int, error) {
	opts = cmd.orderOpt.extract(opts)
	if len(opts) != 0 {
		return ExitCodeFailed, fmt.Errorf("the count of directory does not support options: [%s]", serializeMap(opts))
	}
	cmd.readFileNames(cmd.path, "")
	return ExitCodeSuccess, nil
}

func (cmd *dirCmd) readFileNames(path, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		printfErr(fmt.Errorf("failed to read directory: %v", err))
		os.Exit(ExitCodeFailed)
	}
	entries = cmd.sortEntries(entries)
	prefix += "│----"
	for _, e := range entries {
		fmt.Println(prefix + e.Name())
		if e.IsDir() {
			p := filepath.Join(path, e.Name())
			cmd.readFileNames(p, prefix)
			continue
		}
	}
}

func (cmd *dirCmd) sortEntries(values []os.DirEntry) []os.DirEntry {
	var less func(i, j int) bool
	switch cmd.orderOpt.value() {
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
