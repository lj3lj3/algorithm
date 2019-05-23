package jz

import (
	"reflect"
	"testing"
)

func TestReConstructBinaryTree(t *testing.T) {
	tests := []struct {
		test [][]int
		want *TreeNode
	}{
		{
			test: [][]int{
				{3, 9, 20, 15, 7},
				{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
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
		},
		{
			test: [][]int{
				{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 7},
				{8, 4, 9, 2, 10, 5, 11, 1, 6, 12, 3, 7},
			},
			want: &TreeNode{
				value: 1,
				left: &TreeNode{
					value: 2,
					left: &TreeNode{
						value: 4,
						left: &TreeNode{
							value: 8,
						},
						right: &TreeNode{
							value: 9,
						},
					},
					right: &TreeNode{
						value: 5,
						left: &TreeNode{
							value: 10,
						},
						right: &TreeNode{
							value: 11,
						},
					},
				},
				right: &TreeNode{
					value: 3,
					left: &TreeNode{
						value: 6,
						right: &TreeNode{
							value: 12,
						},
					},
					right: &TreeNode{
						value: 7,
					},
				},
			},
		},
	}

	for _, value := range tests {
		if result := ReConstructBinaryTree(value.test[0], value.test[1]); !reflect.DeepEqual(result, value.want) {
			t.Errorf("printListFromTailToHead failed, test: %v, result: %s, want: %s", value.test, result.String(), value.want.String())
		}
	}
}
