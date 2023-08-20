package option

import (
	"errors"
	"fmt"
)

type sortOption struct {
	isMatched bool
	names     []string

	isCode    bool
	isComment bool
	isBlank   bool
}

func newSortOption() *sortOption {
	return &sortOption{
		names: []string{"--sort", "-s"},
	}
}

func (o *sortOption) extractArgs(args []string) ([]string, error) {
	match := stringsContains(o.names, args[0])
	if !match {
		return args, nil
	}

	if o.isMatched {
		return nil, errors.New("the option of [sort] is duplication")
	}
	o.isMatched = true

	if len(args) <= 1 {
		return nil, errors.New("the option of [sort] missing the value")
	}
	value := args[1]
	switch value {
	case "code":
		o.isCode = true
	case "comment":
		o.isComment = true
	case "blank":
		o.isBlank = true
	default:
		return nil, fmt.Errorf("the value(%s) of option [sort] is error", value)
	}

	return args[2:], nil
}

func (o *sortOption) IsCode() bool {
	return o.isCode || (!o.IsComment() && !o.IsBlank())
}

func (o *sortOption) IsComment() bool {
	return o.isComment
}

func (o *sortOption) IsBlank() bool {
	return o.isBlank
}
