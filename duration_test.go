package types_test

import (
	"fmt"

	"github.com/figoxu/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Duration", func() {
	It("HMS", func() {
		v := 5*60*60 + 4*60 + 3
		fmt.Println(types.Duration(v).HMS())
	})
	It("GetWeekStartAndEndDay", func() {
		day := types.Day(20220306)
		start, end := day.GetWeekStartAndEndDay()
		fmt.Println(start, " ", end)
	})
})
