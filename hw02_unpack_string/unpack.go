package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrAtoi          = errors.New("error convert str to int")
)

func Unpack(input string) (string, error) {
	var b strings.Builder
	var first rune
	for _, r := range input {
		if !unicode.IsDigit(r) {
			if first != 0 {
				b.WriteString(string(first))
			}
			first = r
		}
		if unicode.IsDigit(r) {
			if first == 0 {
				return "", ErrInvalidString
			}
			s, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrAtoi
			}
			b.WriteString(strings.Repeat(string(first), s))
			first = 0
		}
	}
	if first != 0 {
		b.WriteString(string(first))
	}

	return b.String(), nil
}
