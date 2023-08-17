package option

import (
	"errors"
	"fmt"
	"os"

	"github.com/quixote-liu/cloc/util"
)

type outerOption struct {
	isMatched bool
	tags      []string

	path string
}

func newOuterOption() *outerOption {
	return &outerOption{
		tags: []string{"--outer", "-o"},
	}
}

func (o *outerOption) extractArgs(args []string) ([]string, error) {
	match := util.StringsContains(o.tags, args[0])
	if !match {
		return args, nil
	}

	if o.isMatched {
		return nil, errors.New("the option of [file] is duplication")
	}
	o.isMatched = true

	if len(args) <= 1 {
		return nil, errors.New("the option of [file] missing the value")
	}
	path := args[1]
	if _, err := os.Stat(path); os.IsExist(err) {
		return nil, fmt.Errorf("the path(%s) of file is exist, please re-input", path)
	}

	o.path = path

	return args[2:], nil
}

func (o *outerOption) Matched() bool {
	return o.isMatched
}

func (o *outerOption) FilePath() string {
	return o.path
}
