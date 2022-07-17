package main

type sortOption string

const sortName sortOption = "-sort"

const (
	sortValueDesc = "desc"
	sortValueAsc  = "asc"
)

func (p *sortOption) is(opt string) bool {
	return opt == string(*p)
}

func (p *sortOption) matchValue(v string) {

}
