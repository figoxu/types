package types

import (
	"fmt"
	"strconv"
	"time"
)

type Minute int64

// 分钟 200601021504
func NewMinute(ts ...time.Time) Minute {
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	if t.IsZero() {
		return Minute(0)
	}
	v := t.Format(`200601021504`)
	minute, err := strconv.Atoi(v)
	Chk(err)
	if minute == 197001010800 {
		return Minute(0)
	}
	return Minute(minute)
}

func (p Minute) Val() int64 {
	return int64(p)
}

func (p Minute) AsDay() Day {
	return NewDayByTick(p.Tick().Value())
}

func (p Minute) Tick() Tick {
	layout := `200601021504`
	if p == 0 || len(fmt.Sprint(p)) < len(layout) {
		return 0
	}
	t, err := time.ParseInLocation(layout, fmt.Sprint(p), time.Local)
	Chk(err)
	return NewTick(t)
}

func (p Minute) RangeTo(toMin Minute, gaps ...time.Duration) []Minute {
	t := p.Tick().ToTime()
	vs := []Minute{p}
	gap := time.Minute
	if len(gaps) > 0 {
		gap = gaps[0]
		if gap <= time.Minute {
			gap = time.Minute
		}
	}
	next := t.Add(gap)
	fmt.Println("next val : ", NewMinute(next).Val())
	fmt.Println("to val : ", toMin.Val())
	for NewMinute(next).Val() <= toMin.Val() {
		vs = append(vs, NewMinute(next))
		next = next.Add(gap)
	}
	return vs
}

func (p Minute) LastMinute() Minute {
	return NewMinute(p.Tick().ToTime().Add(time.Minute * time.Duration(-1)))
}

func (p Minute) NextMinute() Minute {
	return NewMinute(p.Tick().ToTime().Add(time.Minute * time.Duration(1)))
}

func (p Minute) Add(n int) Minute {
	return NewMinute(p.Tick().ToTime().Add(time.Minute * time.Duration(n)))
}

func (p Minute) Format(layout string) string {
	return p.Tick().ToTime().Format(layout)
}
