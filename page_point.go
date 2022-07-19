package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type pagePoint struct {
	codes    int
	blanks   int
	comments int
}

func newPagePoint() *pagePoint {
	return &pagePoint{}
}

func (p *pagePoint) extract(path string, pj pageJudger) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var mulNotesCount int
	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	for scanner.Scan() {
		line := scanner.Text()

		// if the coding file does not support notes
		if pj.notHaveNotes() {
			if pj.isBlank(line) {
				p.blanks++
			} else {
				p.codes++
			}
			continue
		}

		if mulNotesCount != 0 {
			if pj.tailWithMultilineNotes(line) {
				p.comments += mulNotesCount + 1
				mulNotesCount = 0
				continue
			}
			mulNotesCount++
			continue
		}

		// judge the coding line whether is multiline notes
		if pj.beginWithMultilineNotes(line) {
			if pj.tailWithMultilineNotes(line) {
				p.comments++
				continue
			}
			mulNotesCount++
			continue
		}

		// judeg the codeing line whether is blank
		if pj.isBlank(line) {
			p.blanks++
			continue
		}

		// judge the coding line whether is single notes
		if pj.isSingleLineNotes(line) {
			p.comments++
			continue
		}

		// the coding line is code
		p.codes++
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if mulNotesCount != 0 {
		return fmt.Errorf("the comments using of file %s is error", path)
	}

	return nil
}
