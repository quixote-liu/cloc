package fileparser

import "path/filepath"

type FileParser interface {
}

func NewFileParser(filePath string) FileParser {
	switch filepath.Ext(filePath) {
	case ".go":
		return newGoParser(filePath)
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
