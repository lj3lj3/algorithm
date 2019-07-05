package jz

import (
	"reflect"
	"testing"
)

func TestPrintFromTopToBottom(t *testing.T) {
	tests := []struct {
		test *TreeNode
		want []int
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
			want: []int{3, 9, 20, 15, 7},
		},
		{
			test: &TreeNode{
				value: 3,
				right: &TreeNode{
					value: 9,
				},
			},
			want: []int{3, 9},
		},
		{
			test: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 9,
					left: &TreeNode{
						value: 8,
						right: &TreeNode{
							value: 9,
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
			want: []int{3, 9, 9, 8, 7, 7, 8, 9, 9},
		},
		{
			test: nil,
			want: []int{},
		},
	}

	for _, value := range tests {
		if result := PrintFromTopToBottom(value.test); !reflect.DeepEqual(result, value.want) {
			t.Errorf("PrintFromTopToBottom failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
