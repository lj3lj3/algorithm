package jz

import (
	"reflect"
	"testing"
)

func TestDeleteDuplication(t *testing.T) {
	tests := []struct {
		test *LinkNode
		want *LinkNode
	}{
		{
			test: &LinkNode{
				value: 1,
				next: &LinkNode{
					value: 1,
				},
			},
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
				value: 1,
				next: &LinkNode{
					value: 2,
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
			want: &LinkNode{
				value: 1,
				next: &LinkNode{
					value: 2,
					next: &LinkNode{
						value: 4,
					},
				},
			},
		},
	}

	for _, value := range tests {
		if result := DeleteDuplication(value.test); !reflect.DeepEqual(result, value.want) {
			t.Errorf("DeleteDuplication failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
