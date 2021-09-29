package pauthor

import (
	"regexp"
)

const RegexStr = `^([^<(]+?)?[ \t]*(?:<([^>(]+?)>)?[ \t]*(?:\(([^)]+?)\)|$)`

var _regex = regexp.MustCompile(RegexStr)

func Regex() *regexp.Regexp {
	return _regex
}
