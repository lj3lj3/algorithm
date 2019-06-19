package jz

import (
	"reflect"
	"testing"
)

func TestMirror(t *testing.T) {
	// construct link list
	tests := []struct {
		test *TreeNode
		want *TreeNode
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
			want: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 20,
					left: &TreeNode{
						value: 7,
					},
					right: &TreeNode{
						value: 15,
					},
				},
				right: &TreeNode{
					value: 9,
				},
			},
		},
		{
			test: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 9,
				},
			},
			want: &TreeNode{
				value: 3,
				right: &TreeNode{
					value: 9,
				},
			},
		},
		{
			test: &TreeNode{
				value: 3,
			},
			want: &TreeNode{
				value: 3,
			},
		},
		{
			test: nil,
			want: nil,
		},
	}

	for _, value := range tests {
		Mirror(value.test)
		if !reflect.DeepEqual(value.test, value.want) {
			t.Errorf("Mirror failed, result: %v, want: %v", value.test, value.want)
		}
	}
}
