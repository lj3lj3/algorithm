package jz

import (
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	tests := []struct {
		test     *TreeNode
		expected int
		want     [][]int
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
			expected: 22,
			want:     [][]int{},
		},
		{
			test: &TreeNode{
				value: 3,
				right: &TreeNode{
					value: 9,
					left: &TreeNode{
						value: 10,
					},
				},
			},
			expected: 22,
			want:     [][]int{{3, 9, 10}},
		},
		{
			test: &TreeNode{
				value: 10,
				left: &TreeNode{
					value: 5,
					left: &TreeNode{
						value: 4,
					},
					right: &TreeNode{
						value: 7,
					},
				},
				right: &TreeNode{
					value: 12,
					left: &TreeNode{
						value: 7,
					},
				},
			},
			expected: 22,
			want:     [][]int{{10, 5, 7}, {10, 12}},
		},
		{
			test:     nil,
			expected: 22,
			want:     [][]int{},
		},
	}

	for _, value := range tests {
		if result := findPath(value.test, value.expected); !reflect.DeepEqual(result, value.want) {
			t.Errorf("FindPath failed, test: %v, expected: %d, result: %v, want: %v", value.test, value.expected, result, value.want)
		}
	}
}
