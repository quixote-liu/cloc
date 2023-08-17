package option

import (
	"errors"
	"fmt"

	"github.com/quixote-liu/cloc/util"
)

type orderOption struct {
	isMatched bool
	names     []string

	isDesc bool
	isAsc  bool
}

func newOrderOption() *orderOption {
	return &orderOption{
		names: []string{"--order", "-or"},
	}
}

func (o *orderOption) extractArgs(args []string) ([]string, error) {
	match := util.StringsContains(o.names, args[0])
	if !match {
		return args, nil
	}

	if o.isMatched {
		return nil, errors.New("the option of [order] is duplication")
	}
	o.isMatched = true

	if len(args) <= 1 {
		return nil, errors.New("the option of [order] missing the value")
	}
	value := args[1]
	switch value {
	case "desc":
		o.isDesc = true
	case "asc":
		o.isAsc = true
	default:
		return nil, fmt.Errorf("the value(%s) of option [sort] is error", value)
	}

	return args[2:], nil
}

func (o *orderOption) Matched() bool {
	return o.isMatched
}

func (o *orderOption) IsDesc() bool {
	return o.isDesc
}

func (o *orderOption) IsAsc() bool {
	return o.isAsc
}
