package a1

import (
	"errors"
	"math"
	"strings"
)

// ErrInvalidAtoiArgument ...
var ErrInvalidAtoiArgument = errors.New("a1: 's' cannot be empty, and must be alphabet only")

// MustAtoi ...
func MustAtoi(s string) int {
	i, err := Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Atoi ...
func Atoi(s string) (int, error) {
	if s == "" {
		return 0, ErrInvalidAtoiArgument
	}
	s = strings.ToUpper(s)
	var n float64
	var n2 byte
	var l = len(s)
	for i := 0; i < l; i++ {
		n2 = s[i] - 64
		if n2 > 26 {
			return 0, ErrInvalidAtoiArgument
		}
		n += float64(n2) * math.Pow(26, float64(l-i-1))
	}
	return int(n), nil
}
