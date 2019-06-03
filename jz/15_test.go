package jz

import "testing"

func TestNumberOf1(t *testing.T) {
	tests := []struct {
		test int
		want int
	}{
		{
			test: 0,
			want: 0,
		}, {
			test: 1,
			want: 1,
		}, {
			test: 1024,
			want: 1,
		},
		{
			test: 15, // 10111
			want: 4,
		},
	}

	for _, value := range tests {
		if result := NumberOf1(value.test); result != value.want {
			t.Errorf("NumberOf1 failed, test: %d, result: %d, want: %d", value.test, result, value.want)
		}
	}
}
