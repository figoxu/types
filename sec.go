package types

import (
	"fmt"
	"strconv"
	"time"
)

type Sec int64

func NewSec(ts ...time.Time) Sec {
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}

	v := t.Format(`20060102150405`)
	minute, err := strconv.Atoi(v)
	Chk(err)
	return Sec(minute)
}

func (p Sec) Val() int64 {
	return int64(p)
}

func (p Sec) Tick() Tick {
	layout := `20060102150405`
	t, err := time.ParseInLocation(layout, fmt.Sprint(p), time.Local)
	Chk(err)
	return NewTick(t)
}
