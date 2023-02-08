package types

import (
	"time"
)

// 毫秒 millisecond
type Tick int64

func (p Tick) Value() int64 {
	return int64(p)
}

func (p Tick) ZeroTimeOfDay() time.Time {
	layout := "20060102"
	s := p.ToTime().Format(layout)
	t, err := time.ParseInLocation(layout, s, GetZone())
	Chk(err)
	return t
}

func (p Tick) BeginOfDay() Tick {
	v := NewTick(p.ZeroTimeOfDay()).Value()
	return Tick(v)
}

func (p Tick) EndOfDay() Tick {
	t := p.ZeroTimeOfDay().
		Add(time.Hour * time.Duration(24)).
		Add(time.Nanosecond * time.Duration(-1))
	v := NewTick(t).Value()
	return Tick(v)
}

// ToTime convert tick to local time
func (p Tick) ToTime() time.Time {
	tick := int64(p)
	return time.Unix(tick/1e3, (tick%1e3)*1e6)
}

// NewTick create Tick. default value if now at local time
func NewTick(t ...time.Time) Tick {
	if len(t) == 0 {
		return Tick(time.Now().In(GetZone()).UnixNano() / 1e6)
	}

	return Tick(t[0].UnixNano() / 1e6)
}
