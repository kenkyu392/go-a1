package a1

import "strconv"

// Cell ...
func Cell(row, col int) (string, error) {
	a, err := Itoa(col)
	if err != nil {
		return "", err
	}
	return a + strconv.Itoa(row), nil
}
