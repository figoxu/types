package types_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"

	"github.com/figoxu/types"
)

var _ = Describe("Day", func() {
	It("WeekDayPreRange", func() {
		day := types.NewDayByTick(types.NewTick().Value())
		start, end := day.WeekDayPreRange(time.Saturday)
		fmt.Println(start, " ", end)
		fmt.Println(start.AddDay(-7), " ", end.AddDay(-7))
	})
	It("GetWeekStartAndEndDay", func() {
		day := types.Day(20220306)
		start, end := day.GetWeekStartAndEndDay()
		fmt.Println(start, " ", end)
	})

	FIt("PreNDays", func() {
		day := types.Day(20220519)
		days := day.PreNDays(0)
		fmt.Println("pre 3 days:", days)
	})
	It("NextNDays", func() {
		day := types.Day(20220519)
		days := day.NextNDays(3)
		fmt.Println("pre 3 days:", days)
	})
})
