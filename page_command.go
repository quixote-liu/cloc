package main

import (
	"fmt"
	"path/filepath"
)

type pageCmd struct {
	path string
	sort Optioner
}

func newPageCmd(path string) cmder {
	return &pageCmd{
		path: path,
		sort: newSortOption(),
	}
}

func (cmd *pageCmd) run(opts map[string]string) (code int, err error) {
	opts = cmd.sort.extract(opts)
	if len(opts) != 0 {
		return ExitCodeFailed, fmt.Errorf("the count of file does not support options: [%s]", serializeMap(opts))
	}

	ext := filepath.Ext(cmd.path)
	pj, miss := newPageJudger(ext)
	if miss {
		return ExitCodeFailed, fmt.Errorf("does not support the file(%s)", ext)
	}

	pp := newPagePoint()
	if err = pp.extract(cmd.path, pj); err != nil {
		return ExitCodeFailed, err
	}

	res := fmt.Sprintf("the count of file %s:\n\n", cmd.path)
	v := cmd.sort.value()
	if v == "" || v == sortValueCode {
		res += fmt.Sprintf("codes -------- %d\n", pp.codes)
	}
	if v == "" || v == sortValueBlank {
		res += fmt.Sprintf("blanks -------- %d\n", pp.blanks)
	}
	if v == "" || v == sortValueComment {
		res += fmt.Sprintf("comments -------- %d\n", pp.comments)
	}
	fmt.Println(res)
	code = ExitCodeSuccess

	return
}
