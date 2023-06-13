package main

import (
	"os"
	"path/filepath"
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
	Parse(file os.File) parserResult
}

type parserResult struct {
	CommentLines int
	BlankLines   int
	CodeLines    int
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

func (p *golangParser) Parse(file os.File) parserResult {
	// TODO: add codes
	return parserResult{}
}
