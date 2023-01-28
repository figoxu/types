package types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Day int

func NewDayByTick(tick int64) Day {
	dayString := Tick(tick).ToTime().Format(`20060102`)
	day, err := strconv.Atoi(dayString)
	Chk(err)
	return Day(day)
}

func (p Day) Fmt(layouts ...string) (string, error) {
	layout := "2006-01-02"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	t, err := time.Parse("20060102", fmt.Sprint(p))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return t.Format(layout), nil
}

func (p Day) Value() int {
	return int(p)
}

func (p Day) IsZero() bool {
	return p.Value() == 0
}

func (p Day) TickRange() (Tick, Tick) {
	const dateFmt = "20060102"
	if len(fmt.Sprint(p)) != len(dateFmt) {
		return 0, 0
	}
	t, err := time.ParseInLocation(dateFmt, fmt.Sprint(p), time.Local)
	Chk(err)
	tick := Tick(NewTick(t).Value())
	return tick.BeginOfDay(), tick.EndOfDay()
}

// 当前日期所属的7天周期范围(往回看)
func (p Day) WeekDayPreRange(divideDays ...time.Weekday) (Day, Day) {
	startTick, _ := p.TickRange()
	todayZeroDay := startTick.ZeroTimeOfDay()
	divideDay := time.Sunday
	if len(divideDays) > 0 {
		divideDay = divideDays[0]
	}
	todayWeekDay := todayZeroDay.Weekday()
	dayOffset := divideDay - todayWeekDay
	startOfWeek := startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(dayOffset*24))
	if startOfWeek.Unix() > time.Now().Unix() {
		startOfWeek = startOfWeek.Add(time.Hour * time.Duration(-7*24))
	}
	endOfWeek := startOfWeek.Add(time.Hour * time.Duration(7*24))
	return NewDayByTick(NewTick(startOfWeek).Value()), NewDayByTick(NewTick(endOfWeek).Value())
}

func (p Day) AddDay(day int) Day {
	startTick, _ := p.TickRange()
	t := startTick.ZeroTimeOfDay().AddDate(0, 0, day)
	return NewDayByTick(NewTick(t).Value())
}

func (p Day) Year() int {
	y, err := p.Fmt("2006")
	Chk(err)
	year, err := strconv.Atoi(y)
	Chk(err)
	return year
}

func (p Day) Mon() int {
	m, err := p.Fmt("01")
	Chk(err)
	mon, err := strconv.Atoi(m)
	Chk(err)
	return mon
}

func Today() Day {
	return NewDayByTick(NewTick().Value())
}

func (p Day) GetWeekStartAndEndDay() (Day, Day) {
	startTick, _ := p.TickRange()
	weekDay := startTick.ZeroTimeOfDay().Weekday()
	var startOfWeek time.Time
	switch weekDay {
	case time.Monday:
		startOfWeek = startTick.ZeroTimeOfDay()
	case time.Tuesday:
		startOfWeek = startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(-1*24))
	case time.Wednesday:
		startOfWeek = startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(-2*24))
	case time.Thursday:
		startOfWeek = startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(-3*24))
	case time.Friday:
		startOfWeek = startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(-4*24))
	case time.Saturday:
		startOfWeek = startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(-5*24))
	case time.Sunday:
		startOfWeek = startTick.ZeroTimeOfDay().Add(time.Hour * time.Duration(-6*24))
	}
	endOfWeek := startOfWeek.Add(time.Hour * time.Duration(6*24))
	return NewDayByTick(NewTick(startOfWeek).Value()), NewDayByTick(NewTick(endOfWeek).Value())
}

func (p Day) PreNDays(n int) []Day {
	var ret []Day
	for i := 0; i < n; i++ {
		ret = append(ret, p.AddDay(-i))
	}
	return ret
}

func (p Day) NextNDays(n int) []Day {
	var ret []Day
	for i := 0; i < n; i++ {
		ret = append(ret, p.AddDay(i))
	}
	return ret
}
