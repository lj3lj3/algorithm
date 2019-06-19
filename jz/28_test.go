package jz

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		test *TreeNode
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
			want: false,
		},
		{
			test: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 9,
				},
				right: &TreeNode{
					value: 9,
				},
			},
			want: true,
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
			want: true,
		},
		{
			test: nil,
			want: true,
		},
	}

	for _, value := range tests {
		if result := IsSymmetric(value.test); result != value.want {
			t.Errorf("IsSymmetric failed, result: %v, want: %v", value.test, value.want)
		}
	}
}
