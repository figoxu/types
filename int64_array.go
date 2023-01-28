package types

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ahmetb/go-linq/v3"
)

type Int64Array []int64

func NewInt64Array(v []int64) Int64Array {
	return Int64Array(v)
}

func (p Int64Array) Val() []int64 {
	return p
}

func (p Int64Array) First() int64 {
	if len(p) == 0 {
		return 0
	}
	return p[0]
}

func (p Int64Array) Append(vs ...int64) Int64Array {
	var out []int64
	out = append(out, p...)
	out = append(out, vs...)
	return out
}

func (p Int64Array) Skip(skipVs ...int64) Int64Array {
	var l []int64
	linq.From(p).WhereT(func(x int64) bool {
		skipFlag := linq.From(skipVs).Contains(x)
		return !skipFlag
	}).ToSlice(&l)
	return l
}

func (p Int64Array) Join(sep string) string {
	var vs []string
	linq.From(p).SelectT(func(x int64) string {
		return fmt.Sprint(x)
	}).ToSlice(&vs)
	return strings.Join(vs, sep)
}

func (p Int64Array) Sort(descFlag bool) Int64Array {
	var out []int64
	linq.From(p).OrderByT(func(x int64) int64 {
		if descFlag {
			return x * -1
		}
		return x
	}).ToSlice(&out)
	return out
}

func (p Int64Array) In(vs ...int64) Int64Array {
	var out []int64
	linq.From(p).WhereT(func(x int64) bool {
		return linq.From(vs).Contains(x)
	}).ToSlice(&out)
	return out
}

func (p Int64Array) Distinct() Int64Array {
	var out []int64
	linq.From(p).Distinct().ToSlice(&out)
	return out
}

func (p Int64Array) Contain(v int64) bool {
	return linq.From(p).Contains(v)
}

func (p Int64Array) Size() int {
	return len(p)
}

func (p Int64Array) Max() int64 {
	if len(p) == 0 {
		return 0
	}
	v := linq.From(p).Max()
	return v.(int64)
}

func (p Int64Array) WherePage(page *Page) Int64Array {
	var out []int64
	linq.From(p).Skip(page.Offset()).
		Take(page.GetLimit()).ToSlice(&out)
	return out
}

func (p Int64Array) TotalPageNo(page *Page) int {
	lastPageNo := (p.Size() + page.GetLimit() - 1) / page.GetLimit()
	return lastPageNo
}

func (p Int64Array) DeepEqual(in Int64Array) bool {
	return reflect.DeepEqual(p, in)
}
