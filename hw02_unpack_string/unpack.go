package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var builder strings.Builder
	var prev rune
	var cur rune
	isFirstRune := true

	for _, r := range s {
		cur = r
		if isFirstRune {
			if unicode.IsDigit(r) {
				return "", ErrInvalidString
			}
			prev = r
			isFirstRune = false
			continue
		}
		if unicode.IsDigit(prev) && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(r) {
			count, _ := strconv.Atoi(string(r))
			if count != 0 {
				builder.WriteString(strings.Repeat(string(prev), count))
			}
		} else if !unicode.IsDigit(prev) {
			builder.WriteRune(prev)
		}

		prev = r
	}
	if !unicode.IsDigit(cur) {
		builder.WriteRune(cur)
	}
	return builder.String(), nil
}
