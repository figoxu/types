package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// RmbStr money str like ¥112,121,89
type RmbStr string

func (p RmbStr) Val() string {
	return string(p)
}

func (p RmbStr) Float64() float64 {
	// rm all ',' in str
	str := strings.ReplaceAll(p.Val(), `,`, ``)

	if !strings.Contains(str, "¥") {
		return 0
	}

	v := strings.Split(str, "¥")[1]
	d, err := decimal.NewFromString(v)
	if err != nil {
		return 0
	}
	out, _ := d.Float64()
	return out
}

type IntStr string

func NewIntStr(val int64) IntStr {
	return IntStr(fmt.Sprintf("%v", val))
}

func (p IntStr) Val() string {
	return string(p)
}

func (p IntStr) MustToInt64() int64 {
	out, _ := strconv.ParseInt(string(p), 10, 64)
	return out
}

type IntStrLst []IntStr

func (p IntStrLst) MustToInt64Lst() []int64 {
	var out []int64
	for _, v := range p {
		out = append(out, v.MustToInt64())
	}
	return out
}
