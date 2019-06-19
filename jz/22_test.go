package jz

import (
	"reflect"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	tests := []struct {
		test *LinkNode
		k    int
		want *LinkNode
	}{
		{
			test: &LinkNode{
				value: 1,
				next: &LinkNode{
					value: 1,
				},
			},
			k:    3,
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
			k: 4,
			want: &LinkNode{
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
								next: &LinkNode{
									value: 4,
								},
							},
						},
					},
				},
			},
			k: 2,
			want: &LinkNode{
				value: 3,
				next: &LinkNode{
					value: 4,
				},
			},
		},
	}

	for _, value := range tests {
		if result := FindKthToTail(value.test, value.k); !reflect.DeepEqual(result, value.want) {
			t.Errorf("FindKthToTail failed, test: %v, k: %d, result: %v, want: %v", value.test, value.k, result, value.want)
		}
	}
}
