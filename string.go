package types

import "regexp"

type String string

func NewString(v string) String {
	return String(v)
}

func (p String) Val() string {
	return string(p)
}

func (p String) IsEmail() bool {
	matchFlag := p.IsMatchAnyRegexp(
		`^[\w]+(\.[\w]+)*@[\w]+(\.[\w])+$`,
		`[\w]+(\.[\w]+)*@[\w]+(\.[\w])+`,
	)
	return matchFlag
}

func (p String) IsMatchAnyRegexp(patterns ...string) bool {
	isMatch := func(pattern string) bool {
		reg := regexp.MustCompile(pattern)
		matchFlag := reg.MatchString(p.Val())
		return matchFlag
	}
	if len(patterns) == 0 {
		return false
	}
	for _, pattern := range patterns {
		matchFlag := isMatch(pattern)
		if matchFlag {
			return true
		}
	}
	return false
}
