package jz

import "testing"

func TestFind(t *testing.T) {
	tests := []struct {
		test   [][]int
		target int
		want   bool
	}{
		{
			test: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			want:   true,
		},
		{
			test: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			want:   false,
		},
	}

	for _, value := range tests {
		if result := find(value.test, value.target); result != value.want {
			t.Errorf("find failed, test: %v, target: %d, result: %t, want: %t", value.test, value.target, result, value.want)
		}
	}
}
