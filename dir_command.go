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
	sortOpt  Optioner

	files         int64
	codesTotal    int64
	blanksTotal   int64
	commentsTotal int64
}

func newDirCmd(path string) cmder {
	return &dirCmd{
		path:     path,
		orderOpt: newOrderOption(),
		sortOpt:  newSortOption(),
	}
}

func (cmd *dirCmd) run(opts map[string]string) (int, error) {
	opts = cmd.orderOpt.extract(opts)
	opts = cmd.sortOpt.extract(opts)
	if len(opts) != 0 {
		return ExitCodeFailed, fmt.Errorf("the count of directory does not support options: [%s]", serializeMap(opts))
	}
	cmd.readFileNames(cmd.path, "")
	fmt.Println()
	switch cmd.sortOpt.value() {
	case sortValueFiles:
		fmt.Printf("the all files number: %d\n", cmd.files)
	case sortValueBlank:
		fmt.Printf("[blanks total]: %d\n", cmd.blanksTotal)
	case sortValueCode:
		fmt.Printf("[codes total]: %d\n", cmd.codesTotal)
	case sortValueComment:
		fmt.Printf("[comments total]: %d\n", cmd.commentsTotal)
	default:
		fmt.Printf("[total]: codes: %d, comments: %d, blanks: %d\n\n", cmd.codesTotal,
			cmd.commentsTotal, cmd.blanksTotal)
	}
	return ExitCodeSuccess, nil
}

func (cmd *dirCmd) readFileNames(path, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		printfErr(fmt.Errorf("failed to read directory: %v", err))
		os.Exit(ExitCodeFailed)
	}
	entries = cmd.sortEntries(entries)
	sortValue := cmd.sortOpt.value()
	prefix += "│----"
	for _, e := range entries {
		// deal directory
		p := filepath.Join(path, e.Name())
		if e.IsDir() {
			fmt.Println(prefix + e.Name())
			cmd.readFileNames(p, prefix)
			continue
		}

		// increment files number
		cmd.files++
		if sortValue == sortValueFiles {
			fmt.Println(prefix + e.Name())
			continue
		}

		// read page points
		ext := filepath.Ext(p)
		pj, miss := newPageJudger(ext)
		if miss {
			fmt.Println(prefix + e.Name())
			continue
		}
		pp := newPagePoint()
		if err := pp.extract(p, pj); err != nil {
			fmt.Println(prefix + e.Name() + " [ERROR: read failed]")
			continue
		}

		// total directory page messages
		cmd.codesTotal += int64(pp.codes)
		cmd.commentsTotal += int64(pp.comments)
		cmd.blanksTotal += int64(pp.blanks)

		// print the result
		var tail string
		switch sortValue {
		case sortValueCode:
			tail = fmt.Sprintf("[codes: %d]", pp.codes)
		case sortValueBlank:
			tail = fmt.Sprintf("[blanks: %d]", pp.blanks)
		case sortValueComment:
			tail = fmt.Sprintf("[comments: %d]", pp.comments)
		default:
			tail = fmt.Sprintf("[codes: %d, comments: %d, blanks: %d]", pp.codes, pp.comments, pp.blanks)
		}
		fmt.Println(prefix + e.Name() + " ---------->" + tail)
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
