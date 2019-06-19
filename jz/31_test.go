package jz

import (
	"testing"
)

func TestIsPopOrder(t *testing.T) {
	tests := []struct {
		test [][]int
		want bool
	}{
		{
			test: [][]int{
				{1, 2, 3, 4, 5},
				{4, 5, 3, 2, 1},
			},
			want: true,
		},
		{
			test: [][]int{
				{1, 2, 3, 4, 5},
				{4, 3, 5, 1, 2},
			},
			want: false,
		},
		{
			test: [][]int{
				{},
				{},
			},
			want: true,
		},
	}

	for _, value := range tests {
		if result := IsPopOrder(value.test[0], value.test[1]); result != value.want {
			t.Errorf("IsPopOrder failed, result: %v, %t, want: %t", value.test, result, value.want)
		}
	}
}
