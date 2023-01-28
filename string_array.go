package types

import (
	"reflect"
	"strings"

	"github.com/ahmetb/go-linq/v3"
)

type StringArray []string

func NewStringArray(v ...string) StringArray {
	return StringArray(v)
}

func (p StringArray) Contain(v string) bool {
	return linq.From(p).Contains(v)
}

func (p StringArray) ContainAny(vs ...string) bool {
	for _, v := range vs {
		containFlag := p.Contain(v)
		if containFlag {
			return true
		}
	}
	return false
}

func (p StringArray) Size() int {
	return len(p)
}

func (p StringArray) Skip(skipVs ...string) StringArray {
	var l []string
	linq.From(p).WhereT(func(x string) bool {
		skipFlag := linq.From(skipVs).Contains(x)
		return !skipFlag
	}).ToSlice(&l)
	return l
}

func (p StringArray) Join(separate string) string {
	return strings.Join(p, separate)
}

func (p StringArray) WherePage(page *Page) StringArray {
	var out []string
	linq.From(p).Skip(page.Offset()).
		Take(page.GetLimit()).ToSlice(&out)
	return out
}

func (p StringArray) TotalPageNo(page *Page) int {
	lastPageNo := (p.Size() + page.GetLimit() - 1) / page.GetLimit()
	return lastPageNo
}

func (p StringArray) DeepEqual(in StringArray) bool {
	return reflect.DeepEqual(p, in)
}
