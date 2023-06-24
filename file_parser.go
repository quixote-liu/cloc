package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type FileParser struct {
	parsers []parser
}

func NewFileParser() *FileParser {
	return &FileParser{
		parsers: []parser{}, // TODO: optimize
	}
}

type parser interface {
	Name() string
	Match(fileName string) bool
	Parse(file *os.File) parserResult
}

type parserResult struct {
	CommentLines int
	BlankLines   int
	CodeLines    int
	err          error
}

// parse golang file
type golangParser struct {
	name string
	ext  string
}

func newGolangParser() parser {
	return &golangParser{
		name: "golang",
		ext:  ".go",
	}
}

func (p *golangParser) Name() string {
	return p.name
}

func (p *golangParser) Match(fileName string) bool {
	return filepath.Ext(fileName) == p.ext
}

func (p *golangParser) Parse(file *os.File) parserResult {
	scaner := bufio.NewScanner(file)
	var comments, blanks, codes int
	var isLongCommnets bool
	var isInLongSentence bool
	for scaner.Scan() {
		text := scaner.Text()
		// if the line is blank
		if text == "" {
			blanks++
			continue
		}
		var isInQuoteSentence bool
		if index := strings.Index(text, "/*"); index != -1 {
			for i := index - 1; i >= 0; i-- {
				if text[i] != '"' {
					continue
				}
				if index-1 >= 0 && text[index-1] != '\\' {
					isInQuoteSentence = true
				}
			}
		}

		if strings.Contains(text, "//") {
			// 判断该
		}

	}
	if err := scaner.Err(); err != nil {
		return parserResult{err: err}
	}
	return parserResult{}
}
