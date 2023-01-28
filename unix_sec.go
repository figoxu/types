package types

import (
	"strconv"
	"time"
)

type UnixSec int64

func NewUnixSec(ts ...time.Time) UnixSec {
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return UnixSec(NewTick(t) / 1000)
}

func (p UnixSec) Tick() Tick {
	return Tick(p * 1000)
}

func (p UnixSec) Val() int64 {
	return int64(p)
}

func (p UnixSec) ToString() string {
	return strconv.FormatInt(int64(p), 10)
}
