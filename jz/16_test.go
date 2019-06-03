package jz

import "testing"

func TestPower(t *testing.T) {
	tests := []struct {
		test     float64
		exponent int
		want     float64
	}{
		{
			test:     1,
			exponent: 0,
			want:     1,
		},
		{
			test:     2,
			exponent: -3,
			want:     float64(1) / 8,
		},
		{
			test:     2.3,
			exponent: 9,
			want:     float64(1801.1526614629986),
		},
	}

	for _, value := range tests {
		if result := Power(value.test, value.exponent); result != value.want {
			t.Errorf("Power failed, test: %f, exponent: %d, result: %f, want: %f", value.test, value.exponent, result, value.want)
		}
	}
}
