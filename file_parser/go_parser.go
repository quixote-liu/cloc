package fileparser

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"sort"
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

}

func (p *goparser) parseDir(dir string) error {
	dirfs := os.DirFS(p.filePath)

	fileNames, err := fs.Glob(dirfs, "*")
	if err != nil {
		return fmt.Errorf("get the files from %s failed: %v", p.filePath, err)
	}
	if len(fileNames) == 0 {
		return nil
	}

	// set sort
	if p.options.OrderOption.IsAsc() {
		sort.Sort(sort.StringSlice(fileNames))
	} else if p.options.OrderOption.IsDesc() {
		sort.Sort(sort.Reverse(sort.StringSlice(fileNames)))
	}

	// parse file
	for _, fileName := range fileNames {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			return fmt.Errorf("read status of %s failed: error = %v", fileName, err)
		}
		if fileInfo.IsDir() {
			if err := p.parseDir(fileInfo.Name()); err != nil {
				return err
			}
		} else {
			
		}
	}
}
