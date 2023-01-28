package types

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Int64Array", func() {
	It("WeekDayPreRange", func() {
		var vs Int64Array
		var i int64
		for i = 100; i <= 168; i++ {
			vs = append(vs, i)
		}
		page := &Page{
			Limit: 10,
			Page:  1,
		}
		fmt.Println("PageNo: ", vs.TotalPageNo(page))
		pageResult := vs.WherePage(page)
		for pageResult.Size() > 0 {
			fmt.Println(JsonString(pageResult))
			page.Next()
			pageResult = vs.WherePage(page)
		}
		fmt.Println("Done")
	})
})
