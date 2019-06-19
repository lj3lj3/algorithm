package jz

import (
	"reflect"
	"testing"
)

func TestPrintMatrix(t *testing.T) {
	tests := []struct {
		test [][]int
		want []int
	}{
		{
			test: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			test: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			test: [][]int{
				{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			test: [][]int{},
			want: []int{},
		},
	}

	for _, value := range tests {
		if result := PrintMatrix(value.test); !reflect.DeepEqual(result, value.want) {
			t.Errorf("PrintMatrix failed, result: %v, %v, want: %v", value.test, result, value.want)
		}
	}
}
