package types_test

import (
	"fmt"
	"time"

	"github.com/figoxu/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Minute", func() {
	It("RangeTo", func() {
		from := types.Minute(202109171050)
		to := types.Minute(202109171056)
		vs := from.RangeTo(to)
		fmt.Println(vs)
		fmt.Println(to.AsDay())

		from = types.Minute(202109171100)
		to = types.Minute(202109171205)
		vs = from.RangeTo(to, time.Minute*time.Duration(5))
		fmt.Println(vs)
	})
	It("WeekStartAndEndDay", func() {
		day := types.Day(20220306)
		start, end := day.GetWeekStartAndEndDay()
		fmt.Println(start, " ", end)
	})
	It("Tick", func() {
		cases := []struct {
			Minute types.Minute
			Expect types.Tick
		}{
			{
				types.Minute(202109171050),
				types.Tick(1631847000000),
			},
			{
				types.Minute(202203031330),
				types.Tick(1646285400000),
			},
		}
		for _, t := range cases {
			tick := t.Minute.Tick()
			Ω(tick).To(Equal(t.Expect))
		}
	})

	It("Format", func() {
		cases := []struct {
			Minute types.Minute
			Layout string
			Expect string
		}{
			{
				types.Minute(202109171050),
				"2006/01/02/15:04",
				"2021/09/17/10:50",
			},
			{
				types.Minute(202203031330),
				"15:04",
				"13:30",
			},
			{
				types.Minute(0),
				"15:04",
				"08:00",
			},
		}

		for _, t := range cases {
			Ω(t.Minute.Format(t.Layout)).To(Equal(t.Expect))
		}
	})
})
