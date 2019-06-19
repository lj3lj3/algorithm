package jz

import (
	"reflect"
	"testing"
)

func TestEntryNodeOfLoop(t *testing.T) {
	// Init the list of LinkNodes
	length := 10
	linkNodes := make([]*LinkNode, length)
	for i := 0; i < length; i++ {
		linkNodes[i] = &LinkNode{i + 1, nil}
		if i != 0 {
			linkNodes[i-1].next = linkNodes[i] // link it
		}
	}
	// Create a loop
	entry := 5
	linkNodes[length-1].next = linkNodes[entry]

	// Construct linked list
	tests := []struct {
		test *LinkNode
		want *LinkNode
	}{
		{
			test: linkNodes[0],
			want: linkNodes[entry],
		},
		{
			test: &LinkNode{
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
								},
							},
						},
					},
				},
			},
			want: nil,
		},
	}

	for _, value := range tests {
		if result := EntryNodeOfLoop(value.test); !reflect.DeepEqual(result, value.want) {
			t.Errorf("EntryNodeOfLoop failed, test: %v, result: %v, want: %v", value.test, result, value.want)
		}
	}
}
