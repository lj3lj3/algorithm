package jz

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		test []*LinkNode
		want *LinkNode
	}{
		{
			test: []*LinkNode{
				{
					value: 1,
					next: &LinkNode{
						value: 2,
						next: &LinkNode{
							value: 4,
							next: &LinkNode{
								value: 6,
								next: &LinkNode{
									value: 6,
								},
							},
						},
					},
				},
				{
					value: 1,
					next: &LinkNode{
						value: 3,
						next: &LinkNode{
							value: 5,
						},
					},
				},
			},
			want: &LinkNode{
				value: 1,
				next: &LinkNode{
					value: 1,
					next: &LinkNode{
						value: 2,
						next: &LinkNode{
							value: 3,
							next: &LinkNode{
								value: 4,
								next: &LinkNode{
									value: 5,
									next: &LinkNode{
										value: 6,
										next: &LinkNode{
											value: 6,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			test: []*LinkNode{
				{
					value: 1,
				},
				nil,
			},
			want: &LinkNode{
				value: 1,
			},
		},
		{
			test: []*LinkNode{
				nil,
				nil,
			},
			want: nil,
		},
	}

	for _, value := range tests {
		if result := Merge(value.test[0], value.test[1]); !reflect.DeepEqual(result, value.want) {
			t.Errorf("Merge failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
