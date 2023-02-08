package types

import "time"

type MicSec int64

// 微妙 10e-6秒 : 介于纳秒和毫秒之间
func NewMicSec(t ...time.Time) MicSec {
	if len(t) == 0 {
		return MicSec(time.Now().In(GetZone()).UnixNano() / 1e3)
	}
	return MicSec(t[0].UnixNano() / 1e3)
}

func (p MicSec) Tick() Tick {
	return Tick(p / 1e3)
}

func (p MicSec) Val() int64 {
	return int64(p)
}

func (p MicSec) MonthStartEnd() (MicSec, MicSec) {
	now := p.Tick().ZeroTimeOfDay()
	firstDay := now.AddDate(0, 0, -now.Day()+1)
	start := NewMicSec(firstDay)
	end := NewMicSec(firstDay.AddDate(0, 1, 0))
	return start, end
}

func (p MicSec) DayStartEnd() (MicSec, MicSec) {
	now := p.Tick().ZeroTimeOfDay()
	start := NewMicSec(now)
	end := NewMicSec(now.AddDate(0, 0, 1))
	return start, end
}
