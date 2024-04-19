package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	chars := []rune(str)
	var i, j int
	var result strings.Builder

	if len(chars) == 0 {
		return "", nil
	}
	if unicode.IsDigit(chars[0]) {
		return "", ErrInvalidString
	}

	for i, j = 0, 1; j < len(chars); i, j = i+1, j+1 {
		if unicode.IsDigit(chars[j]) {
			if unicode.IsDigit(chars[i]) {
				return "", ErrInvalidString
			}
			n, err := strconv.Atoi(string(chars[j]))
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(string(chars[i]), n))
		} else if !unicode.IsDigit(chars[i]) {
			result.WriteRune(chars[i])
		}
	}
	// last character is not processed in the cycle
	if !unicode.IsDigit(chars[i]) {
		result.WriteRune(chars[i])
	}

	return result.String(), nil
}
