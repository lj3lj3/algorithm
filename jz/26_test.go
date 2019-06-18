package jz

import (
	"testing"
)

func TestHasSubtree(t *testing.T) {
	// construct link list
	tests := []struct {
		test *TreeNode
		tree *TreeNode
		want bool
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
			tree: &TreeNode{
				value: 20,
				left: &TreeNode{
					value: 15,
				},
				right: &TreeNode{
					value: 7,
				},
			},
			want: true,
		},
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
						left: &TreeNode{
							value: 15,
						},
					},
				},
			},
			tree: &TreeNode{
				value: 20,
				left: &TreeNode{
					value: 15,
				},
			},
			want: false,
		},
		{
			test: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 9,
				},
				right: &TreeNode{
					value: 20,
				},
			},
			tree: nil,
			want: false,
		},
	}

	for _, value := range tests {
		if result := HasSubtree(value.test, value.tree); result != value.want {
			t.Errorf("HasSubtree failed, test: %v, result: %v, want: %t", value.test, result, value.want)
		}
	}
}
