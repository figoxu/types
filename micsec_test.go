package types_test

import (
	"fmt"
	"time"

	"github.com/figoxu/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("MicSec", func() {
	It("Default", func() {
		for i := 0; i < 100; i++ {
			fmt.Println(types.NewMicSec())
			fmt.Println(types.NewTick())
			time.Sleep(time.Duration(3) * time.Nanosecond)
		}
	})
	It("GetWeekStartAndEndDay", func() {
		day := types.Day(20220306)
		start, end := day.GetWeekStartAndEndDay()
		fmt.Println(start, " ", end)
	})
	It("MonthStartEnd", func() {
		day := types.MicSec(1648086757450774)
		start, end := day.MonthStartEnd()
		fmt.Println(start, " ", end)
	})
	It("DayStartEnd", func() {
		day := types.MicSec(1648086757450774)
		start, end := day.DayStartEnd()
		fmt.Println(start, " ", end)
	})
})
