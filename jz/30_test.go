package jz

import "testing"

const actionTop int8 = 2
const actionMin int8 = 3

func TestStackWithMinFunc(t *testing.T) {
	tests := []struct {
		test []*action
		want []int
	}{
		{
			test: []*action{
				{
					action: actionPush,
					data:   2,
				},
				{
					action: actionPush,
					data:   4,
				},
				{
					action: actionTop,
					data:   4,
				},
				{
					action: actionPush,
					data:   0,
				},
				{
					action: actionPush,
					data:   80,
				},
				{
					action: actionMin,
					data:   0,
				},
				{
					action: actionPop,
				},
				{
					action: actionMin,
					data:   0,
				},
				{
					action: actionPop,
				},
				{
					action: actionMin,
					data:   2,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionMin,
					data:   0,
				},
				{
					action: actionPop,
				},
			},
			want: []int{2, 4, 80, 0, 20},
		},
	}

	for _, value := range tests {
		stack := new(StackWithMinFunc)

		for _, test := range value.test {
			if test.action == actionPush {
				stack.push(test.data)
			} else if test.action == actionPop {
				stack.pop()
			} else if test.action == actionTop {
				if result := stack.top(); result != test.data {
					// Failed
					t.Errorf("Top failed, test: %v, result: %d, want: %d", test, result, test.data)
				}
			} else if test.action == actionMin {
				if result := stack.min(); result != test.data {
					// Failed
					t.Errorf("Min failed, test: %v, result: %d, want: %d", test, result, test.data)
				}
			}
		}
	}
}
