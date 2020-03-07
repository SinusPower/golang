package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(in string) (string, error) {
	var result string = in

	return result, nil
}
