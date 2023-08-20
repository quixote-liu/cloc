package fileparser

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/quixote-liu/cloc/option"
)

type FileParser interface {
}

func NewFileParser(filePath string, options *option.Options) FileParser {
	switch filepath.Ext(filePath) {
	case ".go":
		return newGoParser(filePath, options)
	default:
		return nil
	}
}

type PageParserResult struct {
	err          error
	CommentLines int
	BlankLines   int
	CodeLines    int
}

func cleanLine(line string) string {
	return strings.TrimSpace(line)
}

type output struct {
	outer io.ReadWriteCloser
}

func newOutput(options option.Options) *output {
	outer := os.Stdout
	path := options.OuterOption.FilePath()
	if path != "" {
		file, err := os.Open(path)
		if err == nil {
			outer = file
		}
	}
	return &output{outer: outer}
}

func (o *output) WriteResult(filePath string, result PageParserResult) {
	// TODO: optimize
}
