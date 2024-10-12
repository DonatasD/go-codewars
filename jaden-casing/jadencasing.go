package jadencasing

import (
	"strings"
	"unicode"
)

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	var result []string

	for _, w := range words {
		rs := []rune(w)
		rs[0] = unicode.ToUpper(rs[0])
		result = append(result, string(rs))
	}
	return strings.Join(result, " ")
}

func ToJadenCaseV2(str string) string {
	return strings.Title(str)
}
