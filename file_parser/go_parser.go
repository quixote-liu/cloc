package fileparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/quixote-liu/cloc/option"
)

type goparser struct {
	filePath string
	options  *option.Options
}

func newGoParser(filePath string, options *option.Options) *goparser {
	return &goparser{
		filePath: filePath,
		options:  options,
	}
}

func (p *goparser) ParsePage() (PageParserResult, error) {
	return p.parsePage(p.filePath)
}

func (p *goparser) parsePage(filePath string) (PageParserResult, error) {
	parserResult := PageParserResult{}
	file, err := os.Open(p.filePath)
	if err != nil {
		return parserResult, fmt.Errorf("open the file failed, file path = %s, error = %v", p.filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	isMultilineComment := false
	for scanner.Scan() {
		line := cleanLine(scanner.Text())

		// multiline comments
		if strings.HasPrefix(line, "/*") {
			parserResult.CommentLines++
			isMultilineComment = true
			continue
		}
		if strings.HasPrefix(line, "*/") {
			parserResult.CommentLines++
			isMultilineComment = false
			continue
		}
		if isMultilineComment {
			parserResult.CommentLines++
			continue
		}

		// single comment
		if strings.HasPrefix(line, "//") {
			parserResult.CommentLines++
			continue
		}

		// blank line
		if line == "" {
			parserResult.BlankLines++
			continue
		}

		// code line
		parserResult.CodeLines++
	}
	if err := scanner.Err(); err != nil {
		return parserResult, fmt.Errorf("parse file failed: %v", err)
	}

	return parserResult, nil
}

func (p *goparser) ParseDir() error {
	// TODO: optimize
}
