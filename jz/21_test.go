package jz

import (
	"reflect"
	"testing"
)

type testStruct struct {
	test []int
	want []int
}

func TestReOrderArray(t *testing.T) {
	tests := []testStruct{
		{
			test: []int{
				1, 2, 3, 4, 5, 6,
			},
			want: []int{
				1, 3, 5, 2, 4, 6,
			},
		},
		{
			test: []int{},
			want: []int{},
		},
		{
			test: []int{
				1, 2, 2, 5, 5, 6,
			},
			want: []int{
				1, 5, 5, 2, 2, 6,
			},
		},
	}

	tests2 := make([]testStruct, len(tests))
	copy(tests2, tests)

	for _, value := range tests {
		ReOrderArray(value.test)
		if !reflect.DeepEqual(value.test, value.want) {
			t.Errorf("ReOrderArray failed, results: %v, want: %v", value.test, value.want)
		}
	}

	for _, value := range tests2 {
		ReOrderArrayII(value.test)
		if !reflect.DeepEqual(value.test, value.want) {
			t.Errorf("ReOrderArrayII failed, results: %v, want: %v", value.test, value.want)
		}
	}
}
