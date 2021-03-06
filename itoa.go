package a1

import (
	"errors"
)

// ErrInvalidItoaArgument ...
var ErrInvalidItoaArgument = errors.New("a1: 'i' cannot be empty, and must be > 0")

// MustItoa ...
func MustItoa(i int) string {
	s, err := Itoa(i)
	if err != nil {
		panic(err)
	}
	return s
}

// Itoa ...
func Itoa(i int) (string, error) {
	if i <= 0 {
		return "", ErrInvalidItoaArgument
	}
	var s string
	var i2 int
	for i > 0 {
		i2 = (i - 1) % 26
		s = string(rune(i2+65)) + s
		i = (i - i2) / 26
	}
	return s, nil
}
