package jz

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		test *LinkNode
		want *LinkNode
	}{
		{
			test: nil,
			want: nil,
		},
		{
			test: &LinkNode{
				value: 1,
				next: &LinkNode{
					value: 2,
					next: &LinkNode{
						value: 3,
						next: &LinkNode{
							value: 3,
							next: &LinkNode{
								value: 3,
							},
						},
					},
				},
			},
			want: &LinkNode{
				value: 3,
				next: &LinkNode{
					value: 3,
					next: &LinkNode{
						value: 3,
						next: &LinkNode{
							value: 2,
							next: &LinkNode{
								value: 1,
							},
						},
					},
				},
			},
		},
		{
			test: &LinkNode{
				value: 1,
			},
			want: &LinkNode{
				value: 1,
			},
		},
	}

	for _, value := range tests {
		if result := ReverseList(value.test); !reflect.DeepEqual(result, value.want) {
			t.Errorf("ReverseList failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
