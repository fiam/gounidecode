package unidecode

import (
	"unicode"
)

func Unidecode(s string) string {
	str := ""
	for _, c := range s {
		if c <= unicode.MaxASCII {
			str += string(c)
			continue
		}
		if c > unicode.MaxRune {
			/* Ignore reserved chars */
			continue
		}
		if d, ok := transliterations[c]; ok {
			str += d
		}
	}
	return str
}
