package jz

import "testing"

func TestPrint1ToMaxOfNDigits(t *testing.T) {
	tests := []struct {
		test int
	}{
		{
			test: 1,
		},
		{
			test: 2,
		},
		{
			test: 4,
		},
	}

	for _, value := range tests {
		Print1ToMaxOfNDigits(value.test)
	}
}
