package fileparser

import (
	"bufio"
	"fmt"
	"os"
)

type goparser struct {
	filePath string
}

func newGoParser(filePath string) *goparser {
	return &goparser{
		filePath: filePath,
	}
}

func (p *goparser) ParsePage() (PageParserResult, error) {
	parserResult := PageParserResult{}
	file, err := os.Open(p.filePath)
	if err != nil {
		return parserResult, fmt.Errorf("open the file failed, file path = %s, error = %v", p.filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// add logic
}
