package jz

import (
	"testing"
)

func TestVerifySequenceOfBST(t *testing.T) {
	tests := []struct {
		test     *TreeNode
		expected []int
		want     bool
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
			expected: []int{9, 15, 7, 20, 3},
			want:     true,
		},
		{
			test: &TreeNode{
				value: 3,
				right: &TreeNode{
					value: 9,
				},
			},
			expected: []int{3, 9},
			want:     false,
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
			expected: []int{4, 2, 7, 1, 7, 9, 8, 9, 3},
			want:     true,
		},
		{
			test:     nil,
			expected: []int{},
			want:     true,
		},
	}

	for _, value := range tests {
		if result := VerifySequenceOfBST(value.test, value.expected); result != value.want {
			t.Errorf("VerifySequenceOfBST failed, test: %v, expected: %v, result: %t, want: %t", value.test, value.expected, result, value.want)
		}
	}
}
