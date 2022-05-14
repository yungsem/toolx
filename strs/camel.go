package strs

import (
	"github.com/yungsem/toolx/pattern"
	"strings"
)

func SnakeToCamel(str string) (camelCase string) {
	str = strings.ToLower(str)

	if !strings.Contains(str, pattern.Underscore) {
		return UpperCaseFirstLetter(str)
	}

	b := Builder()
	arr := strings.Split(str, pattern.Underscore)
	for _, s := range arr {
		b.Append(UpperCaseFirstLetter(s))
	}

	return b.String()
}
