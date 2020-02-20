package a1

import (
	"errors"
	"testing"
)

func TestItoa(t *testing.T) {

	t.Run("case=valid--1", func(t *testing.T) {
		const want = "A"
		if got, _ := Itoa(1); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--52", func(t *testing.T) {
		const want = "AZ"
		if got, _ := Itoa(52); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--702", func(t *testing.T) {
		const want = "ZZ"
		if got, _ := Itoa(702); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=invalid", func(t *testing.T) {
		if _, err := Itoa(0); !errors.Is(err, ErrInvalidItoaArgument) {
			t.Errorf("got: '%#v', want: '%#v'", err, ErrInvalidAtoiArgument)
		}
	})
}

func TestMustItoa(t *testing.T) {

	t.Run("case=valid--1", func(t *testing.T) {
		const want = "A"
		if got := MustItoa(1); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--52", func(t *testing.T) {
		const want = "AZ"
		if got := MustItoa(52); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--702", func(t *testing.T) {
		const want = "ZZ"
		if got := MustItoa(702); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic")
			}
		}()
		_ = MustItoa(0)
	})
}
