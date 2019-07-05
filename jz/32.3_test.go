package jz

import (
	"reflect"
	"testing"
)

func TestPrintTreeInZhiFormat(t *testing.T) {
	tests := []struct {
		test *TreeNode
		want [][]int
	}{
		{
			test: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 9,
				},
				right: &TreeNode{
					value: 20,
					left: &TreeNode{
						value: 15,
					},
					right: &TreeNode{
						value: 7,
					},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			test: &TreeNode{
				value: 3,
				right: &TreeNode{
					value: 9,
				},
			},
			want: [][]int{{3}, {9}},
		},
		{
			test: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 1,
					left: &TreeNode{
						value: 2,
						right: &TreeNode{
							value: 4,
						},
					},
					right: &TreeNode{
						value: 7,
					},
				},
				right: &TreeNode{
					value: 9,
					left: &TreeNode{
						value: 7,
					},
					right: &TreeNode{
						value: 8,
						left: &TreeNode{
							value: 9,
						},
					},
				},
			},
			want: [][]int{{3}, {9, 1}, {2, 7, 7, 8}, {9, 4}},
		},
		{
			test: nil,
			want: [][]int{},
		},
	}

	for _, value := range tests {
		if result := PrintTreeInZhiFormat(value.test); !reflect.DeepEqual(result, value.want) {
			t.Errorf("PrintTreeInZhiFormat failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
