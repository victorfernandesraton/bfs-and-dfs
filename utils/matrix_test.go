package utils

import "testing"

func TestExpected(t *testing.T) {
	m := [][]uint8{
		{
			1, 2, 3,
		},
		{
			4, 5, 6,
		},
		{
			7, 8, 9,
		},
	}

	if IsAdjacence(m) == false {
		t.Error("Expected true for square matrix")
	}
}
