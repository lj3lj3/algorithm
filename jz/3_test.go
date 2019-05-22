package jz

import "testing"

func TestDuplicate(t *testing.T) {
	tests := []struct {
		test []int
		want int
	}{
		{
			test: []int{2, 3, 1, 0, 2, 5},
			want: 2,
		},
		{
			test: []int{2, 3, 1, 0, 2, 5, 5, 222, 2, 1, 3},
			want: 2,
		},
		{
			test: []int{2, 3, 1, 0, 0, 5, 5, 222, 11, 1, 3},
			want: 0,
		},
	}

	for _, value := range tests {
		if result := duplicate(value.test); result != value.want {
			t.Errorf("duplicated failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}
