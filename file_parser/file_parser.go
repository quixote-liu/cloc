package fileparser

import (
	"path/filepath"

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
