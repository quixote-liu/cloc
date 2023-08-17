package fileparser

import (
	"bufio"
	"fmt"
	"os"

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
	parserResult := PageParserResult{}
	file, err := os.Open(p.filePath)
	if err != nil {
		return parserResult, fmt.Errorf("open the file failed, file path = %s, error = %v", p.filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// add logic
}
