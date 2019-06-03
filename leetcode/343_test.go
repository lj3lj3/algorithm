package leetcode

import "testing"

var tests = []struct {
	test int
	want int
}{
	{
		test: 2,
		want: 1,
	}, {
		test: 10,
		want: 36,
	},
}

func TestIntegerBreak(t *testing.T) {
	for _, value := range tests {
		if result := IntegerBreak(value.test); result != value.want {
			t.Errorf("MovingCount failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}

func TestIntegerBreakII(t *testing.T) {
	for _, value := range tests {
		if result := IntegerBreakII(value.test); result != value.want {
			t.Errorf("MovingCount failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}
