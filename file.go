package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fileCmd struct {
	path string
}

func newFileCmd(path string) cmder {
	return &fileCmd{
		path: path,
	}
}

func (f *fileCmd) run(opts map[string]string) (code int, err error) {
	ext := filepath.Ext(f.path)
	fn, miss := newFileNoter(ext)
	if miss {
		return ExitCodeFailed, fmt.Errorf("does not support the file(%s) with extension '%s'", f.path, ext)
	}
	fp, err := f.extractPoints(fn)
	if err != nil {
		return ExitCodeFailed, err
	}

	// TODO: print result

	return
}

type filePoint struct {
	codes    int
	blanks   int
	comments int
}

func (f *fileCmd) extractPoints(fj fileJudger) (fp filePoint, err error) {
	content, err := os.ReadFile(f.path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	for scanner.Scan() {
		// TODO: optimize logic to extract file points.
	}

	return
}

type fileJudger interface {
	notHaveNote() bool
	matchBeginNote(line string) bool
	matchTailNote(line string) bool
	matchSingleNote(line string) bool
}

type fileJudge struct {
	notHaveN       bool
	singleLineNote string
	multiLineNote  []string
}

func (fj *fileJudge) notHaveNote() bool {
	return fj.notHaveN
}

func (fj *fileJudge) matchBeginNote(line string) bool {
	begin := fj.multiLineNote[0]
	return strings.HasPrefix(line, begin)
}

func (fj *fileJudge) matchTailNote(line string) bool {
	tail := fj.multiLineNote[1]
	return strings.HasSuffix(line, tail)
}

func (fj *fileJudge) matchSingleNote(line string) bool {
	return strings.HasPrefix(line, fj.singleLineNote)
}

var (
	jsExtensions   = []string{"js", "jsx", "mjs", "cjs"}
	jsonExtensions = []string{"json"}
	tsExtensions   = []string{"tsx", "ts"}
	htmlExtensions = []string{"html", "htm"}
	scssExtensions = []string{"scss"}
	cssExtensions  = []string{"css"}
	goExtensions   = []string{"go"}
)

func newFileNoter(ext string) (fj fileJudger, miss bool) {
	ext = strings.TrimPrefix(ext, ".")
	switch {
	case contains(jsExtensions, ext):
		fj = &fileJudge{singleLineNote: "//", multiLineNote: []string{"/*", "*/"}}
	case contains(jsonExtensions, ext):
		fj = &fileJudge{notHaveN: true}
	case contains(tsExtensions, ext):
		fj = &fileJudge{singleLineNote: "//", multiLineNote: []string{"/*", "*/"}}
	case contains(htmlExtensions, ext):
		fj = &fileJudge{multiLineNote: []string{"<!--", "-->"}}
	case contains(scssExtensions, ext):
		fj = &fileJudge{singleLineNote: "//", multiLineNote: []string{"/*", "*/"}}
	case contains(cssExtensions, ext):
		fj = &fileJudge{multiLineNote: []string{"/*", "*/"}}
	case contains(goExtensions, ext):
		fj = &fileJudge{singleLineNote: "//", multiLineNote: []string{"/*", "*/"}}
	default:
		miss = true
	}
	return
}
