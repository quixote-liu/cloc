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
	sort Optioner
}

func newFileCmd(path string) cmder {
	return &fileCmd{
		path: path,
		sort: newSortOption(),
	}
}

func (f *fileCmd) run(opts map[string]string) (code int, err error) {
	opts = f.sort.extract(opts)
	if len(opts) != 0 {
		return ExitCodeFailed, fmt.Errorf("the count of file does not support options: [%s]", serializeMap(opts))
	}

	ext := filepath.Ext(f.path)
	fn, miss := newFileNoter(ext)
	if miss {
		return ExitCodeFailed, fmt.Errorf("does not support the file(%s)", ext)
	}
	fp, err := f.extractPoints(fn)
	if err != nil {
		return ExitCodeFailed, err
	}

	res := fmt.Sprintf("the count of file %s:\n\n", f.path)
	v := f.sort.value()
	if v == "" || v == sortValueCode {
		res += fmt.Sprintf("codes -------- %d\n", fp.codes)
	}
	if v == "" || v == sortValueBlank {
		res += fmt.Sprintf("blanks -------- %d\n", fp.blanks)
	}
	if v == "" || v == sortValueComment {
		res += fmt.Sprintf("comments -------- %d\n", fp.comments)
	}
	fmt.Println(res)
	code = ExitCodeSuccess

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

	var mulNotesCount int
	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	for scanner.Scan() {
		line := scanner.Text()

		// if the coding file does not support notes
		if fj.notHaveNotes() {
			if fj.isBlank(line) {
				fp.blanks++
			} else {
				fp.codes++
			}
			continue
		}

		if mulNotesCount != 0 {
			if fj.tailWithMultilineNotes(line) {
				fp.comments += mulNotesCount + 1
				mulNotesCount = 0
				continue
			}
			mulNotesCount++
			continue
		}

		// judge the coding line whether is multiline notes
		if fj.beginWithMultilineNotes(line) {
			if fj.tailWithMultilineNotes(line) {
				fp.comments++
				continue
			}
			mulNotesCount++
			continue
		}

		// judeg the codeing line whether is blank
		if fj.isBlank(line) {
			fp.blanks++
			continue
		}

		// judge the coding line whether is single notes
		if fj.isSingleLineNotes(line) {
			fp.comments++
			continue
		}

		// the coding line is code
		fp.codes++
	}

	if err = scanner.Err(); err != nil {
		return
	}

	if mulNotesCount != 0 {
		err = fmt.Errorf("the comments using of file %s is error", f.path)
	}

	return
}

type fileJudger interface {
	// judge the coding line whether has notes
	notHaveNotes() bool
	beginWithMultilineNotes(line string) bool
	tailWithMultilineNotes(line string) bool
	isSingleLineNotes(line string) bool

	// judge the coding line whether is blank
	isBlank(line string) bool
}

type fileJudge struct {
	notHaveN            bool
	singleLineNotesChar string
	multilineNotesChars []string
}

func (fj *fileJudge) notHaveNotes() bool {
	return fj.notHaveN
}

func (fj *fileJudge) beginWithMultilineNotes(line string) bool {
	begin := fj.multilineNotesChars[0]
	return strings.HasPrefix(line, begin)
}

func (fj *fileJudge) tailWithMultilineNotes(line string) bool {
	tail := fj.multilineNotesChars[1]
	return strings.HasSuffix(line, tail)
}

func (fj *fileJudge) isSingleLineNotes(line string) bool {
	return strings.HasPrefix(line, fj.singleLineNotesChar)
}

func (fj *fileJudge) isBlank(line string) bool {
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

func newFileNoter(ext string) (fj fileJudger, miss bool) {
	ext = strings.TrimPrefix(ext, ".")
	switch {
	case contains(jsExtensions, ext):
		fj = &fileJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	case contains(jsonExtensions, ext):
		fj = &fileJudge{notHaveN: true}
	case contains(tsExtensions, ext):
		fj = &fileJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	case contains(htmlExtensions, ext):
		fj = &fileJudge{multilineNotesChars: []string{"<!--", "-->"}}
	case contains(scssExtensions, ext):
		fj = &fileJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	case contains(cssExtensions, ext):
		fj = &fileJudge{multilineNotesChars: []string{"/*", "*/"}}
	case contains(goExtensions, ext):
		fj = &fileJudge{singleLineNotesChar: "//", multilineNotesChars: []string{"/*", "*/"}}
	default:
		miss = true
	}
	return
}
