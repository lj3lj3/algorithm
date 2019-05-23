package jz

import (
	"testing"
)

func TestGetNext(t *testing.T) {
	length := 7
	nodes := make([]*TreeLinkNode, length)
	for i := 0; i < length; i++ {
		nodes[i] = &TreeLinkNode{
			value: i,
		}
	}

	// the tree for testing
	tree := nodes[0]
	nodes[0].left = nodes[1]
	nodes[1].parent = nodes[0]
	nodes[1].left = nodes[2]
	nodes[2].parent = nodes[1]
	nodes[1].right = nodes[3]
	nodes[3].parent = nodes[1]
	nodes[3].left = nodes[4]
	nodes[4].parent = nodes[3]
	nodes[3].right = nodes[5]
	nodes[5].parent = nodes[3]
	nodes[0].right = nodes[6]
	nodes[6].parent = nodes[0]

	tests := []struct {
		test *TreeLinkNode
		want *TreeLinkNode
	}{
		{
			test: tree.left,
			want: tree.left.right.left,
		},
		{
			test: tree.left.right.right,
			want: tree.right,
		},
	}

	for _, value := range tests {
		if result := GetNext(value.test); result.value != value.want.value {
			t.Errorf("GetNext failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
