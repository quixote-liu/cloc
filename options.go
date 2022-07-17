package main

import (
	"fmt"
	"strings"
)

func parseRawOptions(raws []string) (map[string]string, int, error) {
	rr := strings.Join(raws, ":")
	opts := strings.Split(rr, "-")
	if opts[0] == "" {
		opts = opts[1:]
	}
	ans := make(map[string]string, len(opts))
	for _, opt := range opts {
		oo := strings.Split(strings.TrimSuffix(opt, ":"), ":")
		if len(oo) != 2 {
			return nil, ExitCodeFailed, fmt.Errorf("the option is error: %s", strings.Join(oo, " "))
		}
		var ok bool
		for _, o := range allOptions {
			if o.match(oo[0], oo[1]) {
				ans[oo[0]] = oo[1]
				ok = true
			}
		}
		if !ok {
			return nil, ExitCodeFailed, fmt.Errorf("the option is error: %s", strings.Join(oo, " "))
		}
	}
	return ans, ExitCodeSuccess, nil
}

var allOptions = []Option{newOrderOption(), newSortOption()}

type Option interface {
	match(tag, value string) bool
}

const (
	orderTag       = "order"
	orderValueDesc = "desc"
	orderValueAsc  = "asc"
)

type orderOption struct {
	tag    string
	values []string
}

func newOrderOption() *orderOption {
	return &orderOption{
		tag:    orderTag,
		values: []string{orderValueAsc, orderValueDesc},
	}
}

func (o *orderOption) match(tag, value string) bool {
	if o.tag != tag {
		return false
	}
	if !contains(o.values, value) {
		return false
	}
	return true
}

const (
	sortTag          = "sort"
	sortValueCode    = "code"
	sortValueFiles   = "files"
	sortValueBlank   = "blank"
	sortValueComment = "comment"
)

type sortOption struct {
	tag    string
	values []string
}

func newSortOption() *sortOption {
	return &sortOption{
		tag:    sortTag,
		values: []string{sortValueCode, sortValueFiles, sortValueBlank, sortValueComment},
	}
}

func (o *sortOption) match(tag, value string) bool {
	if o.tag != tag {
		return false
	}
	if !contains(o.values, value) {
		return false
	}
	return true
}
