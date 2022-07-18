package main

import (
	"fmt"
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
	fmt.Println("options:", opts)
	code = 0
	fmt.Println("in file running")
	return
}

type filePoint struct {
	codes    int
	blanks   int
	comments int
}

func (f *fileCmd) extractPoints(fj fileJudger) filePoint {

}

type fileJudger interface {
	notHaveNote() bool
	matchBeginNote(line string) bool
	matchTailNote(line string) bool
	matchSingleNote(line string) bool
}

type multiLineNote [2]string

type fileJudge struct {
	notHaveN       bool
	singleLineNote string
	multiLineNote  multiLineNote
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
		fj = &fileJudge{singleLineNote: "//", multiLineNote: multiLineNote{"/*", "*/"}}
	case contains(jsonExtensions, ext):
		fj = &fileJudge{notHaveN: true}
	case contains(tsExtensions, ext):
		fj = &fileJudge{singleLineNote: "//", multiLineNote: multiLineNote{"/*", "*/"}}
	case contains(htmlExtensions, ext):
		fj = &fileJudge{multiLineNote: multiLineNote{"<!--", "-->"}}
	case contains(scssExtensions, ext):
		fj = &fileJudge{singleLineNote: "//", multiLineNote: multiLineNote{"/*", "*/"}}
	case contains(cssExtensions, ext):
		fj = &fileJudge{multiLineNote: multiLineNote{"/*", "*/"}}
	default:
		miss = true
	}
	return
}
