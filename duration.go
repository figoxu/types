package types

import "fmt"

type Duration int64 // 秒

func (p Duration) HMS() string {
	hour := p / (60 * 60)
	minute := (p % (60 * 60)) / 60
	sec := p % 60
	var result string
	if hour > 0 {
		result = result + fmt.Sprintf("%d小时", hour)
	}
	if minute > 0 {
		result = result + fmt.Sprintf("%d分", minute)
	}
	if sec > 0 {
		result = result + fmt.Sprintf("%d秒", sec)
	}
	return result
}
