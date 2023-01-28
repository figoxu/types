package types

import "github.com/ahmetb/go-linq/v3"

type Float64Array []float64

func NewFloat64Array(v ...float64) Float64Array {
	return Float64Array(v)
}

func (p Float64Array) Contain(v float64) bool {
	return linq.From(p).Contains(v)
}

func (p Float64Array) Size() int {
	return len(p)
}

func (p Float64Array) First() float64 {
	if len(p) == 0 {
		return 0
	}
	return p[0]
}

func (p Float64Array) Sum() float64 {
	var out float64 = 0
	for _, v := range p {
		out = out + v
	}
	return out
}
