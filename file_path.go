package types

import (
	"strings"
)

type FilePath string

func (p FilePath) Suffix(defaultVal ...string) string {
	var suffix string
	if len(defaultVal) > 0 {
		suffix = defaultVal[0]
	}
	fullPath := string(p)
	idx := strings.LastIndex(fullPath, ".")
	if idx == 0 {
		return suffix
	}
	if idx == -1 {
		return suffix
	}
	return fullPath[idx:]
}
