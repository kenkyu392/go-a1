package a1

import (
	"errors"
	"testing"
)

func TestAtoi(t *testing.T) {

	t.Run("case=valid--A", func(t *testing.T) {
		const want = 1
		if got, _ := Atoi("A"); got != want {
			t.Errorf("got: '%d', want: '%d'", got, want)
		}
	})

	t.Run("case=valid--AZ", func(t *testing.T) {
		const want = 52
		if got, _ := Atoi("AZ"); got != want {
			t.Errorf("got: '%d', want: '%d'", got, want)
		}
	})

	t.Run("case=valid--ZZ", func(t *testing.T) {
		const want = 702
		if got, _ := Atoi("ZZ"); got != want {
			t.Errorf("got: '%d', want: '%d'", got, want)
		}
	})

	t.Run("case=invalid", func(t *testing.T) {
		if _, err := Atoi(""); !errors.Is(err, ErrInvalidAtoiArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
		if _, err := Atoi("A$"); !errors.Is(err, ErrInvalidAtoiArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
		if _, err := Atoi("$A"); !errors.Is(err, ErrInvalidAtoiArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
	})
}

func TestMustAtoi(t *testing.T) {

	t.Run("case=valid--A", func(t *testing.T) {
		const want = 1
		if got := MustAtoi("A"); got != want {
			t.Errorf("got: '%d', want: '%d'", got, want)
		}
	})

	t.Run("case=valid--AZ", func(t *testing.T) {
		const want = 52
		if got := MustAtoi("AZ"); got != want {
			t.Errorf("got: '%d', want: '%d'", got, want)
		}
	})

	t.Run("case=valid--ZZ", func(t *testing.T) {
		const want = 702
		if got := MustAtoi("ZZ"); got != want {
			t.Errorf("got: '%d', want: '%d'", got, want)
		}
	})

	t.Run("case=invalid--empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic")
			}
		}()
		_ = MustAtoi("")
		if _, err := Atoi("A$"); !errors.Is(err, ErrInvalidAtoiArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
		if _, err := Atoi("$A"); !errors.Is(err, ErrInvalidAtoiArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
	})

	t.Run("case=invalid--suffix", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic")
			}
		}()
		_ = MustAtoi("A$")
		if _, err := Atoi("$A"); !errors.Is(err, ErrInvalidAtoiArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
	})

	t.Run("case=invalid--prefix", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic")
			}
		}()
		_ = MustAtoi("$A")
	})
}
