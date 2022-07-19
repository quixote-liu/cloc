package main

import "strings"

type pageJudger interface {
	// judge the page line whether has notes
	notHaveNotes() bool
	beginWithMultilineNotes(line string) bool
	tailWithMultilineNotes(line string) bool
	isSingleLineNotes(line string) bool

	// judge the page line whether is blank
	isBlank(line string) bool
}

type pageJudge struct {
	notHaveN            bool
	singleLineNotesChar string
	multilineNotesChars []string
}

func (pj *pageJudge) notHaveNotes() bool {
	return pj.notHaveN
}

func (pj *pageJudge) beginWithMultilineNotes(line string) bool {
	begin := pj.multilineNotesChars[0]
	return strings.HasPrefix(strings.TrimSpace(line), begin)
}

func (pj *pageJudge) tailWithMultilineNotes(line string) bool {
	tail := pj.multilineNotesChars[1]
	return strings.HasSuffix(strings.TrimSpace(line), tail)
}

func (pj *pageJudge) isSingleLineNotes(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), pj.singleLineNotesChar)
}

func (pj *pageJudge) isBlank(line string) bool {
	return line == ""
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

func newPageJudger(ext string) (pj pageJudger, miss bool) {
	ext = strings.TrimPrefix(ext, ".")
	switch {
	case contains(jsExtensions, ext):
		pj = &pageJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	case contains(jsonExtensions, ext):
		pj = &pageJudge{notHaveN: true}
	case contains(tsExtensions, ext):
		pj = &pageJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	case contains(htmlExtensions, ext):
		pj = &pageJudge{multilineNotesChars: []string{"<!--", "-->"}}
	case contains(scssExtensions, ext):
		pj = &pageJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	case contains(cssExtensions, ext):
		pj = &pageJudge{multilineNotesChars: []string{"/*", "*/"}}
	case contains(goExtensions, ext):
		pj = &pageJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	default:
		miss = true
	}
	return
}
