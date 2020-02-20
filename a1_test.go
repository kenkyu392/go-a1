package a1

import "testing"

func TestCell(t *testing.T) {
	t.Run("case=valid--R1C1", func(t *testing.T) {
		const want = "A1"
		if got, _ := Cell(1, 1); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--R5C26", func(t *testing.T) {
		const want = "Z5"
		if got, _ := Cell(5, 26); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--R20C52", func(t *testing.T) {
		const want = "AZ20"
		if got, _ := Cell(20, 52); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=valid--R52C702", func(t *testing.T) {
		const want = "ZZ52"
		if got, _ := Cell(52, 702); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})
}
