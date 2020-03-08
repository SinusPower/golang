package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack makes basic unpacking of given string.
// This function allows skipping characters by '\'.
func Unpack(in string) (string, error) {
	if in == "" {
		return "", nil
	}

	var b strings.Builder
	runes := []rune(in)
	var prevRune rune
	skipped := false
	digitWritten := false

	for _, r := range runes {
		if prevRune == r && r != '\\' && !digitWritten {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(r) {
			if unicode.IsDigit(prevRune) && !digitWritten {
				return "", ErrInvalidString
			}
			if skipped {
				b.WriteRune(r)
				skipped = false
				digitWritten = true
				prevRune = r
				continue
			}
			count, _ := strconv.Atoi(string(r))                      // we know that r is a digit
			b.WriteString(strings.Repeat(string(prevRune), count-1)) // the first rune we put on previous iteration
		} else if r != '\\' {
			b.WriteRune(r) // put single or the first rune in the group
			digitWritten = true
		} else {
			if !skipped {
				skipped = true
			} else {
				b.WriteRune(r)
				skipped = false
				digitWritten = true
			}
		}
		prevRune = r
	}

	return b.String(), nil
}
